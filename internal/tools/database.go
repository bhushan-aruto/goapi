package tools

import "log"

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}
type DatabseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabseInterface, error) {
	var database DatabseInterface = &mockDB{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &database, nil
}
