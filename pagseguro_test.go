package pagseguro

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagseguroProduction(t *testing.T) {
	input := PagseguroInput{
		Env:   Production,
		Email: os.Getenv("PAGSEGURO_EMAIL"),
		Token: os.Getenv("PAGSEGURO_TOKEN"),
	}

	pag, errPag := New(&input)
	assert.Nil(t, errPag, "this return err is nil")
	assert.Equal(t, os.Getenv("PAGSEGURO_EMAIL"), pag.GetEmail(), "this return is email")
	assert.Equal(t, os.Getenv("PAGSEGURO_TOKEN"), pag.GetToken(), "this return is token")
	assert.Equal(t, fmt.Sprintf("%v", productionURL), pag.GetURL(), "this return is url production")
}

func TestPagseguroSandbox(t *testing.T) {
	input := PagseguroInput{
		Env:   Sandbox,
		Email: os.Getenv("PAGSEGURO_EMAIL"),
		Token: os.Getenv("PAGSEGURO_TOKEN"),
	}

	pag, errPag := New(&input)
	assert.Nil(t, errPag, "this return err is nil")
	assert.Equal(t, os.Getenv("PAGSEGURO_EMAIL"), pag.GetEmail(), "this return is email")
	assert.Equal(t, os.Getenv("PAGSEGURO_TOKEN"), pag.GetToken(), "this return is token")
	assert.Equal(t, fmt.Sprintf("%v", sandboxURL), pag.GetURL(), "this return is url sandbox")
}
