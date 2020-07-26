package banking

import "errors"

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

func (a *BankAccount) Deposit(amount int){
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount int) (error) {
	if a.balance < amount {
		return errors.New("Can't withdraw your account")
	}

	a.balance -= amount
	return nil
}