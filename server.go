package main

import (
  "fmt"
  "net/http"
  "./igo"
  "strconv"
  "./mcbook"
  "io/ioutil"
  "strings"
)

func move(w http.ResponseWriter, r *http.Request) {
  g := *igo.FromId(r.FormValue("id"), 2)
  x,_ := strconv.Atoi(r.FormValue("x"))
  y,_ := strconv.Atoi(r.FormValue("y"))
  g.Move(x,y)
  fmt.Fprintf(w, g.Id() )
}

func main(){
  b := mcbook.Deserialize(ReadFile())

  http.HandleFunc("/joseki", func(w http.ResponseWriter, r *http.Request) {
    g := *igo.FromId(r.FormValue("id"), 2)
    o := make([]string,100)
    c := 0

    for tx:=0; tx<10; tx++ {
      for ty:=0; ty<10; ty++{
        if (g.Board(tx,ty) != 0){
          o[c] = "0/0"
        } else {
          gc := g.Copy()
          gc.Move(tx,ty)
          p := b.Get(gc.Iid())
          o[c] = strconv.Itoa( p.Wins()) +"/" + strconv.Itoa(p.Occured()) 
        }
        c++
      }
    }

    fmt.Fprintf(w, strings.Join(o, ",") )
  })

  http.HandleFunc("/move", move)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "server/" + r.URL.Path[1:])
  })
  http.ListenAndServe(":8080", nil)
}

func ReadFile()(string){
  b, err := ioutil.ReadFile("test.txt")
  if err != nil { return "" }
  return string(b)
}