package middleware

import(
	"errors"
	"net/http"
	"github.com/min-verse/goapitutorial/api"
	"github.com/min-verse/goapitutorial/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid username or token provided")

func Authorization(next http.Handler) http.Handler{
	// w: this is the http.ResponseWriter which dictates what this server responds with
	// r: this is the http.Request which is what quantifies and represents the client's request (headers, body, etc.)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// Get the 'username' from the request query
		var username string = r.URL.Query().Get("username")
		// Get the token from the request header
		var token = r.Header.Get("Authorization")

		var err error

		if username == "" || token == ""{
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil{
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// Calls the next middleware in line OR goes straight back
		// to the primary function which handles logic for this route
		next.ServeHTTP(w, r)
	})
}