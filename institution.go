package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// Institution JSON representation
type Institution struct {
	UUID         string `json:"uuid"`         // Instituion ID
	Name         string `json:"name"`         // Institution full name. Suitable for display to users.
	ShortName    string `json:"shortName"`    // Institution short name. Used in constructed URLs for browser-based authentication. Not intended for display to users.
	AuthURL      string `json:"authUrl"`      // Only returned for schools with browser-based authentication (CAS/Shib). Otherwise is empty string.
	AuthType     int16  `json:"authType"`     // Indicator of whether to ask for credentials (hosted/LDAP) or to use browser-based authentication (CAS/Shib). 1=direct login, 2=web-based
	ProgramName  string `json:"programName"`  // Name of the institution card program. Suitable for display to users.
	PrimaryColor string `json:"primaryColor"` // RGB string (ex: "F84525") representing the primary color for the institution. This corresponds to the GET CBORD Student background color configuration.
}

// CBORDInstitutionSuccessResponse is a success response for /institution
type CBORDInstitutionSuccessResponse struct {
	CBORDSuccessResponse
	Type            string      `json:"type"`
	InstitutionList interface{} `json:"institutionList"`
}

func institutionHandler(w http.ResponseWriter, r *http.Request) {
	institutionUUID, err := uuid.NewUUID()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := CBORDInstitutionSuccessResponse{
		CBORDSuccessResponse: successResponse(),
		Type:                 "institution",
		InstitutionList: []Institution{
			Institution{
				UUID:         institutionUUID.String(),
				Name:         "Vanderbilt University",
				ShortName:    "vu",
				AuthURL:      "https://test-url.com",
				AuthType:     0,
				ProgramName:  "Vu Bucks",
				PrimaryColor: "FFFFFF",
			},
		},
	}

	message, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}
