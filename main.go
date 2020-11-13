package main

import(
  "fmt"
  "strings"
  "os"
  "math/rand"
  "time"
)

func main() {

  rand.Seed(time.Now().UnixNano())

  // deck.go
  mydeck := Deck{}

  // hand.go
  myhand := Hand{}
  cpuhand := Hand{}
  choice := ""

  fmt.Println("Go Fish! v1.0")
  
  mydeck.Init()
  mydeck.Shuffle(3)

  // deal
  count := 7
  for count > 0 {
    cards, err := mydeck.Draw()
    if err != nil {
      panic(err)
    }
    myhand.Add(cards...)

    cards, err = mydeck.Draw()
    if err != nil {
      panic(err)
    }
    cpuhand.Add(cards...)
    count -= 1
  }

  for (len(myhand.Order)>0)&&(len(cpuhand.Order)>0)&&(len(mydeck.Order)>0) {
  myhand.Sort(mydeck.CardRank)

  fmt.Println("\nBooks Human:", myhand.BooksCount, "Computer:", cpuhand.BooksCount)
  for _, c := range myhand.Order {
    fmt.Println(c.String())
  }

  valid_choice := false
  for !valid_choice {
    fmt.Print("\n[M]ake Book, [G]uess, [Q]uit => ")
    fmt.Scanln(&choice)
    switch strings.ToUpper(choice) {
      case "M":
        valid_choice = true
      case "G":
        valid_choice = true
      case "Q":
        os.Exit(0)
      default:
        fmt.Println("Invalid entry. Try again")
    }
  }

  valid_choice = false
  cchoice := ""
  for !valid_choice {
    fmt.Println("\nPossible Ranks: [", strings.Join(myhand.Ranks(),", "),"]")
    fmt.Print("Choose which rank to work with => ")
    fmt.Scanln(&cchoice)
    if contains(myhand.Ranks(), strings.ToUpper(cchoice)) {
      // fmt.Println("Valid choice")
      valid_choice = true
    }else{
      fmt.Println("Invalid choice")
    }
  }

  switch strings.ToUpper(choice) {
    case "M":
      fmt.Println("Trying to make a book of:", strings.ToUpper(cchoice))
      res := myhand.MakeBook(cchoice)
      if res {
        fmt.Println("Worked")
      } else {
        fmt.Println("Didn't Work")
      }
      continue
    case "G":
      fmt.Println("Asking computer for:", strings.ToUpper(cchoice))
      cards, ans := cpuhand.HaveAny(strings.ToUpper(cchoice))
      if ans {
        fmt.Println("They Did!")
        for _, c := range cards {
          fmt.Println(c.String())
        }
        myhand.Add(cards...)
      } else {
        fmt.Println("Go Fish!")
        fish, _ := mydeck.DrawRandom()
        fmt.Println("Drew:", fish.String())
        // if fish equal what we were looking for, we get another turn here
        myhand.Add(fish)
     }
  }

  fmt.Println()

  // Computer

  for _, c := range cpuhand.Ranks() {
    _ = cpuhand.MakeBook(c)
  }
  r := cpuhand.Ranks()
  // if they booked out, don't try to play with an empty hand or rand.Intr complains (max range must be greater than 0)
  if len(r) == 0 {
    continue
  }
  g := r[rand.Intn(len(r))]
  fmt.Println("Computer looking for any:", g)
  cards, ans := myhand.HaveAny(g)
  if ans {
    fmt.Println("You Did!")
    for _, c := range cards {
      fmt.Println(c.String())
    }
    cpuhand.Add(cards...)
  } else {
    fmt.Println("Go Fish!")
    fish, _ := mydeck.DrawRandom()
    //fmt.Println("Drew:", fish.String())
    // if fish equal what we were looking for, we get another turn here
    cpuhand.Add(fish)
  }

//  fmt.Println("Computer hand:")
//  for _, c := range cpuhand.Order {
//    fmt.Println(c.String())
//  }
}

fmt.Println("Game Over")
fmt.Println("Score: Human:", myhand.BooksCount, "Computer:", cpuhand.BooksCount)
if myhand.BooksCount > cpuhand.BooksCount {
  fmt.Println("Human Wins!")
} else if myhand.BooksCount < cpuhand.BooksCount {
  fmt.Println("Computer Wins!")
} else {
  fmt.Println("Its a Tie!")
}
}
