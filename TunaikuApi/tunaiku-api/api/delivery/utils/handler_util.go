package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	. "tunaiku-api/api/utils"
	. "tunaiku-api/struct"
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func IsLoaningRequestValid(loaning *Loaning) (bool, error) {
	validate := validator.New()

	err := validate.Struct(loaning)
	if err != nil {
		return false, err
	}
	return true, nil
}
func IsRequestForLoaningValid(p *Loaning) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}