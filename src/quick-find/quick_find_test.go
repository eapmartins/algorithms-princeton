package quick_find

import "testing"

func TestQuickFindWithNotConnected(t *testing.T) {
	quickFind := QuickFind{arr: nil}
	quickFind.Connected(0, 0)
}
