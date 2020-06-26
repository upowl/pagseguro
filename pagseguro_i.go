package pagseguro

type environment int
type urlEnv string

const (
	// Production Environment Production
	Production environment = 0
	// Sandbox Environment Sandbox
	Sandbox environment = 1

	productionURL urlEnv = "https://ws.pagseguro.uol.com.br/v2"
	sandboxURL    urlEnv = "https://ws.sandbox.pagseguro.uol.com.br/v2"
)

// Pagseguro Pagseguro
type Pagseguro interface {
	GetURL() string
	GetEmail() string
	GetToken() string
}

// New New
func New(input *PagseguroInput) (Pagseguro, error) {
	item := pagseguro{
		environment: input.Env,
		email:       input.Email,
		token:       input.Token,
	}

	return &item, nil
}
