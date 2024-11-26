package tools

import(
	"time"
)

type mockDB struct{}

// Mock data to simulate having Users (LoginDetails)
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

// Mock data to simulate having CoinDetails for each User (LoginDetails)
var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:   100,
		Username: "alex",
	},
	"jason": {
		Coins:   200,
		Username: "jason",
	},
	"marie": {
		Coins:   300,
		Username: "marie",
	},
}

// Method on the mockDB STRUCT to conform to the DatabaseInterface
// Simulates authenticating a user
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	// Initials a LoginDetails struct with default values
	var clientData = LoginDetails{}
	// Checks the "database" (mockLoginDetails) for the username
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

// Method on the mockDB STRUCT to conform to the DatabaseInterface
// Simulates retrieving coin details from the database
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

// Method on the mockDB STRUCT to conform to the DatabaseInterface
// Simulates setting up a database
func (d *mockDB) SetupDatabase() error {
	return nil
}