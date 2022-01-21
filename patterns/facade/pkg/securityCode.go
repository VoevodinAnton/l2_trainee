package pkg

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("Security Code verified")
	return nil
}
