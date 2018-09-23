package main

import "fmt"

func main() {
	m := &MyStruct{"Lars"}
	question(m)
}

func question(something MyInterface) {
	something.Print()
	fmt.Println("How are you?")

	_, ok := something.(*MyStruct)
	if ok {
		fmt.Println("I'm fine")
	} else {
		fmt.Println("Can be better")
	}
}

type MyInterface interface {
	Print()
}

type MyStruct struct {
	Name string
}

func (m *MyStruct) Print() {
	fmt.Printf("Hello %s!\n", m.Name)
}
