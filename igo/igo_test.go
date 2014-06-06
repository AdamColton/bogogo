package igo

import (
  "testing"
  "strconv"
)

func TestInit(t *testing.T){
  game := Game{}
  game.Init()

  if (game.board[0][0] != 0){
    t.Error("Expected board to initilize to 0")
  }

  if (game.turn != 1){
    t.Error("Expected turn to initilize to 1")
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

func TestGetGroupEmptyBoard(t *testing.T) {
  g := Game{}
  g.Init()
  grp := g.GetGroup(0,0,1)
  if (len(grp.Coords) != 100){
    t.Error("expected 100, got " + strconv.Itoa(len(grp.Coords)))
  }
}