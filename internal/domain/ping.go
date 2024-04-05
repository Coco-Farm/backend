package domain

// Ping 구조체는 본 게임에 접속한 클라이언트와 대화하는 테플릿입니다. Pakcet과 똑같다고 보면 됩니다.
type Ping struct {
	Event   string `json:"event"`
	Payload any    `json:"payload"`
}

// Ping은 Ping 구조체의 생성자 입니다.
func NewPing(e string, p any) *Ping {
	return &Ping{
		Event:   e,
		Payload: p,
	}
}

// InputMovedPayload 구조체는 Payload 종류의 클라이언트가 EventMove일 때 전달하는 값의 구조 입니다.
type InputMovedPayload struct {
	X int `json:"x"`
	Y int `json:"y"`
}
