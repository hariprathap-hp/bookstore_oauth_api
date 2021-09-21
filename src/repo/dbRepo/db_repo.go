package dbRepo

import (
	"fmt"
	"strings"

	"github.com/hari/bookstore_oauth_api/src/dbPostgres"
	"github.com/hari/bookstore_oauth_api/src/domain/token"
	"github.com/hari/bookstore_users_api/utils/errors"
)

type DbRepository interface {
	GetAccessToken(string) (*token.AccessToken, *errors.RestErr)
	CreateAccessToken(token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

const (
	getTokenDBQuery  = "select access_token, user_id, client_id, expires from access_tokens where access_token=$1"
	createTokenQuery = "insert into access_tokens (access_token, user_id, client_id, expires) values ($1,$2,$3,$4)"
	updateTokenQuery = "update access_tokens set expires=$1 where access_token=$2"
)

func (db *dbRepository) GetAccessToken(at string) (*token.AccessToken, *errors.RestErr) {
	var result token.AccessToken
	stmt, err := dbPostgres.Client.Prepare(getTokenDBQuery)
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewHTTPInternalServerError("statement creation failed")
	}
	defer stmt.Close()
	getErr := stmt.QueryRow(at).Scan(&result.Access_Token, &result.UserId, &result.ClientId, &result.ExpiryTime)
	if getErr != nil {
		if strings.Contains(getErr.Error(), " no rows in result set") {
			return nil, errors.NewHTTPNotFoundError("enter valid access token")
		}
		return nil, errors.NewHTTPInternalServerError("error while parsing the result")
	}
	return &result, nil
}

func (db *dbRepository) CreateAccessToken(at token.AccessToken) *errors.RestErr {
	stmt, err := dbPostgres.Client.Prepare(createTokenQuery)
	if err != nil {
		return errors.NewHTTPInternalServerError("statement creation failed")
	}
	defer stmt.Close()
	_, createErr := stmt.Exec(at.Access_Token, at.UserId, at.ClientId, at.ExpiryTime)
	if createErr != nil {
		return errors.NewHTTPInternalServerError("creation of access token in DB failed")
	}
	return nil
}

func (db *dbRepository) UpdateExpirationTime(at token.AccessToken) *errors.RestErr {
	stmt, err := dbPostgres.Client.Prepare(updateTokenQuery)
	if err != nil {
		return errors.NewHTTPInternalServerError("statement creation failed")
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(at.ExpiryTime, at.Access_Token)
	if updateErr != nil {
		return errors.NewHTTPInternalServerError("creation of access token in DB failed")
	}
	return nil
}
