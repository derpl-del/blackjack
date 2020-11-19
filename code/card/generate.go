package card

import (
	"fmt"
	"math/rand"
	"time"
)

//Cards Struct
type Cards struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

//Deck type list of cards
type Deck []Cards

//GenerateCard for Generating Card for Deck
func GenerateCard() Deck {
	decks := Deck{}
	CardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	CardValue := []string{"Ace", "Two", "Tree", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	value := 0
	for _, suit := range CardSuits {
		for i, values := range CardValue {
			if i+1 > 11 {
				value = 10
			} else {
				value = i + 1
			}
			name := suit + " " + values
			card := Cards{Name: name, Value: value}
			decks = append(decks, card)
		}
	}
	return decks
}

//ShuffleCard for Shuffling Card in Deck
func (d Deck) ShuffleCard() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	r.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	r.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

//PrintDeck for Print list of Card in Deck
func (d Deck) PrintDeck() {
	fmt.Println(d)
	fmt.Println(len(d))
}

//DrawCard for Print list of Card in Deck
func DrawCard(h Deck, d Deck, num int) (Deck, Deck) {
	h = append(h, d[:num]...)
	d = append(d[:0], d[num:]...)
	//d[len(d)-1], d[0] = d[0], d[len(d)-1]
	return h, d
}
