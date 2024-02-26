package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7",randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return
				continue
			}
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than your have.")
				return
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Bye!")
			fmt.Println("Thank for choosing our bank")
			return
			//break
		}
	}
}

