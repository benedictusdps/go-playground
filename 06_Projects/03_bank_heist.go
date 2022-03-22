package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
  // Random simulator
  rand.Seed(time.Now().UnixNano())

  // To track whether the heist is successful or not
  var isHeistOn bool = true

  // First conditional: Sneak past guards
  var eludedGuards int = rand.Intn(100)
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  // Second conditional: Open the vault
  var openedVault int = rand.Intn(100)
  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault can't be opened.")
  }

  // Third conditional: Leaving
  var leftSafely int = rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("You got caught by a security officer.")
      case 1:
        isHeistOn = false
        fmt.Println("You tripped and fell. The police caught up to you.")
      case 2:
        isHeistOn = false
        fmt.Println("The bag had a hole in it. You lost all your money.")
      case 3:
        isHeistOn = false
        fmt.Println("You were betrayed. Your team took your bag and left you.")
      default:
        fmt.Println("Start the getaway car!")
    }
  // Wrapping up
  if isHeistOn {
    var amtStolen int = 10000 + rand.Intn(1000000)
    fmt.Println("The amount stolen is", amtStolen)
  }
  }
  // Show current status of isHeistOn
  fmt.Println(isHeistOn)
}
