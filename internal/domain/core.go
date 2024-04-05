package domain

import (
	"log"

	"github.com/google/uuid"
)

// Core는 게임의 핵심을 담당하는 구조체입니다.
// 본 구조체에서 룬석, 플레이어, 총알 등 게임의 구성요소가 모이게 되고
// 게임내 필요한 연상을 수행하게 됩니다.
type Core struct {
	Players map[uuid.UUID]*Player
}

// NewCore 함수는 Core 구조체의 생성자 입니다.
func NewCore() *Core {
	return &Core{Players: make(map[uuid.UUID]*Player)}
}

// AddPlayers 함수는 Players 멤버변수의 값을 추가하는 함수입니다.
func (c *Core) AddPlayer(player *Player) {
	c.Players[player.ID] = player
}

// DeletePlayer 함수는 Players 멤버변수의 값을 삭제하는 함수입니다.
func (c *Core) DeletePlayer(playerID uuid.UUID) {
	delete(c.Players, playerID)
}

// MovePlayer 함수는 Player의 움직임을 갱신하는 함수 입니다.
func (c *Core) MovePlayer(playerID uuid.UUID, x int, y int) {
	p, ok := c.Players[playerID]
	if ok {
		log.Printf("존재하지 않은 Player에 접근을 시도하였습니다.\n존재하지 않은 player의 ID playerID: %v 입니다.\n", playerID)
		return
	}
	p.X = x
	p.Y = y
}
