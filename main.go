package main

import (
	"fmt"
	"golangStudy/banking"
)

func main() {
	account := banking.NewAccount("kde")
	account.AccountDeposit(10)
	account.ChangeAccountOwner("Juwon")
	fmt.Println(account)
}
