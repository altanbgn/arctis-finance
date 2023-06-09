package server

import (
	"log"
	"net/http"

	"github.com/altanbgn/arctis-finance/controllers"
)

func Start() error {
	http.HandleFunc("/users", controllers.GetAllUsers)
	http.HandleFunc("/users/{id}", controllers.GetUser)
	http.HandleFunc("/users", controllers.CreateUser)
	http.HandleFunc("/users/{id}", controllers.UpdateUser)
	http.HandleFunc("/users/{id}", controllers.DeleteUser)

	log.Printf("Listening on :8080")
	return http.ListenAndServe(":8080", nil)
}
