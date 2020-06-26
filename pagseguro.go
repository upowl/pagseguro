package pagseguro

type pagseguro struct {
	environment environment
	email       string
	token       string
}

// GetURL GetURL
func (item *pagseguro) GetURL() urlEnv {
	switch item.environment {
	case Production:
		return productionURL
	case Sandbox:
		return sandboxURL
	default:
		return sandboxURL
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
