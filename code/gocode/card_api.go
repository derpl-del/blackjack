package gocode

import (
	"encoding/json"
	"net/http"

	"github.com/derpl-del/blackjack/code/card"
)

//List type
var deck card.Deck
var hand card.Deck
var deal card.Deck

//Response struct
type Response struct {
	UserValue   int    `json:"user_value"`
	DealerValue int    `json:"dealer_value"`
	Winner      string `json:"winner"`
	UserScore   int    `json:"user_score"`
	DealerScore int    `json:"dealer_score"`
}

//ResponseGenerator struct
type ResponseGenerator struct {
	UserHand   card.Deck `json:"user_hand"`
	DealerHand card.Deck `json:"dealer_hand"`
}

//ResponseResultDraw struct
type ResponseResultDraw struct {
	Status string `json:"status"`
	Card   string `json:"card"`
	Value  int    `json:"value"`
	Score  int    `json:"score"`
}

//CardDeckGenerator Func
func CardDeckGenerator(w http.ResponseWriter, r *http.Request) {
	deckPointer := &deck
	handPointer := &hand
	dealPointer := &deal
	*deckPointer = card.GenerateCard()
	deck.ShuffleCard()
	*handPointer = card.Deck{}
	*dealPointer = card.Deck{}
	*handPointer, *deckPointer = card.DrawCard(hand, deck, 2)
	*dealPointer, *deckPointer = card.DrawCard(deal, deck, 1)
	res := ResponseGenerator{DealerHand: deal, UserHand: hand}
	json.NewEncoder(w).Encode(res)
}

//CardDrawUser Func
func CardDrawUser(w http.ResponseWriter, r *http.Request) {
	deckPointer := &deck
	handPointer := &hand
	*handPointer, *deckPointer = card.DrawCard(hand, deck, 1)
	json.NewEncoder(w).Encode(hand)
}

//CardDrawDealer Func
func CardDrawDealer(w http.ResponseWriter, r *http.Request) {
	deckPointer := &deck
	dealPointer := &deal
	*dealPointer, *deckPointer = card.DrawCard(deal, deck, 1)
	json.NewEncoder(w).Encode(deal)
}

//ViewDeck Func
func ViewDeck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(deck)
}

//ViewHand Func
func ViewHand(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(hand)
}

//ViewDealer Func
func ViewDealer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(deal)
}

//ViewResult Func
func ViewResult(w http.ResponseWriter, r *http.Request) {
	Dpoint := 0
	Upoint := 0
	var status string
	Dace := false
	Uace := false
	for _, d := range deal {
		if d.Value == 1 {
			Dace = true
		}
		Dpoint = Dpoint + d.Value
	}
	if Dace == true && Dpoint+10 <= 21 {
		Dpoint = Dpoint + 10
	}
	for _, u := range hand {
		if u.Value == 1 {
			Uace = true
		}
		Upoint = Upoint + u.Value
	}
	if Uace == true && Upoint+10 <= 21 {
		Upoint = Upoint + 10
	}
	if Dpoint == Upoint {
		status = "Draw"
	} else if Dpoint > Upoint {
		status = "Dealer Win"
	} else {
		status = "User Win"
	}
	res := Response{DealerValue: Dpoint, UserValue: Upoint, Winner: status, DealerScore: Dpoint, UserScore: Upoint}
	json.NewEncoder(w).Encode(res)
}

//UserDrawResult Func
func UserDrawResult(w http.ResponseWriter, r *http.Request) {
	Upoint := 0
	var status string
	Uace := false
	cards := hand[len(hand)-1].Name
	value := hand[len(hand)-1].Value
	for _, u := range hand {
		if u.Value == 1 {
			Uace = true
		}
		Upoint = Upoint + u.Value
	}
	if Uace == true && Upoint+10 <= 21 {
		Upoint = Upoint + 10
	}
	if Upoint <= 21 {
		status = "alive"
	} else {
		status = "burnout"
	}
	res := ResponseResultDraw{Card: cards, Status: status, Value: value, Score: Upoint}
	json.NewEncoder(w).Encode(res)
}

//DealerDrawResult Func
func DealerDrawResult(w http.ResponseWriter, r *http.Request) {
	Dpoint := 0
	var status string
	Dace := false
	cards := deal[len(deal)-1].Name
	value := deal[len(deal)-1].Value
	for _, u := range deal {
		if u.Value == 1 {
			Dace = true
		}
		Dpoint = Dpoint + u.Value
	}
	if Dace == true && Dpoint+10 <= 21 {
		Dpoint = Dpoint + 10
	}
	if Dpoint <= 21 {
		status = "alive"
	} else {
		status = "burnout"
	}
	res := ResponseResultDraw{Card: cards, Status: status, Value: value, Score: Dpoint}
	json.NewEncoder(w).Encode(res)
}
