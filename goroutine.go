package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type OriginData struct {
	Uuid     string    `json:"uuid"`
	Date     time.Time `json:"date"`
	Value    float32   `json:"value"`
	Location string    `json:"location"` // Latitude/Longitude
}

type WeatherInfo struct {
	Date        string `json:"date"`
	Temperature string `json:"temperature"`
}

type TransformedData struct {
	OriginData
	weather WeatherInfo
}

func mapWeatherInfo(d OriginData, ch chan<- TransformedData, wg *sync.WaitGroup) error {
	defer wg.Done()

	parts := strings.Split(d.Location, "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid location: %v", d.Location)
	}
	lat, lng := parts[0], parts[1]

	resp, err := http.Get(fmt.Sprintf("https://exampleapi.com/weather?date=%v&lat=%v&lng=%v", d.Date, lat, lng))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var weather WeatherInfo
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return err
	}

	transformed := TransformedData{
		d,
		weather,
	}
	ch <- transformed
	return nil
}

func LoadDW() {
	originDB, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}

	dwDB, err := sql.Open("postgres", "user=password dbname=dbname sslmode=disable")
	if err != nil {
		panic(err)
	}

	rows, err := originDB.Query("SELECT uuid, date, value, location FROM deliveries")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tc := make(chan TransformedData)
	var wg sync.WaitGroup

	go func() {
		for t := range tc {
			if _, err := dwDB.Exec("INSERT INTO dw_deliveries (uuid, date, value, location, temperature) VALUES ($1, $2, $3, $4, $5)", t.Uuid, t.Date, t.Value, t.Location, t.weather.Temperature); err != nil {
				panic(err)
			}
		}
	}()

	for rows.Next() {
		var uuid string
		var timestamp string
		var value float32
		var location string

		if err := rows.Scan(&uuid, &timestamp, &value, &location); err != nil {
			panic(err)
		}

		date, err := time.Parse("2006-01-02 15:04:05", timestamp)
		if err != nil {
			panic(err)
		}

		wg.Add(1)
		go mapWeatherInfo(OriginData{
			Uuid:     uuid,
			Date:     date,
			Value:    value,
			Location: location,
		}, tc, &wg)
	}

	wg.Wait()
	close(tc)
}
