package sessions

import "github.com/upowl/pagseguro"

const (
	path = "/sessions"
)

// Sessions Sessions
type Sessions interface {
	GetToken() (*string, error)
}

type sessions struct {
	url   string
	email string
	token string
}

// SessionsInput SessionsInput
type SessionsInput struct {
	Pagseguro pagseguro.Pagseguro
}

// New New
func New(input *SessionsInput) (Sessions, error) {
	item := sessions{
		url:   input.Pagseguro.GetURL(),
		email: input.Pagseguro.GetEmail(),
		token: input.Pagseguro.GetToken(),
	}

	return &item, nil
}
