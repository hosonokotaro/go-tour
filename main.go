package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	fmt.Println("nil の例、struct を pointer として受け取り、使おうとした場合", ms)

	// NOTE: new struct すると、struct の foo が初期化される。
	ms = new(myStruct)

	// ms.foo は (*ms).foo の syntax sugar
	ms.foo = 24
	fmt.Println("フィールド foo を変更する", ms.foo)
}

type myStruct struct {
	foo int
}
