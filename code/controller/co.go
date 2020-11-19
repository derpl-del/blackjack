package controller

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/blackjack/code/gocode"
	"github.com/gorilla/mux"
)

var r *mux.Router

//CallController for generate api
func CallController() {
	fmt.Println("morning")
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/GenerateDeck", gocode.CardDeckGenerator).Methods("GET")
	r.HandleFunc("/api/v1/DrawCardUser", gocode.CardDrawUser).Methods("GET")
	r.HandleFunc("/api/v1/DrawCardDealer", gocode.CardDrawDealer).Methods("GET")
	r.HandleFunc("/api/v1/ViewDeck", gocode.ViewDeck).Methods("GET")
	r.HandleFunc("/api/v1/ViewHand", gocode.ViewHand).Methods("GET")
	r.HandleFunc("/api/v1/ViewDealer", gocode.ViewDealer).Methods("GET")
	r.HandleFunc("/api/v1/ViewResult", gocode.ViewResult).Methods("GET")
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
