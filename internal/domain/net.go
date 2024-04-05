package domain

import (
	"log"
	"time"

	"github.com/google/uuid"
)

const DefaultTickRate = 30

// Net의 구조체는 클라이언트와 연결음 담당하는 구조체 입니다.
type Net struct {
	*Core
	Conns map[uuid.UUID]*Conn

	// Register requests from the clients.
	EventRegister chan func() (*Conn, *Player)
	// Unregister requests from clients.
	EventUnregister chan uuid.UUID
}

func NewNet(c *Core) *Net {
	return &Net{
		Core:  c,
		Conns: make(map[uuid.UUID]*Conn),

		// register: make(chan struct {
		// 	ID   uuid.UUID
		// 	Conn *Conn
		// }),

		EventRegister:   make(chan func() (*Conn, *Player)),
		EventUnregister: make(chan uuid.UUID),
	}
}

func (n *Net) AddConn(conn *Conn) {
	n.Conns[conn.ID] = conn
}

func (n *Net) DeleteConn(connID uuid.UUID) {
	delete(n.Conns, connID)
}

// 새로운 유저가 들어왔을 떄 실행되느 함수 입니다.
func (n *Net) GenerationPlayer(c *Conn, p *Player) {
	n.EventRegister <- func() (*Conn, *Player) { return c, p }
}

func (n *Net) BroadCast(data any, ignores ...uuid.UUID) {
	var ignore uuid.UUID
	if len(ignores) > 0 {
		ignore = ignores[0]
	}

	for id, conn := range n.Conns {
		if id != ignore {
			if err := conn.WriteJSON(data); err != nil {
				log.Printf("broadCast를 진행하던 중 에러가 발생하였습니다.\n오류가 발생한 연결 소켓: %s:%v\nerror: %v\n", id.String(), conn, err)
			}
		}
	}
}

func (n *Net) Deploy(tickRate int) {
	ticker := time.Tick(time.Second / time.Duration(tickRate))

	for range ticker {
		n.BroadCast(NewPing("broadcast", n.Core.Players))
	}
}

func (n *Net) Run() {
	go n.Deploy(DefaultTickRate)

	for {
		select {
		case f := <-n.EventRegister:
			c, p := f()
			n.AddConn(c)
			n.AddPlayer(p)

		case id := <-n.EventUnregister:
			delete(n.Conns, id)
			n.DeletePlayer(id)
		}
	}
}
