package dining_philosophers

import (
  "log"
  "time"
  "os"
)

// To avoid concurrency problem
// This will create a thread safe output
var fmt = log.New(os.Stdout, "", 0)

// Here are our philosophers
// I remember them from school so ... They may not have lived during the same century :p
var philosophers = []string{"Aristote", "Spinoza", "Kant"}
// , "Marx", "Nietzsche", "Platon", "Thalès", "Pythagore", "Bouddha"

func Philosopher(name string, left_hand, right_hand chan fork) {
  // Philosopher is sitting
  fmt.Println(name, " has now seated to the table !")
  for {
    fmt.Println(name, " is very hungry :(")
    // He tries to pick a fork from his left hand
    <- left_hand
    // He tries to pick a fork from his right hand
    <- right_hand
    // Achievement unlocked ! Philosopher is eating !
    fmt.Println(name, "is eating, miaaaaaam :D")
    // Eating takes a while, but in fact they are not so hungry
    time.Sleep(time.Duration(time.Second * 3))
    fmt.Println(name, "has finished eating ! Time to share forks !")
    // He drops the fork to his left
    left_hand <- 'f'
    // He drops the fork to his right
    right_hand <- 'f'
    fmt.Println(name, " is now thinking, oh my god, his stomach is full :O")
    time.Sleep(time.Duration(time.Second * 2))
  }
}

func LetsDine() {
  fmt.Println("Table is empty, no one is dining, they are busy :/")

  // First, the first left fork is put to left
  left_hand := make(chan fork, 1)
  left_hand <- 'f'
  // This is the first fork, we remember it to reference it as the right fork of the last philosopher
  first_hand := left_hand

  // This channel will let us wait and wait and wait and wait for the dinner to be infinite !
  time := make(chan bool, 0)

  for i := 1; i < len(philosophers); i++ {
    // Put a fork to philosophers's right side
    right_hand := make(chan fork, 1)
    right_hand <- 'f'
    // Philosopher arrives at the table !
    go Philosopher(philosophers[i], left_hand, right_hand)
    // Current philosophers right hand is next philosopher's left hand
    left_hand = right_hand
  }
  // Last philosopher arrives at the table
  go Philosopher(philosophers[0], left_hand, first_hand)

  // Time to dine !!!!!!!!
  <- time
}