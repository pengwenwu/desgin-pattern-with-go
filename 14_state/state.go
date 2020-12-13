package state

import "fmt"

type Week interface {
	Today()
	Next(ctx *DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{today:&Sunday{}}
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

type Sunday struct {

}

func (*Sunday) Today() {
	fmt.Printf("Sunday\n")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

type Monday struct {

}

func (m *Monday) Today() {
	fmt.Printf("Monday\n")
}

func (m *Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct {

}

func (t Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

func (t Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct {

}

func (w Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (w Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct {

}

func (t Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (t Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct {

}

func (f Friday) Today() {
	fmt.Printf("Friday\n")
}

func (f Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct {

}

func (s Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (s Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}


