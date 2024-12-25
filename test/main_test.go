// test/main_test.go
package main

import (
	"testing"

	"automation_appWeb_test/test/api" // Corrija a importação do pacote api
	"automation_appWeb_test/test/steps"
)

func TestLogin(t *testing.T) {
    service := api.NewAPIService()  // Use a função NewAPIService corretamente
    steps := &steps.LoginSteps{Service: service}

    err := steps.iLoginWithEmailAndPassword(t, "eve.holt@reqres.in", "pistol")
    if err != nil {
        t.Fatalf("Teste falhou: %s", err)
    }
}
