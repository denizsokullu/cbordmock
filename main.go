package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// institution
	router.HandleFunc("/institution", institutionHandler).Methods("GET")

	// patron
	router.HandleFunc("/patron/payment_token", patronPaymentTokenHandler).Methods("GET") // testing only

	router.HandleFunc("/patron/{paymentToken}/token/balance", patronBalanceHandler).Methods("GET")
	router.HandleFunc("/patron/{paymentToken}/token/tender", patronTenderHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
