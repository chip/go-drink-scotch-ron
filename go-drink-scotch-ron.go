package main

import (
    "fmt"
    "math/rand"
    "time"
    "io/ioutil"
    "strings"
)

func main() {
  filename := "ron-burgundy-quotes.txt"
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("Uh-oh... ", err)
  } else {
    quotes := strings.Split(string(content), "\n")
    length := len(quotes)
    rand.Seed( time.Now().UTC().UnixNano())
    index := rand.Intn(length)

    fmt.Println(quotes[index])
  }
}
