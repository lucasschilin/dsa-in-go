package main

import (
	"fmt"

	"github.com/lucasschilin/dsa-in-go/leetcode"
)

func main() {
	fns := leetcode.LengthOfLogestSubstring("abcabcbb")

	for name, fn := range fns {
		exec(name, fn)
	}
}

func exec(name string, fn func() any) {
	fmt.Println("\n" + name + ":")
	fmt.Println("Resultado:", fn())
}
