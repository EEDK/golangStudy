package main

import (
	"fmt"
	"golangStudy/banking"
)

func main() {
	account := banking.NewAccount("kde")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
