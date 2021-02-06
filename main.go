package main

import (
	"Simple-Form-Submission/models"
	"fmt"
	"net/http"

	"Simple-Form-Submission/controllers"

	"github.com/gorilla/mux"
)

//constants to connect with database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "siteusers"
)

func main() {

	//sprintf converts the string to the format given
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserInfoService(psqlInfo)

	//Pass the intialised db instance to new users
	userController := controllers.NewUsers(us)

	router := mux.NewRouter()

	router.HandleFunc("/signup", userController.New).Methods("GET")
	router.HandleFunc("/signup", userController.Create).Methods("POST")

	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", router)

	us.AutoMigrate()
}
