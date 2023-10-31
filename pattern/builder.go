package main

import "fmt"

// Car is a complex object that can be constructed using Builder.
type Car struct {
	Wheels int
	Seats  int
	Color  string
}

// Builder interface defines the methods to construct a Car.
type Builder interface {
	SetWheels(wheels int) Builder
	SetSeats(seats int) Builder
	SetColor(color string) Builder
	Build() Car
}

// CarBuilder is the concrete implementation of Builder interface.
type CarBuilder struct {
	car Car
}

func (b *CarBuilder) SetWheels(wheels int) Builder {
	b.car.Wheels = wheels
	return b
}

func (b *CarBuilder) SetSeats(seats int) Builder {
	b.car.Seats = seats
	return b
}

func (b *CarBuilder) SetColor(color string) Builder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) Build() Car {
	return b.car
}

func main() {
	builder := &CarBuilder{}
	car := builder.SetWheels(4).SetSeats(5).SetColor("Red").Build()

	fmt.Printf("Car: Wheels=%d, Seats=%d, Color=%s\n", car.Wheels, car.Seats, car.Color)
}
