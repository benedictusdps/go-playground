package main

import (
	"fmt"
	"time"
)

func main() {
  // Print title
  fmt.Println("Here comes the Gopher!")
  // Prints the Go gopher ASCII Art
  fmt.Println("      `.-::::::-.`")
  fmt.Println("  .:-::::::::::::::-:.")
  fmt.Println("  `_:::    ::    :::_`")
  fmt.Println("   .:( ^   :: ^   ):.")
  fmt.Println("   `:::   (..)   :::.")
  fmt.Println("   `:::::::UU:::::::`")
  fmt.Println("   .::::::::::::::::.")
  fmt.Println("   O::::::::::::::::O")
  fmt.Println("   -::::::::::::::::-")
  fmt.Println("   `::::::::::::::::`")
  fmt.Println("    .::::::::::::::.")
  fmt.Println("      oO:::::::Oo")
  // Prints the current time in UTC
  fmt.Println(time.Now())
}