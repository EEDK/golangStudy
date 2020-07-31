package banking

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("Your account have few money to handle this activity")

// BankAccount struct
type BankAccount struct{
	owner string
	balance int
}

// NewAccount is Constructor
func NewAccount(owner string) *BankAccount {
	account := BankAccount{owner: owner , balance: 0}
	return &account
}

// Getter
func (a BankAccount) Balance() int {
	return a.balance
}

func (a BankAccount) Owner() string{
	return a.owner
}

// Activity
func (a *BankAccount) AccountDeposit(amount int){
	a.balance += amount
}

func (a BankAccount) String() string{
	return fmt.Sprint(a.Owner(), "'s account.\nHas ", a.Balance())
}

func (a *BankAccount) AccountWithdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

func (a *BankAccount) ChangeAccountOwner(newOwner string)  {
	a.owner = newOwner
}