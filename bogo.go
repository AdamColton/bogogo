/*
todo
- load book
- threading
- server
*/

package main

import (
  "./igo"
  "./mcbook"
  "os"
  "bufio"
  "time"
  "runtime"
  "fmt"
  "io/ioutil"
)

func main() {
  th := runtime.NumCPU()
  runtime.GOMAXPROCS(th)
  fmt.Print("Reading File...")
  b := mcbook.Deserialize(ReadFile())
  fmt.Println("Done")
  w := Writer(b)
  stop := make([]chan bool, th)

  for i:=0; i<th; i++ {
    s := make(chan bool, 1)
    go Player(i, b, w, s)
    stop[i] = s
  }

  for {
    time.Sleep(time.Hour)
  }
}

func Player(id int, b mcbook.Book, ch chan<- *igo.Game, stop <-chan bool) {
  for{
    select{
      case <-stop:
        fmt.Println("Exit: ", id)
        return
      default:
        ch <- b.PlayAGame(2)
    }
  }
}

func Writer(book mcbook.Book)(chan<- *igo.Game) {
  ch := make(chan *igo.Game, 100)
  go func(ch <-chan *igo.Game) {
    c := 0
    for {
      g := <-ch
      book.RecordGame(g)
      fmt.Print(".")
      c++
      if (c==1000){
        fmt.Print("\nSerializing...")
        s := book.Serialize()
        fmt.Println("Done")
        go WriteFile(s)
        c = 0
      }
    }
  }(ch)
  return ch
}

func WriteFile(s string) {
  fo, err := os.Create("test.txt")
  defer func() {
      if err := fo.Close(); err != nil {
          panic(err)
      }
  }()
  if err != nil { panic(err) }
  w := bufio.NewWriter(fo)
  fmt.Fprint(w, s)
  w.Flush()
}

func ReadFile()(string){
  b, err := ioutil.ReadFile("test.txt")
  if err != nil { return "" }
  return string(b)
}

func cheat(){
  fmt.Println(time.Minute)
}