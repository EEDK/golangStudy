package banking

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

func (a BankAccount) Balance() int {
	return a.balance
}

// Get Actual account Change 
func (a *BankAccount) Deposit(amount int){
	a.balance += amount
}