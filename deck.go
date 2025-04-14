package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck () deck {
	cards := deck{}

	setsOfCardType := deck{
		"Spades",
		"Clubs",
		"Hearts",
		"Diamonds",
	}
	setsOfCardValue := deck{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, cardType := range setsOfCardType{
		for _,value:= range setsOfCardValue{
			cards = append(cards, value+" of "+cardType)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}
func (d deck) print() {
	for i, card := range d {

		fmt.Println(i+1, card)
	}
}

func (d deck) toString () string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	
	return err
}


func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err!=nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

s := strings.Split(string(bs), ",")

return deck(s)
}

func (d deck) suffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		newPosition := r.Intn(len(d)-1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}