package main

import (
	"fmt"
)

func main() {
	list1 := map[string]string{"key1": "aaa", "key2": "bbb", "key3": "ccc"}
	fmt.Println("list1", list1)

	list2 := list1

	list1["key1"] = "ddd"
	fmt.Println("list2 が list1[\"key1\"] の変更によって同じように変わっている", list2)
}
