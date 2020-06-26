package pagseguro

import "fmt"

type pagseguro struct {
	environment environment
	email       string
	token       string
}

// GetURL GetURL
func (item *pagseguro) GetURL() string {
	switch item.environment {
	case Production:
		return fmt.Sprintf("%v", productionURL)
	case Sandbox:
		return fmt.Sprintf("%v", sandboxURL)
	default:
		return fmt.Sprintf("%v", sandboxURL)
	}
}

// GetEmail GetEmail
func (item *pagseguro) GetEmail() string {
	return item.email
}

// GetToken GetToken
func (item *pagseguro) GetToken() string {
	return item.token
}
