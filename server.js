const express = require('express')
const http = require('http')
const socketIO = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = socketIO(server, {path: '/socket.io'});

// server의 정보 
let players = {};

io.on("connection", (socket) => {
    console.log("New player: ", socket.id)
    players[socket.id] = {x: 0, y: 0};
    socket.emit("welcome", players)
        
    socket.on("event", (data) => {
        console.log("event detected | player : ", socket.id);
        players[socket.id] = data;
    });

    socket.on("disconnect", () => {
        console.log("logout player: ", socket.id);
        delete players[socket.id];
    });
});

setInterval(() => {
    io.emit("update", players);
}, 1000 / 30);

server.listen(8080, () => {
    console.log("Server listening on port 8080")
});