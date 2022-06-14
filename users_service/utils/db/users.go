package db

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	user "github.com/GeTechG/microservices_cinema/users_service/entities"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
)

var users map[uuid.UUID]user.User

func loadUsers() map[uuid.UUID]user.User {
	b, err := ioutil.ReadFile("users.json")
	if err != nil {
		return make(map[uuid.UUID]user.User)
	}
	var tempUsers map[uuid.UUID]user.User
	err = json.Unmarshal(b, &tempUsers)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return tempUsers
}

func AddNewUser(newUser user.User) error {
	for _, v := range users {
		if v.Login == newUser.Login {
			return errors.New("this is login exists")
		}
	}
	users[newUser.Id] = newUser
	return nil
}

func RemoveUser(uuid uuid.UUID) {
	delete(users, uuid)
}

func GetUserByUuid(uuid uuid.UUID) user.User {
	return users[uuid]
}

func UpdateUserByUuid(uuid uuid.UUID, user user.User) {
	_, exists := users[uuid]
	if exists {
		users[uuid] = user
	}
}

func CheckUserPair(login string, password string) (*user.User, bool) {
	sha256pass := sha256.Sum256([]byte(password))
	hashCheckPassword := base64.StdEncoding.EncodeToString(sha256pass[:])
	for _, u := range users {
		if u.Login == login {
			hashPassword := base64.StdEncoding.EncodeToString(u.PasswordHash[:])
			if hashCheckPassword == hashPassword {
				return &u, true
			}
		}
	}
	return nil, false
}
