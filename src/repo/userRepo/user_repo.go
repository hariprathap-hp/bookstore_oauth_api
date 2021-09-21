package userRepo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hariprathap-hp/bookstore_oauth_api/src/domain/userLogin"
	"github.com/hariprathap-hp/bookstore_users_api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	baseURL = "http://localhost:8080"
	timeout = time.Millisecond * 100
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: baseURL,
		Timeout: timeout,
	}
)

type UserRepository interface {
	LoginUser(string, string) (*userLogin.User, *errors.RestErr)
}

type userRepository struct {
}

func NewUserRepo() UserRepository {
	return &userRepository{}
}

func (u *userRepository) LoginUser(email, password string) (*userLogin.User, *errors.RestErr) {
	postBody := userLogin.LoginUser{
		Email:    email,
		Password: password,
	}
	fmt.Println("PostBody", postBody)
	response := usersRestClient.Post("/users/login", postBody)
	if response == nil {
		return nil, errors.NewHTTPInternalServerError("invalid restclient response when trying to login user")
	}
	fmt.Println(response)
	var user userLogin.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewHTTPInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
