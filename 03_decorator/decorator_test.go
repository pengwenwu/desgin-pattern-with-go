package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpAddDecorator(c, 5)
	c = WarpMulDecorator(c, 2)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
}