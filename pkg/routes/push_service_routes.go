package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sk25469/push_noti_service/pkg/controller"
)

func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.SignUpHandler).Methods("POST")
	r.HandleFunc("/login", controller.SignInHandler).Methods("POST")

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server is listening on :8080🚀")
	http.ListenAndServe(":8080", nil)
}
