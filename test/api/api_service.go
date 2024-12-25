package api

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"

	"automation_appWeb_test/test/config"
)

type APIService struct {
    Client  *resty.Client
    BaseURL string
}

func NewAPIService() *APIService {
    client := resty.New()
    return &APIService{
        Client:  client,
        BaseURL: config.BaseURL,
    }
}

func (s *APIService) GetAllUsers() (*resty.Response, error) {
    resp, err := s.Client.R().Get(s.BaseURL + config.UsersPath)
    fmt.Println("Request URL:", s.BaseURL+config.UsersPath)
    fmt.Println("Response Status Code:", resp.StatusCode())
    fmt.Println("Response Body:", string(resp.Body()))
    return resp, err
}

func (s *APIService) GetSingleUser(userID int) (*resty.Response, error) {
    url := s.BaseURL + config.SingleUserPath + strconv.Itoa(userID)
    resp, err := s.Client.R().Get(url)
    fmt.Println("Request URL:", url)
    fmt.Println("Response Status Code:", resp.StatusCode())
    fmt.Println("Response Body:", string(resp.Body()))
    return resp, err
}

func (s *APIService) CreateUser(body interface{}) (*resty.Response, error) {
    fmt.Println("Request Body:", body)
    resp, err := s.Client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Post(s.BaseURL + config.UsersPath)
    fmt.Println("Request URL:", s.BaseURL+config.UsersPath)
    fmt.Println("Response Status Code:", resp.StatusCode())
    fmt.Println("Response Body:", string(resp.Body()))
    return resp, err
}

// Adicione o método LoginUser para testar a funcionalidade de login
func (s *APIService) LoginUser(body interface{}) (*resty.Response, error) {
    fmt.Println("Request Body:", body) // Exibe o corpo da requisição
    resp, err := s.Client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Post(s.BaseURL + config.LoginPath)
    fmt.Println("Request URL:", s.BaseURL+config.LoginPath)
    fmt.Println("Response Status Code:", resp.StatusCode())
    fmt.Println("Response Body:", string(resp.Body())) // Exibe o corpo da resposta
    return resp, err
}
