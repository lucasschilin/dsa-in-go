package main

import (
	"fmt"

	"github.com/lucasschilin/dsa-in-go/leetcode"
)

// l1 representa a lista [9,9,9,9,9,9,9]
var l1 = leetcode.ListNode{
	Val: 9, // primeiro nó
	Next: &leetcode.ListNode{
		Val: 9, // segundo nó
		Next: &leetcode.ListNode{
			Val: 9, // terceiro nó
			Next: &leetcode.ListNode{
				Val: 9, // quarto nó
				Next: &leetcode.ListNode{
					Val: 9, // quinto nó
					Next: &leetcode.ListNode{
						Val: 9, // sexto nó
						Next: &leetcode.ListNode{
							Val:  9,   // sétimo nó
							Next: nil, // fim da lista
						},
					},
				},
			},
		},
	},
}

// l2 representa a lista [9,9,9,9]
var l2 = leetcode.ListNode{
	Val: 9, // primeiro nó
	Next: &leetcode.ListNode{
		Val: 9, // segundo nó
		Next: &leetcode.ListNode{
			Val: 9, // terceiro nó
			Next: &leetcode.ListNode{
				Val:  9,   // quarto nó
				Next: nil, // fim da lista
			},
		},
	},
}

func main() {
	fns := leetcode.AddTwoNumbers(&l1, &l2)

	for name, fn := range fns {
		exec(name, fn)
	}
}

func exec(name string, fn func() any) {
	fmt.Println("\n" + name + ":")
	fmt.Println("Resultado:", fn())
}
