package receiptGenerator

import (
	"testing"
)

func TestNewBill(t *testing.T) {
	want := "Barry's Bill"
	testBill := newBill(want)
	got := testBill.name
	if got != want {
		t.Errorf("error! got %s and expected %s", got, want)
	}
}

func TestAddItem(t *testing.T) {
	billName := "Barry's Bill"
	testBill := newBill(billName)
	testItem := "food"
	testItemPrice := 19.99
	testBill.addItem(testItem, testItemPrice)
	if testBill.items[testItem] != testItemPrice {
		t.Errorf("error! item %s costing %0.2f was not added to the bill", testItem, testItemPrice)
	}
}

func TestAddTip(t *testing.T) {
	billName := "Barry's Bill"
	testBill := newBill(billName)
	var testTipAmount float64 = 100
	testBill.updateTip(testTipAmount)
	if testBill.tip != testTipAmount {
		t.Errorf("error! the tip amount was not %f", testTipAmount)
	}
}
func BenchmarkNewBill(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newBill("bill name")
	}
}
