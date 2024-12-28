package main

import (
	"fmt"
	"net/http"

	hand "github.com/bhushan-aruto/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	// Initialize router
	r := chi.NewRouter()
	hand.Handler(r)

	fmt.Println("Starting Go API service...")

	fmt.Println(`
      OOOOO   AAAAA   PPPPPP   III   EEEEEEE
G        O   O   A   A   P    P    I    E
G  GG    O   O   AAAAA   PPPPPP    I    EEEEE
G    G   O   O   A   A   P         I    E
 GGGGG   OOOOO   A   A   P        III   EEEEEEE `)

	// Start server
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
