package disjoint

import "fmt"

type DisjointSet struct {
	parent map[int]int
	rank   []int
}

func New(size int) DisjointSet {
	return DisjointSet {
		parent: make(map[int]int),
		rank: make([]int, size),
	}
}

func (d *DisjointSet) Union(x, y int) {
	px, py := d.Find(x), d.Find(y)
	if d.rank[py] > d.rank[px] {
		delete(d.parent, px)
		d.parent[px], d.rank[py] = py, d.rank[py]+1
	} else {
		delete(d.parent, py)
		d.parent[py], d.rank[px] = px, d.rank[px]+1
	}
	fmt.Println("Added: (", x, y, "), Parents:", d.parent, d.rank)
}

func (d *DisjointSet) Find(x int) int {
	for {
		fmt.Println("Finding:", x)
		p, e := d.parent[x]
		if !e || p == x { break }
		x = p
	}
	return x
} 

func (d *DisjointSet) Components() int {
	return len(d.parent)
}