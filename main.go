package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/valyala/fastjson"
)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

var p fastjson.Parser = fastjson.Parser{}
var g *Game = NewGame()

func main() {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnEvent("/", "playerMovement", func(c socketio.Conn, msg string) {
		v, _ := p.Parse(msg)
		g.MovePlayer(
			string(v.GetStringBytes("id")),
			v.GetInt("x"),
			v.GetInt("y"),
		)
	})

	server.BroadcastToNamespace()

	// server.BroadcastToNamespace("/", "state", func(c socketio.Broadcast)) {}

	// server.BroadcastToRoom("/", )

	// server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
	// 	log.Println("notice: ", msg)
	// 	s.Emit("reply", "have "+string(msg))
	// })

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		log.Println("chat: ", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(c socketio.Conn, err error) {
		log.Println("meet error: ", err)
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatal("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))

	log.Println("Serving at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
