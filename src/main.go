package main

import (
	"awesomeProject/src/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("dongeon", 1000 )

	account.Deposit(80)
	fmt.Println(account.Balance())

	err := account.Withdraw(10000)
	if err != nil {
		fmt.Println(err)
	}
}
