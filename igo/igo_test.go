package igo

import (
  "testing"
  "strconv"
)

func TestInit(t *testing.T){
  game := Game{}
  game.Init(2)

  if (game.board[0][0] != 0){
    t.Error("Expected board to initilize to 0")
  }

  if (game.turn != 1){
    t.Error("Expected turn to initilize to 1")
  }
}

func TestMove(t *testing.T) {
  game := Game{}
  game.Init(2)
  game.Move(0,0)
  game.Move(0,1)
  game.Move(0,2)

  if (game.Board(0,0) != 1){
    t.Error( "expected 0,0 to be 1, got "+strconv.Itoa(game.Board(0,0)) )
  }
  if (game.Board(0,1) != 2){
    t.Error( "expected 0,1 to be 2, got "+strconv.Itoa(game.Board(0,1)) )
  }
  if (game.Board(0,2) != 1){
    t.Error( "expected 0,2 to be 1, got "+strconv.Itoa(game.Board(0,2)) )
  }
}

func TestCoordQueue(t *testing.T) {
  q := make(map[[2]int]bool)
  i := [2]int{1,2}
  q[i] = true
  if (!q[[2]int{1,2}]){
    t.Error("expected true")
  }
  if (q[[2]int{2,3}]){
    t.Error("expected false")
  }
}

func TestOneEye(t *testing.T) {
  game := Game{}
  game.Init(2)

  game.Move(0,1)
  game.Move(0,2)
  game.Move(1,0)
  game.Move(1,2)
  win, gs, chn := game.Move(1,1)

  if (len(chn) != 3){
    t.Error("Expected len 3, got ", len(chn))
  }

  if (win) {
    t.Error("did not expect a win")
  }

  if (len(gs) != 2){
    t.Error("Expected 2 groups, got ", len(gs))
  }

  c := 0
  for _, g := range gs {
    if (g.Eye()){
      c++
    }
  }
  if (c != 1){
    t.Error("Expected one eye, got ", c)
  }
}

func TestWin(t *testing.T) {
  game := Game{}
  game.Init(2)

  game.Move(1,0)
  game.Move(2,0)
  game.Move(0,1)
  game.Move(2,1)
  game.Move(1,1)
  game.Move(2,2)
  game.Move(1,2)
  game.Move(2,3)
  game.Move(1,3)
  game.Move(2,4)

  win, _, _ := game.Move(0,3)

  if (!win) {
    t.Error("Expected a win")
  }
}