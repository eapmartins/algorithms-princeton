package quick_union

import "fmt"

type QuickUnion struct {
	arr []int
	sz  []int
}

func (qu *QuickUnion) Union(p, q int) error {
	err := qu.validIndex(p, q)
	if err != nil {
		return err
	}

	pID := qu.Root(p)
	qID := qu.Root(q)

	if pID == qID {
		return nil
	}

	if qu.sz[pID] < qu.sz[qID] {
		qu.arr[pID] = qID
		qu.sz[qID] += qu.sz[pID]
	} else {
		qu.arr[qID] = pID
		qu.sz[pID] += qu.sz[qID]
	}

	return nil
}

func (qu *QuickUnion) Root(n int) int {

	for qu.arr[n] != n {
		qu.arr[n] = qu.arr[qu.arr[n]] // path compression by halving the path length to root
		n = qu.arr[n]
	}
	return n
}

func (qu *QuickUnion) Connected(p, q int) (bool, error) {
	err := qu.validIndex(p, q)
	if err != nil {
		return false, err
	}

	return qu.Root(p) == qu.Root(q), nil
}

func (qu *QuickUnion) validIndex(p, q int) error {

	if qu.arr == nil {
		return fmt.Errorf("arr should not be null")
	}

	if p > len(qu.arr) || q > len(qu.arr) || p < 0 || q < 0 {
		return fmt.Errorf("index out of range %d and %d", p, q)
	}

	return nil
}
