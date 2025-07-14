package common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ParseErr(errs error) (status int, err error) {
	if errs == nil {
		return http.StatusOK, nil
	}

	newErrMes := ""
	status = 0

	var ve validator.ValidationErrors
	if errors.As(errs, &ve) {
		status = http.StatusBadRequest
		for _, v := range ve {
			if v.Tag() == "required" {
				newErrMes += fmt.Sprintf("Field %s must be provided;", v.Field())
			}
			if v.Tag() == "email" {
				newErrMes += fmt.Sprintf("Field %s must contains email;", v.Field())
			}
			if v.Tag() == "min" {
				newErrMes += fmt.Sprintf("Minimal lenght for field %s is %v;", v.Field(), v.Param())
			}
			if v.Tag() == "max" {
				newErrMes += fmt.Sprintf("Maximum lenght for field %s is %v;", v.Field(), v.Param())
			}
		}
	}

	switch {
	case errors.Is(errs, ErrSubscriptionNotFound):
		status = http.StatusNotFound
		newErrMes += fmt.Sprintf("%v;", ErrSubscriptionNotFound)
	case errors.Is(errs, ErrDateFormat):
		status = http.StatusBadRequest
		newErrMes += fmt.Sprintf("%v;", ErrDateFormat)
	case errors.Is(errs, ErrBeginDateAfterEndDate):
		status = http.StatusBadRequest
		newErrMes += fmt.Sprintf("%v;", ErrBeginDateAfterEndDate)
	case errors.Is(errs, ErrBadType):
		status = http.StatusBadRequest
		newErrMes += fmt.Sprintf("%v;", ErrBadType)
	case errors.Is(errs, ErrBadURL):
		status = http.StatusBadRequest
		newErrMes += fmt.Sprintf("%v;", ErrBadURL)
	case newErrMes == "":
		status = http.StatusInternalServerError
		newErrMes = "Internal server error"
	}

	return status, errors.New(newErrMes)
}
