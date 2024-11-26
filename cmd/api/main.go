package main

import(
	"fmt"
	"net/http"
	// Chi is a Golang API framework (can use another one if needed)
	"github.com/go-chi/chi"
	// Importing own module from internal/handlers folder
	"github.com/min-verse/goapitutorial/internal/handlers"
	// Importing logrus to use as a logging service
	log "github.com/sirupsen/logrus"
	// Install external pacakges like chi and logrus
	// using the following command in the command line: go mod tidy
	// These will be added in the go.mod file under require(...)
)

func main(){
	log.SetReportCaller(true)
	// Struct to set up API
	var r *chi.Mux = chi.NewRouter()

	// Defined in internals/handlers
	// This will set up our router
	// which adds the endpoint definitions
	handlers.Handler(r)

	fmt.Println("Starting Coin Balance GO API Server...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil{
		log.Error(err)
	}
}