package wallet

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(value int) {
	w.balance = w.balance + value
}
func (w Wallet) Balance() int {
	return w.balance
}
