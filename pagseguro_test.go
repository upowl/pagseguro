package pagseguro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagseguroProduction(t *testing.T) {
	input := PagseguroInput{
		Env:   Production,
		Email: "email@email.com.br",
		Token: "98ouyhbouihb07ho8yb879b",
	}

	pag, errPag := New(&input)
	assert.Nil(t, errPag, "this return err is nil")
	assert.Equal(t, "email@email.com.br", pag.GetEmail(), "this return is email")
	assert.Equal(t, "98ouyhbouihb07ho8yb879b", pag.GetToken(), "this return is token")
	assert.Equal(t, productionURL, pag.GetURL(), "this return is url production")
}

func TestPagseguroSandbox(t *testing.T) {
	input := PagseguroInput{
		Env:   Sandbox,
		Email: "email@email.com.br",
		Token: "98ouyhbouihb07ho8yb879b",
	}

	pag, errPag := New(&input)
	assert.Nil(t, errPag, "this return err is nil")
	assert.Equal(t, "email@email.com.br", pag.GetEmail(), "this return is email")
	assert.Equal(t, "98ouyhbouihb07ho8yb879b", pag.GetToken(), "this return is token")
	assert.Equal(t, sandboxURL, pag.GetURL(), "this return is url sandbox")
}
