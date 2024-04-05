package domain

import (
	"github.com/google/uuid"
)

// Player는 마법사의 키우기의 가상 캐릭터의 구성요소를 정의한 구조체 입니다.
type Player struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	X    int       `json:"x"`
	Y    int       `json:"y"`
}

// NewPlayer 함수는 Player를 생성하여 반환해주는 함수 입니다.
func NewPlayer(id uuid.UUID, name string) *Player {
	return &Player{ID: id, Name: name}
}
