package segment_tree

import (
	"fmt"
)

type node struct {
	val int
	left int
	right int
}

type SegmentTree struct {
	height   int
	data     []int
	capacity int
	tree     []node
}

func(s *SegmentTree) Update(index int, e int) {

} 

func(s *SegmentTree) RangeSum(left, right int) int {

}

func(s *SegmentTree) updateSG(e int, left, right int) {
	if left == right {

	}
}
