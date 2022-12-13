package main

func main() {
	unionFind := QuickFindUnionFind{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	unionFind.Union(4, 3)
	unionFind.Union(3, 8)

	connected := unionFind.Connected(4, 8)
	println("4 and 8 are connected:", connected)

	connected = unionFind.Connected(4, 9)
	println("4 and 9 are connected:", connected)

	unionFind.Union(0, 1)
	connected = unionFind.Connected(0, 1)
	println("0 and 1 are connected:", connected)
}

type QuickFindUnionFind struct {
	arr []int
}

func (qf *QuickFindUnionFind) Union(p, q int) {
	pID := qf.arr[p]
	qID := qf.arr[q]

	for i := 0; i < len(qf.arr); i++ {
		if qf.arr[i] == pID {
			qf.arr[i] = qID
		}
	}
}

func (qf *QuickFindUnionFind) Connected(p, q int) bool {
	return qf.arr[p] == qf.arr[q]
}