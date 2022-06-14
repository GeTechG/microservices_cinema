package user_handler

import (
	"encoding/json"
	"errors"
	user "github.com/GeTechG/microservices_cinema/users_service/entities"
	"github.com/GeTechG/microservices_cinema/users_service/utils"
	"github.com/GeTechG/microservices_cinema/users_service/utils/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

type RequestUser struct {
	Login        string `json:"login" validate:"required"`
	Password     string `json:"password" validate:"required"`
	DateBirthday string `json:"date_birthday" validate:"required"`
}

const layoutTime string = "2006-01-02"

func NewUser(w http.ResponseWriter, r *http.Request) {
	var err error
	utils.Request(r, &w, http.MethodPost, true, func() {
		var requestUser RequestUser

		err = utils.ParseJson(r, &requestUser)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}

		validate := validator.New()

		err = validate.Struct(requestUser)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}

		date, err := time.Parse(layoutTime, requestUser.DateBirthday)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}
		newUser := user.NewUser(requestUser.Login, requestUser.Password, date)
		err = db.AddNewUser(newUser)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}

		utils.SetResponse(&w, http.StatusOK, []byte(newUser.Id.String()))
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	utils.Request(r, &w, http.MethodPost, true, func() {
		var requestUser RequestUser
		err := utils.ParseJson(r, &requestUser)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}

		logginedUser, exists := db.CheckUserPair(requestUser.Login, requestUser.Password)
		if !exists {
			utils.SetResponseException(&w, errors.New("not found this pair login and password"))
			return
		}
		utils.SetResponse(&w, http.StatusOK, []byte(logginedUser.Id.String()))
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.Request(r, &w, http.MethodPost, false, func() {
		bytesBody, err := io.ReadAll(r.Body)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}
		body := string(bytesBody)

		userUuid, err := uuid.Parse(body)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}
		userByUuid := db.GetUserByUuid(userUuid)

		response, err := json.Marshal(userByUuid)
		if err != nil {
			utils.SetResponseException(&w, err)
			return
		}
		utils.SetResponse(&w, http.StatusOK, response)
	})
}
