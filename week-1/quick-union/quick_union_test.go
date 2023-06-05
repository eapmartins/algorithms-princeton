package quick_union

import "testing"

func TestQuickUnionFindRootOfOneLevelTree(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	index := qu.Root(3)
	if index != 3 {
		t.Errorf("index should be 3")
	}
}

func TestQuickUnionFindRootOfTwoLevelTree(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 4, 4, 5, 6, 7, 8, 9}}
	index := qu.Root(3)

	if index != 4 {
		t.Errorf("index should be 4")
	}
}

func TestQuickUnionFindRootOfThreeLevelsTree(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 4, 8, 5, 6, 7, 8, 9}}
	index := qu.Root(3)

	if index != 8 {
		t.Errorf("index should be 8")
	}
}

func TestQuickUnionConnected(t *testing.T) {
	qu := QuickUnion{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
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
}
