package main

import (
	"fmt"
	"net/http"

	"assesment.com/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/customerdetails", handlers.PostCustomerDetailHandler).Methods("POST")
	router.HandleFunc("/customerdetails/{customid}", handlers.GetCustomerHandler).Methods("GET")
	router.HandleFunc("/customerdetails/{customid}", handlers.DeleteCustomerHandler).Methods("DELETE")
	router.HandleFunc("/customerdetails/{customid}", handlers.UpdateCustomerHandler).Methods("PUT")
	router.HandleFunc("/customerdetails", handlers.ListCustomerHandler).Methods("GET")

	http.ListenAndServe(":9898", router)
	fmt.Println("Server started...")
}
