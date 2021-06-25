package authorization_test

import (
	"testing"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
)

func TestSendEmail(t *testing.T) {
	err := authorization.LoadMail("../cmd/certificates/email.json")
	if err != nil {
		t.Error("Error in loadMail")
	}
	err = authorization.ConfirmEmail("wegel63355@paseacuba.com", "EstoEsUnTokenDePrueba8")
	if err != nil {
		t.Errorf("Error in confirmEmail %v", err)
	}
}
