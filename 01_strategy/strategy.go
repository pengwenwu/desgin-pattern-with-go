package strategy

import "fmt"

type Duck struct {
	FlyBehavior FlyBehavior
	QuackBehavior QuackBehavior
}

func (duck *Duck) PerformFly() {
	duck.FlyBehavior.Fly()
}

func (duck *Duck) PerformQuack() {
	duck.QuackBehavior.Quack()
}

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

type FlyWithWings struct {
}

func (*FlyWithWings) Fly() {
	fmt.Println("I am flying!")
}

type FlyNoWay struct {
}

func (*FlyNoWay) Fly() {
	fmt.Println("I can not fly!")
}

// 呱呱叫
type Quack struct {
}

// 不会叫
type MuteQuack struct {
}

// 吱吱叫
type Squeak struct {
}

func (*Quack) Quack() {
	fmt.Println("Quack!")
}

func (*Squeak) Quack() {
	fmt.Println("Squeak!")
}

func NewDuckContext(flyBehavior FlyBehavior, quackBehavior QuackBehavior) *Duck {
	return &Duck{
		FlyBehavior:  flyBehavior,
		QuackBehavior: quackBehavior,
	}
}
