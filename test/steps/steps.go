// steps/login_steps.go
package steps

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"automation_appWeb_test/test/api"
)

type LoginSteps struct {
    Service      *api.APIService
    ResponseCode int
    ResponseBody map[string]interface{}
}

// Método correto, com a primeira letra maiúscula para ser exportado
func (l *LoginSteps) ILoginWithEmailAndPassword(t *testing.T, email, password string) error {
    // Define o corpo do login
    body := map[string]string{"email": email, "password": password}
    // Chama o método correto para login
    resp, err := l.Service.LoginUser(body) // Certifique-se de usar o método correto para o login
    if err != nil {
        return err
    }
    l.ResponseCode = resp.StatusCode()

    l.ResponseBody = make(map[string]interface{})
    if err := json.Unmarshal(resp.Body(), &l.ResponseBody); err != nil {
        return err
    }

    if l.ResponseCode >= 400 {
        if l.ResponseBody["error"] == nil {
            return errors.New("erro na resposta, mas campo 'error' está ausente")
        }
    }
    return nil
}

func TestLoginWithEmailAndPassword(t *testing.T) {
    service := api.NewAPIService()
    steps := &LoginSteps{Service: service}

    err := steps.ILoginWithEmailAndPassword(t, "eve.holt@reqres.in", "pistol")
    assert.NoError(t, err)
    assert.Equal(t, 200, steps.ResponseCode)

    assert.Contains(t, steps.ResponseBody, "id")
    assert.Contains(t, steps.ResponseBody, "email")
}