package bankSystem

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit 存钱
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// Withdraw 取钱
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

// Balance 查看金额
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
