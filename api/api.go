package api

import(
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct{
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct{
	// Success Code
	Code int
	// Account Balance
	Balance int64
}

// Error Response
type Error struct{
	// Error Code
	Code int

	// Error Message
	Message string
}

// Function to write an error that will be inherited in multiple places
func writeError(w http.ResponseWriter, message string, code int){
	// Creating an Error Struct
	var resp Error = Error{
		Code: code,
		Message: message,
	}

	// Sets the response header "Content-Type" as "application/json"
	w.Header().Set("Content-Type", "application/json")
	// Sets the error code in the header
	w.WriteHeader(code)

	// Put the Error Struct from above as a JSON response
	// Initializes the w http.ResponseWriter as a NewEncoder and then encodes the Error Struct
	json.NewEncoder(w).Encode(resp)
}

var(
	// When the user does something that results in an error
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// When the server internally runs into an error
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Internal Error Occurred", http.StatusInternalServerError)
	}
)