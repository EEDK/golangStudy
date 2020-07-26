package main

import (
	"fmt"
	"golangStudy/banking"
	"log"
)

func main() {
	account := banking.NewAccount("kde")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil{
		log.Fatalln(err)
	}
}
