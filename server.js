const express = require('express')
const http = require('http')
const socketIO = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = socketIO(server);

// server의 정보 
let players = {};

io.on("connection", (socket) => {
    players[socket.id] = {x: 0, y: 0};
    socket.emit("welcome", players)
        
    io.on("event", (data) => {
        players[socket.id] = data;
    });
});


io.on("disconnect", (socket) => {
    delete players[socket.id];
});

setInterval(() => {
    io.emit("update", players);
}, 1000 / 30);

server.listen(8080, () => {
    console.log("Server listening on port 8080")
});