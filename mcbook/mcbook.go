package mcbook

import (
  "../igo"
  "math/rand"
  "strconv"
)

type position struct {
  occured, wins int
}

type Book map[string]*position

func (self Book) get(i string)(*position) {
  p, exists := self[i]
  if (!exists) {
    p = &position{}
    self[i] = p
  }
  return p
}

func (self Book) Move(game *igo.Game)(bool) {
  x := rand.Intn(10)
  y := rand.Intn(10)
  c := game.Copy()
  c.Move(x,y)
  p := self.get(c.Id())
  v := (p.wins*1000)/(p.occured+1)

  for i:=0; i<10; i++ {
    tx := rand.Intn(10)
    ty := rand.Intn(10)
    c := game.Copy()
    c.Move(x,y)
    p := self.get(c.Id())
    tv := (p.wins*1000 + rand.Intn(10))/(p.occured+1)
    if (tv > v){
      x = tx
      y = ty
      v = tv
    }
  }

  win,_,_,_ := game.Move(x,y)
  return win
}

func (self Book) RecordGame(game *igo.Game) {
  w := game.Winner() - 1
  l := 1 - w
  for _, i := range game.Moves[w] {
    p := self.get(i)
    p.occured++
    p.wins++
  }
  for _, i := range game.Moves[l] {
    p := self.get(i)
    p.occured++
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

func (self Book) Serialize()([]byte) {
  s := ""
  for k,v := range self {
    if (v.occured > 0){
      s += k + " "
      s += strconv.Itoa(v.wins) + " "
      s += strconv.Itoa(v.occured) + "\n"
    }
  }
  return []byte(s)
}