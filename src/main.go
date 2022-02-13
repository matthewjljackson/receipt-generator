package receiptGenerator

import (
	"bufio"
	"fmt"
	"os"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	newBill := newBill(name)
	fmt.Printf("New bill created under the name %v \n", newBill.name)
	return newBill
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
