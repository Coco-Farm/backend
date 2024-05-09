package model

import "errors"

var (
	ErrInvalidDirection = errors.New("invalid direditon input")
)

// Player 구조체는 본 게임의 사용자를 표현합니다.
type Player struct {
	ID string `json:"id"`
	X  int    `json:"x"`
	Y  int    `json:"y"`

	// Speed is ignored by this package.
	Speed int `json:"-"`
}

func (p *Player) Move(direction string) error {
	switch direction {
	case "up":
		p.Y += 1 * p.Speed
	case "down":
		p.Y -= 1 * p.Speed
	case "left":
		p.X -= 1 * p.Speed
	case "right":
		p.X += 1 * p.Speed
	default:
		return ErrInvalidDirection
	}

	return nil
}
