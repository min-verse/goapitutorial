package handlers

import(
	// Imports the framework
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	// Imports middleware package
	"github.com/min-verse/goapitutorial/internal/middleware"
)

// This function is capitalized so that it can be EXPORTED and then IMPORTED elsewhere
func Handler(r *chi.Mux){
	// Middleware is function that gets called before primary function
	// (primary function handles endpoint logic)

	// Global Middleware
	r.Use(chimiddle.StripSlashes)

	// Can define a route simply as:
	// r.Get('/route', callback)
	// r.Post('/route', callback)

	// Defines a route and a callback function
	// which will handle that route's logic
	// (This is basically like Express.js)
	r.Route("/account", func(router chi.Router){
		// Middleware for "/account" route specifically
		// (OPPOSITE of Global Middleware as this is route-specific)

		// This Authorization function will be created in
		// the middleware package later (middleware imported above)
		router.Use(middleware.Authorization)

		// Defines the route at "/account/coins"
		router.Get("/coins", GetCoinBalance)
	})
}