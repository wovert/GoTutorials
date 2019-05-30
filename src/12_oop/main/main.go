package main

import (
	"fmt"
)

type Integer struct {
	value int
}

type Point struct {
	px float32
	py float32
}

func (point *Point) setXY(px, py float32) {
	point.px = px
	point.py = py
}

func (point *Point) getXY() (float32, float32) {
	return point.px, point.py
}

func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}


func main() {
	a := Integer{1}
	b := Integer{2}

	fmt.Printf("%v\n", a.compare(b))

	point := new(Point)
	point.setXY(1.23, 4.56)
	px, py := point.getXY()
	fmt.Println(px, py)
}