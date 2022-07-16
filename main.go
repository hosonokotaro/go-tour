package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a

	fmt.Println("pointer はアドレスを受け取る", a, b)
	fmt.Println("pointer を値として受け取る", a, *b)

	a = 27
	fmt.Println("a を変更する", a, *b)
	*b = 35
	fmt.Println("b を変更する", a, *b)
}
