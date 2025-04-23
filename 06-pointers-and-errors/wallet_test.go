package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want BitCoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := BitCoin(10)

		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		wallet.Withdraw(5)

		want := BitCoin(5)

		assertBalance(t, wallet, want)
	})
}
