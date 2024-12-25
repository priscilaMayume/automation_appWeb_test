package api

import (
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
    return s.Client.R().Get(s.BaseURL + config.UsersPath)
}

func (s *APIService) GetSingleUser(userID int) (*resty.Response, error) {
    return s.Client.R().Get(s.BaseURL + config.SingleUserPath + strconv.Itoa(userID))
}

func (s *APIService) CreateUser(body interface{}) (*resty.Response, error) {
    return s.Client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Post(s.BaseURL + config.UsersPath)
}

// Adicione o método LoginUser para testar a funcionalidade de login
func (s *APIService) LoginUser(body interface{}) (*resty.Response, error) {
    return s.Client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Post(s.BaseURL + config.LoginPath)
}
