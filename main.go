package main

import (
	"backend/internal/domain"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	// player1 := domain.NewPlayer(uuid.New(), "JunBumHan")

	// // 사용자가 들어왔다?
	// core.AddPlayer(player1)

	// // 사용자가 움직였다면?
	// core.MovePlayer(player1.ID, 10, 5)

	// // 사용자가 없어진다면?
	// core.DeletePlayer(player1.ID)

	// 서버 네트워크 생성
	net := domain.NewNet(domain.NewCore())

	// 각각의 사용자들에게 broadcast
	go net.Deploy(domain.DefaultTickRate)

	http.HandleFunc("GET /ws/{name}", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("websocket으로 Upgrade하던 중 오류가 발생했습니다.\nerror: %v\n", err)
			return
		}

		net.AddUser(uuid.New(), r.PathValue("name"), domain.WrapConn(conn))
	})

	log.Fatal(http.ListenAndServe(":9190", nil))

}
