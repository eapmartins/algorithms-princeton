package quick_find

import "testing"

func TestQuickFindNotConnected(t *testing.T) {
	quickFind := QuickFind{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	connected, err := quickFind.Connected(4, 8)

	if connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, false)
	}
}

func TestQuickFindConnected(t *testing.T) {
	quickFind := QuickFind{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	err := quickFind.Union(4, 3)
	connected, err := quickFind.Connected(4, 8)

	if connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, false)
	}

	err = quickFind.Union(3, 8)
	connected, err = quickFind.Connected(4, 8)

	if !connected && err == nil {
		t.Errorf("Output %t not equal to expected %t", connected, true)
	}
}
