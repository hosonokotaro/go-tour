package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println(j)
		}(i)
	}
}
