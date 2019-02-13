package entity

import (
	uiza "api-wrapper-go"
	"errors"
)

func CheckCreateParamsIsValid(params *uiza.EntityCreateParams) error {
	var err error
	if params.URL == nil || params.InputType == nil || params.Name == nil {
		return errors.New("Missing params")
	}
	if (params.URL == nil || *params.URL == "") && *params.InputType != uiza.InputTypeS3UIZA {
		return errors.New("Wrong params URL")
	}
	return err
}
