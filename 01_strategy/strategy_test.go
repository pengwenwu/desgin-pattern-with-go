package strategy

import (
	"testing"
)

func TestNewDuckContext(t *testing.T) {
	duck := NewDuckContext(&FlyWithWings{}, &Quack{})
	duck.PerformFly()
	duck.PerformQuack()

	duck = NewDuckContext(&FlyNoWay{}, &Squeak{})
	duck.PerformFly()
	duck.PerformQuack()
}