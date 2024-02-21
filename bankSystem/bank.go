package bankSystem

import "fmt"

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

// Withdraw 取钱
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance 查看金额
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
