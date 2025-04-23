package wallet

import "fmt"

type BitCoin int

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(value BitCoin) {
	w.balance += value
}

func (w *Wallet) Withdraw(value BitCoin) {
	w.balance -= value
}

// keep method receivers the same for consistency
func (w *Wallet) Balance() BitCoin {
	// return (*w).balance // struct pointer is automatically dereferenced
	return w.balance
}
