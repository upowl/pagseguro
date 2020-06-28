package sessions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upowl/pagseguro"
)

func TestPagseguroSessions(t *testing.T) {
	pagseguroInput := pagseguro.PagseguroInput{
		Env:   pagseguro.Sandbox,
		Email: os.Getenv("PAGSEGURO_EMAIL"),
		Token: os.Getenv("PAGSEGURO_TOKEN"),
	}

	pag, _ := pagseguro.New(&pagseguroInput)

	input := SessionsInput{
		Pagseguro: pag,
	}

	sess, errSess := New(&input)
	assert.Nil(t, errSess, "this return err is nil")

	token, errToken := sess.GetToken()
	assert.Error(t, errToken, "this return err is nil")
	typeString := ""
	assert.IsType(t, &typeString, token, "this return is type string pointer")
}
