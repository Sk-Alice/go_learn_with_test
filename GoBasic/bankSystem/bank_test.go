package bankSystem

// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/pointers-and-errors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		_ = wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	// 测试--透支操作，不进行任何操作
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Error("didn't get an error but wanted one")
	}

	if !errors.Is(want, got) {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
