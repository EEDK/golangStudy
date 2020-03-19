package accounts

import "errors"

// struct 작성할 go파일

//Account = 계좌 정보
type Account struct{
	owner string
	balance int
}

//NewAccount = create Count
func NewAccount(owner string , balance int) *Account{
	account := Account{owner: owner, balance: balance}
	return &account
}

// Deposit x 는 balance 변경해주는 메소드
func (theAccount *Account) Deposit(amount int){
	theAccount.balance += amount
}

// balnace만 출력하는 메소드
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x 만큼 제거
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Cant Withdraw 멍청아")
	}
	a.balance -= amount
	return nil
}