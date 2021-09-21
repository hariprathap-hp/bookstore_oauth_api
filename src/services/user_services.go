package services

import (
	"fmt"
	"strings"

	"github.com/hari/bookstore_oauth_api/src/domain/token"
	"github.com/hari/bookstore_oauth_api/src/repo/dbRepo"
	"github.com/hari/bookstore_oauth_api/src/repo/userRepo"
	"github.com/hari/bookstore_users_api/utils/errors"
)

type UserService interface {
	GetToken(string) (*token.AccessToken, *errors.RestErr)
	CreateToken(token.AccessTokenRequest) (*token.AccessToken, *errors.RestErr)
	UpdateToken()
}

type userService struct {
	dbRepo       dbRepo.DbRepository
	restUserRepo userRepo.UserRepository
}

func NewService(dbrepo dbRepo.DbRepository, restRepo userRepo.UserRepository) UserService {
	return &userService{
		dbRepo:       dbrepo,
		restUserRepo: restRepo,
	}
}

func (s *userService) GetToken(at string) (*token.AccessToken, *errors.RestErr) {
	at = strings.TrimSpace(at)
	if len(at) == 0 {
		return nil, errors.NewHTTPBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetAccessToken(at)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *userService) CreateToken(request token.AccessTokenRequest) (*token.AccessToken, *errors.RestErr) {
	fmt.Println("Create Token")
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	//get New Access Token
	at := token.GetNewAccessToken(user.Id)
	at.Generate()
	createErr := s.dbRepo.CreateAccessToken(at)
	if createErr != nil {
		return nil, createErr
	}
	return &at, nil
}

func (s *userService) UpdateToken() {

}
