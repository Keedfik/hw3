package main

import (
	"fmt"
	"math"
)

type Shape struct {
	Name string
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	return 0.0
}

type Circle struct {
	Shape
	Radius float64
}

func (cir Circle) Area() float64 {
	return math.Pi * cir.Radius * cir.Radius
}

type Rectangle struct {
	Shape
	Length float64
	Width  float64
}

func (rec Rectangle) Area() float64 {
	return rec.Length * rec.Width
}

func main() {
	var rad float64
	fmt.Println("Input radius")
	fmt.Scanf("%f\n", &rad)

	circle := Circle{Shape{"Circle"}, rad}
	fmt.Printf("%s: Area = %.1f\n", circle.GetName(), circle.Area())

	var len, wid float64
	fmt.Println("Input length")
	fmt.Scanf("%f\n", &len)
	fmt.Println("Input width")
	fmt.Scanf("%f\n", &wid)

	rectangle := Rectangle{Shape{"Rectangle"}, len, wid}
	fmt.Printf("%s: Area = %.1f\n", rectangle.GetName(), rectangle.Area())
}
