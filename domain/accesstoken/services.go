package accesstoken

import (
	"strings"

	"github.com/yesseneon/bookstore_oauth_api/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RESTError)
	Create(AccessToken) *errors.RESTError
	UpdateExpirationTime(AccessToken) *errors.RESTError
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RESTError)
	Create(AccessToken) *errors.RESTError
	UpdateExpirationTime(AccessToken) *errors.RESTError
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
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

func (s *service) Create(at AccessToken) *errors.RESTError {
	if restErr := at.Validate(); restErr != nil {
		return restErr
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RESTError {
	if restErr := at.Validate(); restErr != nil {
		return restErr
	}

	return s.repository.UpdateExpirationTime(at)
}
