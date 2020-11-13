package main

import(
  "fmt"
  "math/rand"
  "time"
  "errors"
)

type Deck struct {
  Order []*Card
  Acelow bool
  CardRank []string
}

// deck constructor
func (d *Deck) Init() {
  d.Order = nil
  d.CardRank = []string{ Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace }
  c := Card{}
  for _, suit := range c.Suits() {
    for _, rank := range c.Ranks() {
      cc, err := NewCard(suit, rank)
      if err != nil {
        panic(err)
      }
      d.Order = append(d.Order, cc)
    }
  }
}

// shuffle the deck
func (d *Deck) Shuffle(count ...uint32) {
  var counter uint32
  if len(count) == 0 {
    counter = 1
  } else if count[0] == 0 {
    counter = 1
  } else {
    counter = count[0]
  }

  for counter > 0 {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(d.Order), func(i, j int) { d.Order[i], d.Order[j] = d.Order[j], d.Order[i] })
    counter -= 1
  }
}

// draw one from somewhere random in the deck
func (d *Deck) DrawRandom() (*Card, error) {
  rand.Seed(time.Now().UnixNano())
  if len(d.Order) == 0 {
    return &Card{}, errors.New("No cards left in deck")
  }
  index := rand.Intn(len(d.Order))
  card := d.Order[index]
  switch index {
    case 0:
      d.Order = d.Order[1:]
    case 1:
      d.Order = append([]*Card{d.Order[0]}, d.Order[2:len(d.Order)]...)
    case len(d.Order):
      d.Order = d.Order[:len(d.Order)-1]
    default:
      d.Order = append(d.Order[:index], d.Order[index+1:]...)
  }
  return card, nil
}

// draw count from top of deck
func (d *Deck) Draw(count ...uint32) ([]*Card, error) {
  var counter uint32
  var hand []*Card
  if len(count) == 0 {
    counter = 1
  } else if count[0] == 0 {
    counter = 1
  } else {
    counter = count[0]
  }
  if uint32(len(d.Order)) < counter {
    return hand, errors.New(fmt.Sprintf("Not enough cards available want:", counter, "have:", len(d.Order)))
  }

  for counter > 0 {
    hand = append(hand, d.Order[0])
    d.Order = d.Order[1:len(d.Order)]
    counter -= 1
  }

  return hand, nil

}
