package main

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type CarBuilder struct {
	color    Color
	wheels   Wheels
	topSpeed Speed
}

func NewBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) Paint(c Color) *CarBuilder {
	b.color = c
	return b
}

func (b *CarBuilder) Wheels(w Wheels) *CarBuilder {
	b.wheels = w
	return b
}

func (b *CarBuilder) TopSpeed(s Speed) *CarBuilder {
	b.topSpeed = s
	return b
}

func (b *CarBuilder) Build() *Car {
	return &Car{
		color:    b.color,
		wheels:   b.wheels,
		topSpeed: b.topSpeed,
	}
}

type Car struct {
	color    Color
	wheels   Wheels
	topSpeed Speed
}

func (c *Car) Drive() error {
	fmt.Printf("Driving a %s car with %s wheels at %v MPH\n", c.color, c.wheels, c.topSpeed)
	return nil
}

func (c *Car) Stop() error {
	fmt.Println("Car stopped")
	return nil
}

func main() {
	assembly := NewBuilder().Paint(RedColor)

	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	_ = familyCar.Drive()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	_ = sportsCar.Drive()

}
