package main

// // client가 Connect 요청하고 "newUser"를 emit하는 함수 입니다.
// func HandleNewUser(c socketio.Conn) {
// 	// Map 크기는 5000 x 5000
// 	p := &Player{
// 		ID: c.ID(),
// 		X:  5,
// 		Y:  5,
// 	}
// 	game.AddPlayer(p)
// 	c.Emit("newUser", p)
// }

// // client의 연결 접속이 끊겼을 때 처리하는 함수 입니다.
// func HandleDisconnectUser(c socketio.Conn) {
// 	defer c.Close()

// 	game.RemovePlayer(c.ID())
// 	fmt.Printf("[Disconnected] Disconnected User: %s (%s)\n", c.RemoteAddr(), c.ID())
// }
