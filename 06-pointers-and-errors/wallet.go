package wallet

import (
	"errors"
	"fmt"
)

// BitCoin is a type for int used for the balance of a wallet
type BitCoin int

// ErrInsufficientFunds describes when a wallet is withdrawn from
// beyond its balance
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet describes a crypto wallet
type Wallet struct {
	balance BitCoin
}

// Deposit adds to a wallet's balance
func (w *Wallet) Deposit(value BitCoin) {
	w.balance += value
}

// Withdraw reduces the wallet balance by the given value
func (w *Wallet) Withdraw(value BitCoin) error {
	// valid but redundant to dereference on struct pointer
	if value > w.balance {
		return ErrInsufficientFunds
	}

	(*w).balance -= value

	return nil
}

// Balance returns the wallet's balance
// keep method receivers the same for consistency
func (w *Wallet) Balance() BitCoin {
	// return (*w).balance // struct pointer is automatically dereferenced
	return w.balance
}
