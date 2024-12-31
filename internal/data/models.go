package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrEditConflict       = errors.New("edit conflict")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
)

type Models struct {
	Cars CarModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Cars: CarModel{DB: db},
	}
}
