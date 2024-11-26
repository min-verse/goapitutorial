package tools

import(
	log "github.com/sirupsen/logrus"
)

// Database collections

// These are database models essentially
type LoginDetails struct{
	AuthToken string
	Username string
}

type CoinDetails struct{
	Coins int64
	Username string
}

// Using interface here so we can swap out databases easily
type DatabaseInterface interface{
	// We will define these methods
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// Function to set up a database and it returns
// both a DatabaseInterface and an error value
func NewDatabase() (*DatabaseInterface, error){
	// mockDB STRUCT that implements interface
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil{
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
