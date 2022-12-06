package facade

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("wallet balance added successfully")
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("wallet balance is not sufficient")
	}

	fmt.Println("wallet balance is sufficient")
	w.balance -= amount

	return nil
}
