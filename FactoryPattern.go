package main

import "fmt"

type Operater interface {
	Operate(x int, y int) int
}

type Addition struct{}
type Subtraction struct{}

type Factory struct{}

func (a *Addition) Operate(x int, y int) int {
	result := x + y
	return result
}

func (s *Subtraction) Operate(x int, y int) int {
	if x > y {
		result := x - y
		return result
	} else {
		result := y - x
		return result
	}
}

func (f Factory) CreateOperate(symbol string) Operater {
	switch symbol {
	case "+":
		return &Addition{}
	case "-":
		return &Subtraction{}
	default:
		panic("无效运算符号")
	}
}

func main() {
	var f Factory
	add := f.CreateOperate("+")
	result := add.Operate(2, 3)
	fmt.Printf("result is %d", result)
}
