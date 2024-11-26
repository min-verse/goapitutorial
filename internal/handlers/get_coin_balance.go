package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/min-verse/goapitutorial/api"
	"github.com/min-verse/goapitutorial/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	// Initialize a CoinBalanceParams STRUCT defined in api folder
	var params = api.CoinBalanceParams{}
	// Use a schema.Decoder from the 'gorilla/schema' package
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// Grabs the values in the URL and then sets them to the CoinBalanceParams STRUCT
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Get CoinDetails from the "database" to GetUserCoins from the params.Username
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Create a CoinBalanceResponse STRUCT with
	// the balance and status
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	// Sets the response hader and also encodes the CoinBalanceReponse STRUCT
	// as a JSON in the body of the response (if there's an error, it will error out)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}