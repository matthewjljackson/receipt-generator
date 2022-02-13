package receiptGenerator

import (
	"testing"
)

func TestNewBill(t *testing.T) {
	want := "Barry's Bill"
	testBill := newBill(want)
	got := testBill.name
	if got != want {
		t.Error("error", got, want)
	}
}

func BenchmarkNewBill(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newBill("bill name")
	}
}
