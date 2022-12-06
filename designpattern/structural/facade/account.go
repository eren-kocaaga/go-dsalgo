package facade

import "fmt"

type Account struct {
	name string
}

func newAccount(name string) *Account {
	return &Account{name: name}
}

func (a *Account) checkAccount(name string) error {
	if a.name != name {
		return fmt.Errorf("account name is incorrect")
	}

	fmt.Println("account verified")
	return nil
}
