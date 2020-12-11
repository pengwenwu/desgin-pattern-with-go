package iterator

import "fmt"

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

type Aggregate interface {
	Iterator() Iterator
}

type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumberIterator{
		numbers: n,
		next:    n.start,
	}
}

type NumberIterator struct {
	numbers *Numbers
	next int
}

func (n *NumberIterator) First() {
	n.next = n.numbers.start
}

func (n *NumberIterator) IsDone() bool {
	return n.next > n.numbers.end
}

func (n *NumberIterator) Next() interface{} {
	if !n.IsDone() {
		next := n.next
		n.next++
		return next
	}
	return nil
}

func IteratorPrint(i Iterator)  {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}
