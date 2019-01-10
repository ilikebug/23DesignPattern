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

/*
工厂方法模式：

通过工厂方法模式的类图可以看到，工厂方法模式有四个要素：

	工厂接口。工厂接口是工厂方法模式的核心，与调用者直接交互用来提供产品。在实际编程中，有时候也会使用一个抽象类来作为与调用者交互的接口，其本质上是一样的。
	工厂实现。在编程中，工厂实现决定如何实例化产品，是实现扩展的途径，需要有多少种产品，就需要有多少个具体的工厂实现。
	产品接口。产品接口的主要目的是定义产品的规范，所有的产品实现都必须遵循产品接口定义的规范。产品接口是调用者最为关心的，产品接口定义的优劣直接决定了调用者代码的稳定性。同样，产品接口也可以用抽象类来代替，但要注意最好不要违反里氏替换原则。
	产品实现。实现产品接口的具体类，决定了产品在客户端中的具体行为。
适用场景：

	不管是简单工厂模式，工厂方法模式还是抽象工厂模式，他们具有类似的特性，所以他们的适用场景也是类似的。

	首先，作为一种创建类模式，在任何需要生成复杂对象的地方，都可以使用工厂方法模式。有一点需要注意的地方就是复杂对象适合使用工厂模式，而简单对象，特别是只需要通过new就可以完成创建的对象，无需使用工厂模式。如果使用工厂模式，就需要引入一个工厂类，会增加系统的复杂度。

	其次，工厂模式是一种典型的解耦模式，迪米特法则在工厂模式中表现的尤为明显。假如调用者自己组装产品需要增加依赖关系时，可以考虑使用工厂模式。将会大大降低对象之间的耦合度。

	再次，由于工厂模式是依靠抽象架构的，它把实例化产品的任务交由实现类完成，扩展性比较好。也就是说，当需要系统有比较好的扩展性时，可以考虑工厂模式，不同的产品用不同的实现工厂来组装。
*/
