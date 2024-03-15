package main

type Player struct {
	ID string `json:"id"`
	X  int    `json:"x"`
	Y  int    `json:"y"`
}

type Game struct {
	Players map[string]*Player
}

func NewGame() *Game {
	return &Game{Players: make(map[string]*Player)}
}

func (g *Game) AddPlayer(p *Player) {
	g.Players[p.ID] = p
}

func (g *Game) RemovePlayer(playerID string) {
	delete(g.Players, playerID)
}

func (g *Game) MovePlayer(playerID string, deltaX, deltaY int) {
	player := g.Players[playerID]
	player.X += deltaX
	player.Y += deltaY
}
