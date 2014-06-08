package igo

type Game struct {
  turn, sides, winner int
  board [10][10]int
  Moves [2][]string
}

type Coord [2]int

type group struct {
  Coords []Coord
  eye bool
}

type chain struct{
  pieces map[Coord]bool
  liberty bool
}

func (self *group) Eye()(bool) { return self.eye }

func (self *Game) Board(x,y int)(int) { return self.board[x][y] }
func (self *Game) Turn()(int) { return self.turn}
func (self *Game) Winner()(int) { return self.winner }

func (self *Game) Copy()(Game) {
  g := Game{}
  g.turn = self.turn
  g.sides = self.sides
  for i:=0; i<2; i++ {
    g.Moves[i] = make([]string, 0, len(g.Moves[i]))
    for _, p := range self.Moves[i] {
      g.Moves[i] = append(g.Moves[i], p)
    }
  }
  for x := 0 ; x<10 ; x++ {
    for y := 0 ; y<10 ; y++ {
      g.board[x][y] = self.board[x][y]
    }
  }
  return g
}

func (self *Game) Init(sides int)(){
  self.turn = 1
  self.sides = sides
  if (self.sides > 2){
    self.sides = 2
  }
  self.Moves[0] = make([]string,0)
  self.Moves[1] = make([]string,0)
}

func (self *Game) Move(x, y int)(bool, []group, chain, string){
  win := false
  gs := make([]group,0,0)
  chn := chain{}
  id := ""
  if (self.board[x][y] == 0){
    self.board[x][y] = self.turn
    self.removeDeadPieces()
    id = self.Id()
    self.Moves[self.turn-1] = append(self.Moves[self.turn-1], id)
    win, gs, chn = self.CheckForWin(x,y)
    if (win){
      self.winner = self.turn
    }
    self.turn = 3 - self.turn
  }
  return win, gs, chn, id
}

func (self *Game) GetChain(x,y int)(chain){
  p := self.board[x][y]
  proc := make(map[Coord]bool)
  q := make([]Coord, 0, 100)
  c := Coord{x,y}
  chn := chain{}
  chn.pieces = make(map[Coord]bool)
  dirs := [...]Coord{ {-1,0}, {1,0}, {0, -1}, {0, 1}}

  q = append(q,c)
  proc[c] = true

  for len(q) > 0 {
    c, q = q[len(q)-1], q[:len(q)-1]
    if (c[0] >= 0 && c[0] < 10 && c[1] >= 0 && c[1] < 10){
      if (self.board[c[0]][c[1]] == p) {
        chn.pieces[c] = true
        for _, dir := range dirs {
          t := Coord{dir[0]+c[0], dir[1]+c[1]}
          if (!proc[t]) {
            proc[t] = true
            q = append(q,t)
          }
        }
      } else if (self.board[c[0]][c[1]] == 0){
        chn.liberty = true
      }
    }
  }
  return chn
}

func (self *Game) GetGroup(x, y int, chn chain)(group){
  g := group{}
  g.eye = true
  proc := make(map[Coord]bool)
  q := make([]Coord, 0, 100)
  c := Coord{x,y}
  dirs := [...]Coord{ {-1,0}, {1,0}, {0, -1}, {0, 1}}

  q = append(q, c)
  proc[c] = true

  for(len(q) > 0){
    c, q = q[len(q)-1], q[:len(q)-1]
    if (c[0] >= 0 && c[0] < 10 && c[1] >= 0 && c[1] < 10 && !chn.pieces[c]){
      g.Coords = append(g.Coords, c)
      for _, dir := range dirs {
        t := Coord{dir[0]+c[0], dir[1]+c[1]}
        if (!proc[t]){
          proc[t] = true
          q = append(q, t)
          if( g.eye && 
              ( t[0] == 10 ||
                t[1] == 10 ||
                (self.sides < 2 && t[0] < 0) ||
                (self.sides == 0 && t[1] < 0))) {
            g.eye = false
          }
        }
      }
    }
  }
  return g
}

func (self *Game) GetGroups(chn chain)([]group) {
  gs := make([]group, 0)
  proc := make(map[Coord]bool)
  for x := 0 ; x<10 ; x++ {
    for y := 0 ; y<10 ; y++ {
      c := Coord{x,y}
      if (!chn.pieces[c] && !proc[c]){
        g := self.GetGroup(x, y, chn)
        gs = append(gs, g)
        for _, c := range g.Coords {
          proc[c] = true
        }
      }
    }
  }
  return gs
}

func (self *Game) CheckForWin(x,y int)(bool, []group, chain) {
  chn := self.GetChain(x,y)
  gs := self.GetGroups(chn)
  c := 0
  for _, g := range gs {
    if (g.eye){
      c++
    }
  }
  return c >= 2, gs, chn
}

func (self *Game) Id()(string){
  s := ""
  for x := 0; x<10; x++ {
    for y := 0; y<10; y++ {
      v := "0"
      if (self.board[x][y] != 0){
        v = "1"
        if (self.board[x][y] != self.turn){
          v = "2"
        }
      }
      s += v
    }
  }
  return s
}

func (self *Game) removeDeadPieces() {
  proc := make(map[Coord]bool)
  for x := 0 ; x<10 ; x++ {
    for y := 0 ; y<10 ; y++ {
      c := Coord{x,y}
      if (self.board[x][y] != 0 && !proc[c]){
        chn := self.GetChain(x,y)
        v := 1
        if (!chn.liberty){
          v = 0
        }
        for c, _ := range chn.pieces{
          proc[c] = true
          self.board[c[0]][c[1]] *= v
        }
      }
    }
  }
}