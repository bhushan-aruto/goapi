package internal

import (
	"encoding/json"
	"net/http"

	"github.com/bhushan-aruto/goapi/api"
	"github.com/bhushan-aruto/goapi/internal/tools"
	"github.com/gorilla/schema"

	// "github.com/opencontainers/image-spec/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalence(w http.ResponseWriter, r *http.Request) {
	var parms = api.CoinBalenceparams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&parms, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	if tokenDetails = (*database).GetUserCoins(parms.Username); tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var response = api.CoinBalanceResponse{
		Balence: tokenDetails.Coins,
		Code:    http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
		api.InternalErrorHandler(w)
		return
	}
}
