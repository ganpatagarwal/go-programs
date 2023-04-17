package main

import (
	"fmt"
	"log"
)

type first struct {
	name string
}

func (f first) greetFirst() string {
	return fmt.Sprintf("hello: %s", f.name)
}

func (f first) greet() string {
	return fmt.Sprintf("hello: %s", f.name)
}

func newFirst(name string) first {
	return first{name: name}
}

type second struct {
	name string
}

func (f *second) greetSecond() string {
	return fmt.Sprintf("hello: %s", f.name)
}

func (f *second) greet() string {
	return fmt.Sprintf("hello: %s", f.name)
}

func newSecond(name string) *second {
	return &second{name: name}
}

type greet interface {
	greet() string
}

func knownInterface(i greet) string {
	return i.greet()
}

func unknownInterface(i interface{}) string {
	f, ok := i.(first)
	if ok {
		return f.greetFirst()
	}

	s, ok := i.(*second)
	if ok {
		return s.greetSecond()
	}

	return "invalid assertion"
}

func main() {
	log.Println(unknownInterface(newFirst("John")))
	log.Println(unknownInterface(newSecond("Doe")))

	log.Println(knownInterface(newFirst("first")))
	log.Println(knownInterface(newFirst("second")))
}
