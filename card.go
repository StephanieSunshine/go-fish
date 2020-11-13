package main

import (
  "errors"
)

const(
  Spade   = "\u2660"
  Heart   = "\u2665"
  Club    = "\u2663"
  Diamond = "\u2666"

  Ace     = "A"
  Two     = "2"
  Three   = "3"
  Four    = "4"
  Five    = "5"
  Six     = "6"
  Seven   = "7"
  Eight   = "8"
  Nine    = "9"
  Ten     = "10"
  Jack    = "J"
  Queen   = "Q"
  King    = "K"
)

type Card struct {
  Suit string
  Rank string
}

func (c *Card) Suits() []string {
  return []string{Spade, Heart, Club, Diamond}
}

func (c *Card) Ranks() []string {
  return []string{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}

func (c *Card) String() string {
  return c.Rank+c.Suit
}

func NewCard(s string, r string) (*Card, error) {
  c := new(Card)
  var err error = nil 
  switch s {
    case Spade:
      c.Suit = Spade
    case Heart:
      c.Suit = Heart
    case Club:
      c.Suit = Club
    case Diamond:
      c.Suit = Diamond
    default:
      err = errors.New("Invalid Suit")
  }
  switch r {
    case Ace:
      c.Rank = Ace
    case Two:
      c.Rank = Two
    case Three:
      c.Rank = Three
    case Four:
      c.Rank = Four
    case Five:
      c.Rank = Five
    case Six:
      c.Rank = Six
    case Seven:
      c.Rank = Seven
    case Eight:
      c.Rank = Eight
    case Nine:
      c.Rank = Nine
    case Ten:
      c.Rank = Ten
    case Jack:
      c.Rank = Jack
    case Queen:
      c.Rank = Queen
    case King:
      c.Rank = King
    default:
      err = errors.New("Invalid Rank")
  }
  return c, err
}
