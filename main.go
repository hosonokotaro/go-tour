package main

import (
	"fmt"
	// "math/rand"
	// "math"
	// "time"
)

func main() {
	// ランダム値を unixtime で初期化する
	// rand.Seed(time.Now().UnixNano())

	// 分割代入の例
	// a, b := swap("hello", "world")

	a, b := split(17)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

// return 値に名前を付けると、変数宣言と同様の効果を得る。この return は、戻り値すべてを返す
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// return に複数の変数を指定できる、取り出すときはどうするかまだわからない
func swap(a, b string) (string, string) {
	return b, a
}

// まとめて型を指定するとき
func add(x, y int) int {
	return x + y
}
