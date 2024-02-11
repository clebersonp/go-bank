package main

// to import my packages: full path (module name + directories name).

import (
	"example.com/bank/fileops"
	"fmt"
)

const fileName = "balance.txt"

func main() {
	accountBalance, err := fileops.readFloatFromFile(fileName, 0)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
		// panic crash the application with stack error
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	// infinity loop
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			writeFloatToFile(accountBalance, fileName)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			writeFloatToFile(accountBalance, fileName)
		case 4:
			fmt.Println("Your balance:", accountBalance)
		default:
			fmt.Println("Your balance:", accountBalance)
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
