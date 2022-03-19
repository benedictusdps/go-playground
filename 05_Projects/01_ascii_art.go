package main

import (
  "fmt"
  "time"
)

func main() {
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