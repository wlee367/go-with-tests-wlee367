package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

/**
* In Go if a symbol (variables, types, functions et al) starts
* with a lowercase symbol then it is private
* outside the package it's defined in.*
 */
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/**
so the language permits us to write w.balance, without an explicit dereference. These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.
Technically you do not need to change Balance to use a pointer receiver as taking a copy of the balance is fine. However, by convention you should keep your method receiver types the same for consistency.
*/
