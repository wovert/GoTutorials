package main

import (
  "fmt"
)

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("cat实现了move接口")
}

func (c *cat)eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
  
}