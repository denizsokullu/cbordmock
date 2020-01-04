package main

// CBORDSuccessResponse is a baseline cbord success response
type CBORDSuccessResponse struct {
	Code    string `json:"responseCode"`
	Message string `json:"responseMessage"`
}

// CBORDFailureResponse is a baseline cbord bad request response
type CBORDFailureResponse struct {
	Code   string      `json:"systemErrorCode"`
	Type   string      `json:"type"`
	Title  string      `json:"title"`
	Detail interface{} `json:"detail"`
}

func successResponse() CBORDSuccessResponse {
	return CBORDSuccessResponse{
		Code:    "200",
		Message: "OK",
	}
}
