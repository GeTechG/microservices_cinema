package main

import (
	user_handler "github.com/GeTechG/microservices_cinema/users_service/handlers"
	"github.com/GeTechG/microservices_cinema/users_service/utils/db"
	"log"
	"net/http"
)

func main() {
	db.LoadDb()
	http.HandleFunc("/v1/new_user", user_handler.NewUser)
	http.HandleFunc("/v1/login", user_handler.Login)

	http.HandleFunc("/api/get_user", user_handler.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
