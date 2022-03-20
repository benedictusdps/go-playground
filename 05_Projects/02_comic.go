package main

import "fmt"

func main() {
  // First comic: Mr. GoToSleep
  var publisher string = "DizzyBooks Publishing Inc."
  var writer string = "Tracey Hatchet"
  var artist string = "Jewel Tampson"
  var title string = "Mr. GoToSleep"
  var year int = 1997
  var pageNumber int = 14
  var grade float32 = 6.5

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "at", year, "with", pageNumber, "pages graded at", grade,)
  fmt.Println("")

  // Second comic: Epic
  publisher = "DizzyBooks Publishing Inc."
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  title = "Epic Vol. 1"
  year = 2013
  pageNumber = 160
  grade = 9.0

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "at", year, "with", pageNumber, "pages graded at", grade,)
}