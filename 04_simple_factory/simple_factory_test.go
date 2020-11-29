package simple_factory

import (
	"testing"
)

func Test_helloApi_Say(t *testing.T) {
	api := NewApi(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
}

func Test_hiApi_Say(t *testing.T) {
	api := NewApi(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Type2 test fail")
	}
}