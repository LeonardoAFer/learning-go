package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileBalance string = "banana.txt"

func main() {
	var accountBalance, err = getBalanceFromFile(fileBalance)

	if err != nil {
		fmt.Println(err)
		return
		//panic(err)
	}

	for {
		fmt.Println("1 - Check balance")
		fmt.Println("2 - Deposit money")
		fmt.Println("3 - Withdraw money")
		fmt.Println("4 - Exit")

		var choice int
		fmt.Scan(choice)

		switch choice {
		case 1:
			fmt.Println(accountBalance)
			var deposit float64
			fmt.Scan(deposit)
			if deposit <= 0 {
				continue
			}
			accountBalance += deposit
		case 2:
			var deposit float64
			fmt.Scan(deposit)
			if deposit <= 0 {
				continue
			}
			accountBalance += deposit
			writeBalanceToFile(accountBalance)
		case 3:
			var withdraw float64
			fmt.Scan(withdraw)
			accountBalance -= withdraw
			writeBalanceToFile(accountBalance)
		case 4:
			return
		}
	}
}

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile("banana.txt", []byte(balanceString), 0644)
}

func getBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Error reading file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Println(err)
	}
	return balance, nil
}
