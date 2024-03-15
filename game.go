package main

var g *Game = NewGame()
var gCnt = 0

type Player struct {
	ID int `json:"id"`
	X  int `json:"x"`
	Y  int `json:"y"`
}

type Game struct {
	Players map[int]*Player
}

func NewGame() *Game {
	return &Game{Players: make(map[int]*Player)}
}

func (g *Game) AddPlayer(p *Player) {
	g.Players[p.ID] = p
}

func (g *Game) RemovePlayer(playerID int) {
	delete(g.Players, playerID)
}

func (g *Game) MovePlayer(playerID, deltaX, deltaY int) {
	player := g.Players[playerID]
	player.X += deltaX
	player.Y += deltaY
}
