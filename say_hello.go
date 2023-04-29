package go_module_sample

import "fmt"

func SayHello() string {
	return "Hello"
}

func SayHelloName(name string) {
	fmt.Println("Hello", name)
}