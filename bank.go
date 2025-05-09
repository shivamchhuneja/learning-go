package main

import (
	"fmt"
	"os"
)

func bankingMode() {
	bankService(0.0)
}

func bankService(balance float64) {
	var choice int

	fmt.Println("\nWelcome to Go Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		checkBalance(balance)
		bankService(balance)
	case 2:
		balance = depositMoney(balance)
		bankService(balance)
	case 3:
		balance = withdrawMoney(balance)
		bankService(balance)
	case 4:
		fmt.Println("Thanks for banking with us!")
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Try again.")
		bankService(balance)
	}
}

func checkBalance(balance float64) {
	fmt.Printf("Current Balance: %.2f\n", balance)
}

func depositMoney(balance float64) float64 {
	var amount float64
	fmt.Printf("Your current balance is %.2f\nHow much would you like to deposit? ", balance)
	fmt.Scan(&amount)
	balance += amount
	fmt.Printf("Done. New Balance is %.2f\n", balance)
	return balance
}

func withdrawMoney(balance float64) float64 {
	var amount float64
	fmt.Printf("Your current balance is %.2f\nHow much would you like to withdraw? ", balance)
	fmt.Scan(&amount)

	if amount > balance {
		fmt.Println("❌ Insufficient funds. Try a smaller amount.")
	} else {
		balance -= amount
		fmt.Printf("Withdrawal successful. New Balance is %.2f\n", balance)
	}

	return balance
}
