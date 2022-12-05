package main

import (
	"fmt"
	"dsalgo/cord"
	"dsalgo/disjoint"
)

func main() {
	TestDisjoinSet()
}

// Suppose you are given a grid of 1's and 0's. All adjacent 1's are connected components.
// For example, in the following case you have 2 connected components because you have two "islands" of 1's.
//
// 1 1 0 0 1 1
// 1 0 0 0 1 1 
// 1 0 0 0 0 0
//
func TestDisjoinSet() {
	fmt.Println("TestDisjoinSet")
	grid := [][]int{
		[]int{1,1,0,0,1,1},
		[]int{1,0,0,0,1,1},
		[]int{1,1,0,0,0,0},
		[]int{1,0,0,0,0,0},
		[]int{1,0,0,0,0,0},
		[]int{1,0,0,0,0,0},
	}
	n := len(grid)
	ds := disjoint.New(n)
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if i!=j && grid[i][j] == 1 {
				ds.Union(i, j)
			}
		}
	}
	fmt.Println(ds.Components())
	ds.Union(0,3)
	fmt.Println(ds.Components())
	ds.Union(0,2)
	fmt.Println(ds.Components())
}

func TestCord() {
	fmt.Println("Cord Tree")
	tree := cord.Tree{}
	root := cord.InternalNode {
		Len: 26,
		Left: cord.LeafNode {
			Len: 10,
			Value: []rune("abcdefghij"),
		},
		Right: cord.InternalNode {
			Len: 16,
			Left: cord.LeafNode {
				Len: 6,
				Value: []rune("klmnop"),
			},
			Right: cord.LeafNode {
				Len: 10,
				Value: []rune("qrstuvwxyz"),
			},
		},
	}
	tree.New(root)
	_, i2 := tree.FindIndex(2)
	_, i11 := tree.FindIndex(11)
	_, i22 := tree.FindIndex(22)
	_, i25 := tree.FindIndex(25)
	_, i26 := tree.FindIndex(26)
	fmt.Printf("2: %s, 11: %s, 22: %s, 25: %s, 26: %s", string(i2), string(i11), string(i22), string(i25), string(i26))
}