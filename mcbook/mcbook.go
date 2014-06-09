package mcbook

import (
  "../igo"
  "math/rand"
  "strconv"
  "strings"
)

type position struct {
  occured, wins int
}
func (self *position) Occured()(int) { return self.occured }
func (self *position) Wins()(int) { return self.wins }

type Book map[string]*position

func (self Book) Get(i string)(*position) {
  p, exists := self[i]
  if (!exists) {
    p = &position{}
  }
  return p
}

func (self Book) Move(game *igo.Game)(bool) {
  x := 0
  y := 0
  v := -1

  for tx:=0; tx<10; tx++ {
    for ty:=0; ty<10; ty++{
      if (game.Board(tx,ty) != 0){
        continue
      }
      c := game.Copy()
      c.Move(tx,ty)
      p := self.Get(c.Id())
      tv := (p.wins*1000 + rand.Intn(10))/(p.occured+1)
      if (tv > v){
        x = tx
        y = ty
        v = tv
      }
    }
  }

  win,_,_,_ := game.Move(x,y)
  return win
}

func (self Book) RecordGame(game *igo.Game) {
  w := game.Winner() - 1
  l := 1 - w
  for _, i := range game.Moves[w] {
    p := self.Get(i)
    p.occured++
    p.wins++
    self[i] = p
  }
  for _, i := range game.Moves[l] {
    p := self.Get(i)
    p.occured++
    self[i] = p
  }
}

func (self Book) PlayAGame(sides int)(*igo.Game) {
  game := igo.Game{}
  game.Init(sides)

  win := false

  for !win {
    win = self.Move(&game)
  }

  return &game
}

func (self Book) Serialize()(string) {
  s := make([]string,0,len(self))
  for k,v := range self {
    if (v.occured > 0){
      s = append(s, k + strconv.Itoa(v.wins) + "/" + strconv.Itoa(v.occured))
    }
  }
  return strings.Join(s, "\n")
}

func Deserialize(s string)(Book) {
  b := Book{}
  l := strings.Split(s, "\n")
  cur := 0
  for cur < (len(l)-1) {
    k := ""
    for i:=0; i<10; i++ {
      k += l[cur] + "\n"
      cur++
    }
    v := strings.Split(l[cur], "/")
    p := b.Get(k)
    p.wins,_ = strconv.Atoi( v[0] )
    p.occured,_ = strconv.Atoi( v[1] )
    b[k] = p
    cur++
  }
  return b 
}