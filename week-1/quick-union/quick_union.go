package quick_union

import "fmt"

type QuickUnion struct {
	arr []int
}

func (qu *QuickUnion) Union(p, q int) error {
	err := qu.validIndex(p, q)
	if err != nil {
		return err
	}

	pID := qu.Root(p)
	qID := qu.Root(q)

	qu.arr[pID] = qID

	return nil
}

func (qu *QuickUnion) Root(n int) int {

	for qu.arr[n] != n {
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

	if p > len(qu.arr) || q > len(qu.arr) || p <= 0 || q <= 0 {
		return fmt.Errorf("index out of range %d and %d", p, q)
	}

	return nil
}
