package quick_find

import "fmt"

type QuickFind struct {
	arr []int
}

func (qf *QuickFind) Union(p, q int) {

	qf.validIndex(p, q)

	pID := qf.arr[p]
	qID := qf.arr[q]

	for i := 0; i < len(qf.arr); i++ {
		if qf.arr[i] == pID {
			qf.arr[i] = qID
		}
	}
}

func (qf *QuickFind) Connected(p, q int) bool {
	return qf.arr[p] == qf.arr[q]
}

func (qf *QuickFind) validIndex(p, q int) {
	if p < len(qf.arr) && q < len(qf.arr) && p >= 0 && q >= 0 {
		return
	}
	panic(fmt.Sprintf("index out of range %d and %d", p, q))
}

func (qf *QuickFind) Print() {
	for i := 0; i < len(qf.arr); i++ {
		fmt.Printf("%d:%d ", i, qf.arr[i])
	}
	println()
}
