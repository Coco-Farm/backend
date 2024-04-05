package domain

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Conn은 websocket의 인스턴스의 Warp하는 구조체 입니다.
type Conn struct {
	ID uuid.UUID
	*websocket.Conn
}

// WrapConn 함수는 *websocket.Conn을 Wrap합니다.
func WrapConn(id uuid.UUID, target *websocket.Conn) *Conn {
	return &Conn{
		ID:   id,
		Conn: target,
	}
}

// ReadPump는 client가 전송한 값을 대신 전달해주는 매개체 수행 함수입니다.
func (c *Conn) ReadPump() {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("client의 값을 수신하던 중 에러 발생: %v\n", err)
			}
			break
		}

		// 아마 여기서 message를 해석하여 어떤  event인지에 따라서 payload 값을 parsing 해야 할듯
	}
}
