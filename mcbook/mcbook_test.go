package mcbook

import (
  "testing"
  "../igo"
)

func TestBookRecord(t *testing.T) {
  b := Book{}
  b["0"] = b.get("0")
  b.get("0").occured = 10
  if (b["0"].occured != 10) {
    t.Error("expected 10")
  }
}

func TestMove(t *testing.T) {
  b := Book{}
  g := igo.Game{}
  g.Init(2)

  if (g.Turn() != 1){
    t.Error("it should be player 1's turn")
  }

  b.Move(&g)

  if (g.Turn() != 2){
    t.Error("it should be player 2's turn")
  }
}

func TestWin(t *testing.T) {
  game := igo.Game{}
  game.Init(2)

  _,_,_,mv1 := game.Move(1,0)
  _,_,_,mv2 := game.Move(2,0)
  game.Move(0,1)
  game.Move(2,1)
  game.Move(1,1)
  game.Move(2,2)
  game.Move(1,2)
  game.Move(2,3)
  game.Move(1,3)
  game.Move(2,4)
  game.Move(0,3)

  if (game.Winner() != 1) {
    t.Error("Expected a win for player 1")
  }

  b := Book{}
  b.RecordGame(&game)

  if (b.get(mv1).wins != 1 && b.get(mv1).occured != 1){
    t.Error("Expected move one to register one win")
  }

  if (b.get(mv2).wins != 0 && b.get(mv2).occured != 1){
    t.Error("Expected move two to register one occurance and no wins")
  }
}

func TestPlayAGame(t *testing.T) {
  b := Book{}
  b.RecordGame( b.PlayAGame(2) )

  if (len(b) < 11){
    t.Error("did not record game")
  }
}

func TestDeserialize(t *testing.T) {
  b := Book{}
  b.RecordGame( b.PlayAGame(2) )
  
  b2 := Deserialize(b.Serialize())

  for k,v := range b {
    v2 := b2.get(k)
    if (v.occured != v2.occured || v.wins != v2.wins){
      t.Error("Serialize/Deserialize cycle failed")
      return
    }
  }
}