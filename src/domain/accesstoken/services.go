package accesstoken

import (
	"strings"

	"github.com/yesseneon/bookstore_oauth_api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RESTError)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RESTError)
}

type service struct {
	repository Repository
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RESTError) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.BadRequest("Invalid access token id")
	}

	accessToken, restErr := s.repository.GetByID(accessTokenID)
	if restErr != nil {
		return nil, restErr
	}

	return accessToken, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
