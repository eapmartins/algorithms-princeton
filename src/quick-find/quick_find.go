package quick_find

import "fmt"

type QuickFind struct {
	arr []int
}

func (qf *QuickFind) Union(p, q int) error {

	err := qf.validIndex(p, q)
	if err != nil {
		return err
	}

	pID := qf.arr[p]
	qID := qf.arr[q]

	for i := 0; i < len(qf.arr); i++ {
		if qf.arr[i] == pID {
			qf.arr[i] = qID
		}
	}

	return nil
}

func (qf *QuickFind) Connected(p, q int) (bool, error) {
	err := qf.validIndex(p, q)
	if err != nil {
		return false, err
	}

	return qf.arr[p] == qf.arr[q], nil
}

func (qf *QuickFind) validIndex(p, q int) error {

	if qf.arr == nil {
		return fmt.Errorf("arr should not be null")
	}

	if p > len(qf.arr) || q > len(qf.arr) || p <= 0 || q <= 0 {
		return fmt.Errorf("index out of range %d and %d", p, q)
	}

	return nil
}
