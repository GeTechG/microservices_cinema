package user

import (
	"crypto/sha256"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	PasswordHash [32]byte  `json:"-"`
	DateBirthday time.Time `json:"date_birthday"`
}

func NewUser(login, password string, dateBirthday time.Time) User {
	return User{
		Id:           uuid.New(),
		Login:        login,
		PasswordHash: sha256.Sum256([]byte(password)),
		DateBirthday: dateBirthday,
	}
}
