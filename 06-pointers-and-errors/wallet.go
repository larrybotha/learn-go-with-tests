package wallet

import (
	"errors"
	"fmt"
)

/*
Type aliases can be defined using the following syntax:
type [TypeName] [originalType]
*/
type Bitcoin int

/*
Any value that has a String() method is implementing the method defined on the
Stringer interface

String() defines the default way a value is printed

When printing formattd strings, we'll need to use %s for String() to print our
custom representation of the value
*/
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

/*
this type of receiver is called a value received
*/
func (w Wallet) DepositNoPointer(value Bitcoin) {
	/*
		The address here is different from the address inside the test, because
		we are not referencing the same value here, but instead a copy of the
		value where this function was called

		If we modify the copy, as we are doing in this functionn, the actual value
		of the instance property remains unchanged
	*/
	fmt.Printf("address of wallet in DepositNoPointer func is %#v\n", &w.balance)

	w.balance += value
}

/*
This method modifies an instance property, so the received type must
use a pointer

This type of receiver is called a pointer receiver
*/
func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value

	// also equivalent to
	//(*w).balance += value
}

/*
As a convention in Go, one should keep method receives consistent

This receiver does not need to be a pointer receiver, but because we
have a pointer receiver on Deposit, we make this a pointer receiver, too
*/
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount

	return nil
}
