package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Taro",
	}
	fmt.Println("geeter struct を実装した g", g)

	g.greet()
	fmt.Println("name を変更した g", g)
}

type greeter struct {
	greeting string
	name     string
}

// NOTE: *greeter は pointer として受け取り、その参照を変数内で変更している例
// もし、pointer ではなく value を渡した場合はコストがかかることを注意される
func (g *greeter) greet() {
	fmt.Println("g の struct はここに渡されている。greet() は main 関数から実行できる", g.greeting, g.name)
	g.name = "幽霊"
}
