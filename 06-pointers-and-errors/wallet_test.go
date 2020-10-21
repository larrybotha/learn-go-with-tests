package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("test without pointers", func(t *testing.T) {
		wallet := Wallet{}

		fmt.Printf("address of wallet in DepositNoPointer test is %v\n", &wallet.balance)

		wallet.DepositNoPointer(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(0)

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("test with pointers", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}
