package db

import (
	"github.com/yesseneon/bookstore_oauth_api/src/clients/cassandra"
	"github.com/yesseneon/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/yesseneon/bookstore_oauth_api/src/utils/errors"
)

type DBRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RESTError)
}

type dbRepository struct {
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RESTError) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.InternalServerError()
}

func NewRepository() DBRepository {
	return &dbRepository{}
}
