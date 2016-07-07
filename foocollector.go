package main

import (
	"fmt"
)

type FooCollector struct {
	foo string
}

func NewFooCollector() (FooCollector) {
	return &fooCollector {
		foo: "bar"
	}
}

func main() {
	rad := NewFooCollector()
	fmt.Println(rad.foo)
}