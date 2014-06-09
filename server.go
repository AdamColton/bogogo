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
    t := r.FormValue("turn")
    o := make([]string,100)
    c := 0
    id := ""

    for tx:=0; tx<10; tx++ {
      for ty:=0; ty<10; ty++{
        if (g.Board(tx,ty) != 0){
          o[c] = "0"
        } else {
          gc := g.Copy()
          gc.Move(tx,ty)
          if (t == "1"){
            id = gc.Iid()
          } else {
            id = gc.Id()
          }
          p := b.Get(id)
          o[c] = strconv.Itoa( (1000*p.Wins())/(p.Occured()+1) )
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