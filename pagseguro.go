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
	GetSession()
	SendPayment()
	CancelPayment()
	ReversePayment()
	ConsultNotification()
	ConsultPayment()
	ConsultPaymentDetails()
}

type pagseguro struct{}

func (item *pagseguro) GetSession()            {}
func (item *pagseguro) SendPayment()           {}
func (item *pagseguro) CancelPayment()         {}
func (item *pagseguro) ReversePayment()        {}
func (item *pagseguro) ConsultNotification()   {}
func (item *pagseguro) ConsultPayment()        {}
func (item *pagseguro) ConsultPaymentDetails() {}

// PagseguroInput PagseguroInput
type PagseguroInput struct {
	Env   environment
	Email string
	Token string
}

// New New
func New(input *PagseguroInput) (Pagseguro, error) {
	item := pagseguro{}
	return &item, nil
}
