package main

import (
	"fmt"
	"strings"
)

func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func main() {
	//零参数
	fmt.Println(joinstr())

	//多个参数
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println(joinstr("Geeks", "for", "Geeks"))
	fmt.Println("G", "E", "E", "k", "S")

}
