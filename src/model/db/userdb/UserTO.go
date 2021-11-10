package userdb

import (
	"github.com/satori/go.uuid"
	)

type UserTo struct {
	ID             uuid.UUID
	Email          string
	EmailConfirmed bool
	Password       string
}
