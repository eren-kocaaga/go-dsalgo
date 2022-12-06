package facade

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("security code is incorrect")
	}

	fmt.Println("security code verified")

	return nil
}
