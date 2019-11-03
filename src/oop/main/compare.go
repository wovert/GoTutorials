package main


import (
	"fmt"
)

type Integer struct {
	value int
}
// 定义Integer结构体的方法 compare
func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}

type Point struct {
	x float32
	y float32
}

// func(接受任何类型)
func (point *Point) setXY(x, y float32) {
	point.x = x
	point.y = y
}

func(point *Point) getXY() (float32, float32) {
	return point.x, point.y
}

func main() {
	a := Integer{1}
	b := Integer{2}
	fmt.Printf("%v", b.compare(a))
	fmt.Println()

	point := new(Point)
	point.setXY(1.23, 4.36)
	x , y := point.getXY()
	fmt.Println(x, y)

}