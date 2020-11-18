package main

import (
	"fmt"

	"github.com/derpl-del/blackjack/code/card"
)

func main() {
	fmt.Println("====================================++++++++++++++++++++++++++++++++++++++==========================")
	deck := card.GenerateCard()
	deck.ShuffleCard()
	deck.PrintDeck()
	fmt.Println(len(deck))
	draw, deck := card.DrawCard(deck)
	draw.PrintDeck()
	deck.PrintDeck()
}
