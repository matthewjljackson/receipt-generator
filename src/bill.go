package receiptGenerator

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	bill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return bill
}

func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ...$%0.2f \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-15v ...$%0.2f \n", "tip:", b.tip)
	fs += fmt.Sprintf("%-15v ...$%0.2f", "total:", total+b.tip)
	return fs
}

func (b *bill) updateTip(tipAmount float64) {
	b.tip += tipAmount
}

func (b *bill) addItem(item string, price float64) {
	b.items[item] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Your bill was saved to file")
}
