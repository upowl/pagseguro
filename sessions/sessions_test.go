package sessions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upowl/pagseguro"
)

func TestPagseguroSessions(t *testing.T) {
	pagseguroInput := pagseguro.PagseguroInput{
		Env:   pagseguro.Sandbox,
		Email: "email@email.com.br",
		Token: "aksdjlaskdjlkasjd",
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
