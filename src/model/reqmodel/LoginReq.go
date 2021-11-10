package reqmodel

import "github.com/go-playground/validator"

type UserReq struct {
	Email          string    `json:"email" validate:"required,email"`
	Password       string    `json:"password" validate:"required"`
}

type UserValidator struct {
	Validator *validator.Validate
}

func (uv *UserValidator) Validate(i interface{}) error {
	return uv.Validator.Struct(i)
}