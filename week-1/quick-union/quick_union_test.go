package quick_union

import (
	"fmt"
	"testing"
)

func TestQuickUnionFindRootOfOneLevelTree(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, sz: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	index := qu.Root(3)
	if index != 3 {
		t.Errorf("index should be 3")
	}
}

func TestQuickUnionFindRootOfTwoLevelTree(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 4, 4, 5, 6, 7, 8, 9}, sz: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	index := qu.Root(3)

	if index != 4 {
		t.Errorf("index should be 4")
	}
}

func TestQuickUnionFindRootOfThreeLevelsTree(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 4, 8, 5, 6, 7, 8, 9}, sz: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	index := qu.Root(3)

	if index != 8 {
		t.Errorf("index should be 8")
	}
}

func TestQuickUnionConnected(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, sz: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	connected, err := qu.Connected(3, 4)
	if connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, false)
	}

	err = qu.Union(3, 4)
	connected, err = qu.Connected(3, 4)
	if !connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, true)
	}

	err = qu.Union(4, 8)
	connected, err = qu.Connected(3, 8)
	if !connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, true)
	}

	err = qu.Union(8, 9)
	connected, err = qu.Connected(3, 9)
	if !connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, true)
	}

	err = qu.Union(9, 3)
	connected, err = qu.Connected(3, 9)
	if !connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, true)
	}

	connected, err = qu.Connected(1, 2)
	if connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, false)
	}
}

func TestQuickUnionRootWithCompression(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, sz: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}

	printResult(qu.arr)
	_ = qu.Union(4, 3)
	_ = qu.Union(3, 8)
	_ = qu.Union(6, 5)
	_ = qu.Union(9, 4)
	_ = qu.Union(2, 1)
	_ = qu.Union(5, 0)
	_ = qu.Union(7, 2)
	_ = qu.Union(6, 1)
	_ = qu.Union(7, 3)
	printResult(qu.arr)

	index := qu.Root(3)
	if index != 6 {
		t.Errorf("index should be 6")
	}

}

func printResult(arr []int) {
	fmt.Printf("%v", arr)
	println()
}
