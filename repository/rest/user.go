package rest

import (
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/yesseneon/bookstore_oauth_api/domain/user"
	"github.com/yesseneon/bookstore_utils/errors"
)

type RESTUserRepository interface {
	LoginUser(string, string) (*user.User, *errors.RESTError)
}

type userRepository struct{}

func NewRepository() RESTUserRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string, password string) (*user.User, *errors.RESTError) {
	data := user.LoginData{
		Email:    email,
		Password: password,
	}

	var u *user.User
	var restErr *errors.RESTError
	resp, err := resty.New().R().
		SetBody(data).
		SetResult(u).
		SetError(restErr).
		Post("http://localhost:8081/users/login")

	log.Println(resp)

	if err != nil {
		log.Println(err)
		log.Println(restErr)
		return nil, restErr
	}

	return u, nil
}
