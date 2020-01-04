package main

import (
	"encoding/json"
	"net/http"
)

// Tender is the tender data
type Tender struct {
	AccountDisplayName string `json:"accountDisplayName"` // Account display name
	PaymentSystemType  int16  `json:"paymentSystemType"`  // Payment system type. 1=OPCS, 2=CS Gold, 3=Credit Card
	AccountTender      string `json:"accountTender"`      // Account tender
	AccountType        int16  `json:"accountType"`        // Account type. 1=meals, 2=inclining balance/charge, 3=declining balance
	Balance            string `json:"balance"`            // Balance
}

// CBORDPatronPaymentTokenSuccessResponse is a success response for /patron/payment_token (Testing only!)
type CBORDPatronPaymentTokenSuccessResponse struct {
	CBORDSuccessResponse
	PaymentToken string `json:"paymentToken"`
}

// CBORDPatronBalanceSuccessResponse is a success response for /patron/{token}/token/balance
type CBORDPatronBalanceSuccessResponse struct {
	CBORDSuccessResponse
	Balance     string `json:"balance"`
	CashlessUID string `json:"cashlessUid"`
}

// CBORDPatronTenderSuccessResponse is a success response for /patron/{token}/token/tender
type CBORDPatronTenderSuccessResponse struct {
	CBORDSuccessResponse
	TenderList   []Tender `json:"tenderList"`
	UniversityID string   `json:"universityId"`
	CashlessUID  string   `json:"cashlessUid"`
}

func patronPaymentTokenHandler(w http.ResponseWriter, r *http.Request) {
	response := CBORDPatronPaymentTokenSuccessResponse{
		CBORDSuccessResponse: successResponse(),
		PaymentToken:         "123456789",
	}

	message, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func patronBalanceHandler(w http.ResponseWriter, r *http.Request) {
	response := CBORDPatronBalanceSuccessResponse{
		CBORDSuccessResponse: successResponse(),
		Balance:              "50.00",
		CashlessUID:          "123456789",
	}

	message, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func patronTenderHandler(w http.ResponseWriter, r *http.Request) {
	response := CBORDPatronTenderSuccessResponse{
		CBORDSuccessResponse: successResponse(),
		TenderList: []Tender{
			Tender{
				AccountDisplayName: "VU Bucks",
				PaymentSystemType:  0,
				AccountTender:      "tender",
				AccountType:        1,
				Balance:            "50.00",
			},
		},
		UniversityID: "123123123",
		CashlessUID:  "123456789",
	}

	message, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}
