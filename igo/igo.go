package igo

type Game struct {
  turn int
  board [10][10]int
  sides int
}

type Coord [2]int

type group struct {
  Coords []Coord
  eye bool
}

type chain map[Coord]bool

func (self *Game) Board(x,y int)(int) { return self.board[x][y] }
func (self *group) Eye()(bool) { return self.eye }

func (self *Game) Init(sides int)(){
  self.turn = 1
  self.sides = sides
  if (self.sides > 2){
    self.sides = 2
  }
}

func (self *Game) Move(x, y int)(bool, []group, chain){
  win := false
  gs := make([]group,0,0)
  chn := chain{}
  if (self.board[x][y] == 0){
    self.board[x][y] = self.turn
    win, gs, chn = self.CheckForWin(x,y)
    self.turn = 3 - self.turn
  }
  return win, gs, chn
}

func (self *Game) GetChain(x,y int)(chain){
  p := self.board[x][y]
  proc := make(map[Coord]bool)
  q := make([]Coord, 0, 100)
  c := Coord{x,y}
  chn := chain{} //? make(chain)
  dirs := [...]Coord{ {-1,0}, {1,0}, {0, -1}, {0, 1}}

  q = append(q,c)
  proc[c] = true

  for len(q) > 0 {
    c, q = q[len(q)-1], q[:len(q)-1]
    if (c[0] >= 0 && c[0] < 10 && c[1] >= 0 && c[1] < 10 && self.board[c[0]][c[1]] == p){
      chn[c] = true
      for _, dir := range dirs {
        t := Coord{dir[0]+c[0], dir[1]+c[1]}
        if (!proc[t]) {
          proc[t] = true
          q = append(q,t)
        }
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
    if (c[0] >= 0 && c[0] < 10 && c[1] >= 0 && c[1] < 10 && !chn[c]){
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
      if (!chn[c] && !proc[c]){
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