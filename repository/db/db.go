package db

import (
	goerrors "errors"
	"log"

	"github.com/gocql/gocql"
	"github.com/yesseneon/bookstore-oauth-api/clients/cassandra"
	"github.com/yesseneon/bookstore-oauth-api/domain/accesstoken"
	"github.com/yesseneon/bookstore-utils/errors"
)

type DBRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RESTError)
	Create(accesstoken.AccessToken) *errors.RESTError
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RESTError
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RESTError) {
	var at accesstoken.AccessToken
	query := cassandra.GetSession().Query("SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?", id)
	err := query.Scan(&at.AccessToken, &at.UserID, &at.ClientID, &at.Expires)
	if err != nil {
		if goerrors.Is(err, gocql.ErrNotFound) {
			return nil, errors.NotFound()
		}

		return nil, errors.InternalServerError()
	}

	return &at, nil
}

func (r *dbRepository) Create(at accesstoken.AccessToken) *errors.RESTError {
	err := cassandra.GetSession().Query("INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)", at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec()
	if err != nil {
		log.Println(err)
		return errors.InternalServerError()
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RESTError {
	err := cassandra.GetSession().Query("UPDATE access_tokens SET expires=? WHERE access_token=?", at.Expires, at.AccessToken).Exec()
	if err != nil {
		return errors.InternalServerError()
	}

	return nil
}
