package main

/*import(
  "fmt"
)*/

type Hand struct {
  Order []*Card
  Books []*Card
  BooksCount int
}

// Just pass Deck.CardRank
func (h *Hand) Sort(rank []string) {
  var q []*Card
  for _, r := range rank {
    for _, c := range h.Order {
      if c.Rank == r {
        q = append(q, c)
      }
    }
  }
  h.Order = q
}

func (h *Hand) Add(c ...*Card) {
  h.Order = append(h.Order, c...)
}

func (h *Hand) Ranks() []string {
  var r []string
  for _, c := range h.Order {
    if !contains(r, c.Rank) {
      r = append(r, c.Rank)
    }
  }
  return r
}

func (h *Hand) MakeBook(b string) bool {
  if !contains(h.Ranks(), b) {
    return false
  }
  q := []*Card{}
  s := []*Card{}
  for _, c := range h.Order {
    if c.Rank == b {
      q = append(q, c)
    }else{
      s = append(s, c)
    }
  }
  if len(q) == 4 {
    h.Books = append(h.Books, q...)
    h.BooksCount += 1
    h.Order = s
    return true
  }
  return false
}

func (h *Hand) HaveAny(r string) ([]*Card, bool) {
  if !contains(h.Ranks(), r) {
    return []*Card{}, false
  }
  q := []*Card{}
  s := []*Card{}
  for _, c := range h.Order {
    if c.Rank == r {
      q = append(q, c)
    }else{
      s = append(s, c)
    }
  }
  h.Order = s
  return q, true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
