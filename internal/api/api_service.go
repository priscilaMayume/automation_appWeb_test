package api

import (
	"strconv"

	"github.com/go-resty/resty/v2"
)

type APIService struct {
	Client  *resty.Client
	BaseURL string
}

func NewAPIService(baseURL string) *APIService {
	client := resty.New()
	return &APIService{
		Client:  client,
		BaseURL: baseURL,
	}
}

func (s *APIService) GetAllUsers() (*resty.Response, error) {
	return s.Client.R().Get(s.BaseURL + "/users")
}

func (s *APIService) GetSingleUser(userID int) (*resty.Response, error) {
	return s.Client.R().Get(s.BaseURL + "/users/" + strconv.Itoa(userID))
}

func (s *APIService) CreateUser(body interface{}) (*resty.Response, error) {
	return s.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.BaseURL + "/users")
}

func (s *APIService) LoginUser(body interface{}) (*resty.Response, error) {
	return s.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.BaseURL + "/login")
}
