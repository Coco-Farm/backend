package model

// Player 구조체는 본 게임의 사용자를 표현합니다.
type Player struct {
	ID string `json:"id"`
	X  int    `json:"x"`
	Y  int    `json:"y"`

	// PlayerSpeed is ignored by this package.
	PlayerSpeed int `json:"-"`
}
