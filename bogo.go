/*
todo
- load book
- threading
- server
*/

package main

import (
  //"./igo"
  "./mcbook"
  "io/ioutil"
  "fmt"
)

func main() {
  b := mcbook.Book{}
  wins := [2]int{}
  for i:=0; i<10000; i++ {
    g := b.PlayAGame(2)
    b.RecordGame(g)
    wins[g.Winner()-1]++
  }
  fmt.Println(wins)
  ioutil.WriteFile("test.txt", b.Serialize(), 0644)
}