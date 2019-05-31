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

type Animal interface {
	Fly()
	Run()
}

type Animal2 interface {
	Fly()
}

type Bird struct {

}

func (bird Bird) Fly() {
	fmt.Println("Bird is flying!!!")
}

func (bird Bird) Run() {
	fmt.Println("Bird is Running!!!")
}


func main() {
	// a := Integer{1}
	// b := Integer{2}

	// fmt.Printf("%v\n", a.compare(b))

	// point := new(Point)
	// point.setXY(1.23, 4.56)
	// px, py := point.getXY()
	// fmt.Println(px, py)

	// var animal Animal
	// var animal2 Animal2

	// bird := new(Bird)

	// animal = bird
	// animal2 = bird

	// animal2.Fly()
	// animal.Run()

	// animal = animal2 // error
	// animal2 = animal // 大的接口赋值给小的接口
	// animal2.Fly()

	var v1 interface{}
	v1 = 6.76
	
	// 判断接口类型错误
	if v, ok := v1.(float64); ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println(v, ok)
	}
	switch v := v1.(type) {
		case int:
		case float32:
		case float64:
			fmt.Printf("%v is float64", v)
		case string:
			fmt.Printf("%vt is string", v)
	}
}