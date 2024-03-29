package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func readBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("failed to find balance file")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
	fmt.Println("Balance updated! New amount:", balanceText)
}

func main() {
	accountBalance, err := readBalanceFromFile(fileName)

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
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Your balance")
		fmt.Println("5. Exit")

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
			writeBalanceToFile(accountBalance, fileName)
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
			writeBalanceToFile(accountBalance, fileName)
		case 4:
			fmt.Println("Your balance:", accountBalance)
		default:
			fmt.Println("Your balance:", accountBalance)
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
