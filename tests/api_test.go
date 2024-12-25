package tests

import (
	"automation_test_golang/internal/api"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	service := api.NewAPIService("https://reqres.in/api")

	resp, err := service.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var body map[string]interface{}
	json.Unmarshal(resp.Body(), &body)
	assert.NotNil(t, body["data"])
}

func TestCreateUser(t *testing.T) {
	service := api.NewAPIService("https://reqres.in/api")

	payload := map[string]string{
		"name": "Test User",
		"job":  "Engineer",
	}

	resp, err := service.CreateUser(payload)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())

	var body map[string]interface{}
	json.Unmarshal(resp.Body(), &body)
	assert.Contains(t, body, "name")
	assert.Contains(t, body, "job")
	assert.Equal(t, "Test User", body["name"])
	assert.Equal(t, "Engineer", body["job"])
}
