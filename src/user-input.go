package receiptGenerator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("What is the item name? ", reader)
		price, _ := getInput("What does the item cost? ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
		fmt.Println(b.format())
	case "t":
		tip, _ := getInput("How much would you like to tip? ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		promptOptions(b)

	default:
		fmt.Println("Please select a valid option")
		promptOptions(b)
	}
}
