package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			// use %s so the String() prints our custom string value
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("expecting error but got nil")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got.Error(), want)
		}
	}

	t.Run("Deposit without pointers", func(t *testing.T) {
		wallet := Wallet{}

		// print with %#v - use Go's Stringer so that we can see the address without
		// our own String() implementation on Bitcoin overriding it
		fmt.Printf("address of wallet in DepositNoPointer test is %#v\n", &wallet.balance)

		wallet.DepositNoPointer(Bitcoin(10))

		want := Bitcoin(0)

		assertBalance(t, wallet, want)
	})

	t.Run("Deposit with pointers", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}

		wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(5)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}
