package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalenceparams struct {
	Username string
}

// respose balenc
type CoinBalanceResponse struct {
	Code    int   //status code
	Balence int64 //  acount bal
}

// errro response
type Error struct {
	Code    int    //error code
	Message string // erro measg
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content- type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An Unexpected Error Occcured.", http.StatusInternalServerError)
	}
)
