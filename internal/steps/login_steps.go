package steps

import (
	"automation_test_golang/internal/api"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LoginSteps struct {
	Service      *api.APIService
	ResponseCode int
	ResponseBody map[string]interface{}
}

func (l *LoginSteps) iLoginWithEmailAndPassword(t *testing.T, email, password string) error {
	body := map[string]string{"email": email, "password": password}
	resp, err := l.Service.LoginUser(body)
	if err != nil {
		return err
	}
	l.ResponseCode = resp.StatusCode()

	// Processar a resposta JSON
	l.ResponseBody = make(map[string]interface{})
	if err := json.Unmarshal(resp.Body(), &l.ResponseBody); err != nil {
		return err
	}

	// Validar o campo 'error' em caso de falha
	if l.ResponseCode >= 400 {
		if l.ResponseBody["error"] == nil {
			return errors.New("erro na resposta, mas campo 'error' está ausente")
		}
	}
	return nil
}

func TestLoginWithEmailAndPassword(t *testing.T) {
	service := api.NewAPIService("https://reqres.in/api")
	steps := &LoginSteps{Service: service}

	// Simulando o login
	err := steps.iLoginWithEmailAndPassword(t, "eve.holt@reqres.in", "pistol")
	assert.NoError(t, err)
	assert.Equal(t, 200, steps.ResponseCode)

	// Validando a resposta JSON
	assert.Contains(t, steps.ResponseBody, "id")
	assert.Contains(t, steps.ResponseBody, "email")
}
