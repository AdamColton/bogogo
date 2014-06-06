package igo

type Game struct {
  turn int
  board [10][10]int
  sides int
}

type Coord [2]int

type group struct {
  player int
  Coords []Coord
}

func (self *Game) Init()(){
  self.turn = 1
}

func (self *Game) Move(x, y int)(){
  if (self.board[x][y] == 0){
    self.board[x][y] = self.turn
    self.turn = 3 - self.turn
  }
}

func (self *Game) GetGroup(x, y, player int)(group){
  g := group{}
  g.player = player
  processed := make(map[Coord]bool)
  q := make([]Coord, 0)
  i := Coord{x,y}
  dirs := [...]Coord{ {-1,0}, {1,0}, {0, -1}, {0, 1}}

  q = append(q, i)
  processed[i] = true

  for(len(q) > 0){
    i, q = q[len(q)-1], q[:len(q)-1]
    if (i[0] >= 0 && i[0] < 10 && i[1] >= 0 && i[1] < 10 && self.board[i[0]][i[1]] != player){
      g.Coords = append(g.Coords, i)
      for _, dir := range dirs {
        t := Coord{dir[0]+i[0], dir[1]+i[1]}
        if (!processed[t]){
          processed[t] = true
          q = append(q, t)
        }
      }
    }
  }
  return g
}