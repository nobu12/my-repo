package main

import "fmt"

var version string // ビルド時にldflagsフラグ経由でバージョンを埋め込むための変数

func main() {
	fmt.Printf("ddExamdaplardeeedeaAdd %s\n", version)
}
