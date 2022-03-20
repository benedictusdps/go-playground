package main

import "fmt"

func main() {
  // Create the variable jellybeanCounter
  // here and make its type int8
  var jellybeanCounter int8

  // Go will raise errors both for
  // jellybeanCounter being unused
  // and for "fmt" being imported
  // and unused.
  
  // Uncomment the following line
  // and watch the program run
  // successfully!
  
  fmt.Println(jellybeanCounter)
}