package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetail = map[string]LoginDetails{
	"naruto": {
		AuthToken: "23456789",
		Username:  "naruto",
	},
	"rock_le": {
		AuthToken: "23456",
		Username:  "rock_lee",
	},
	"light": {
		AuthToken: "23459",
		Username:  "light",
	},
}

// Mocked data for user coin details
var mockCoinDetails = map[string]CoinDetails{
	"naruto": {
		Coins:    100,
		Username: "naruto",
	},
	"rocklee": {
		Coins:    400,
		Username: "rocklee",
	},
	"light": {
		Coins:    600,
		Username: "light",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {

	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetail[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {

	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

// SetupDatabase method implements the DatabseInterface
func (d *mockDB) SetupDatabase() error {
	// Simulate database setup
	return nil
}
