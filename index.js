import express from 'express';
import { createServer } from 'node:http';
import { fileURLToPath } from 'node:url';
import { dirname, join } from 'node:path';
import { Server } from 'socket.io';

const app = express();
const server = createServer(app);
const io = new Server(server);

const __dirname = dirname(fileURLToPath(import.meta.url));

app.get("/", (req, res) => {
    res.sendFile(join(__dirname, "index.html"));
});

// client에서 socket.io() 함수를 실행하면 이 코드가 자동적으로 실행된다. 
io.on('connection', (socket) => {
    socket.on('chat message', (msg) => {
        io.emit("chat message", msg);
    });
  });

server.listen(3000, () => {
    console.log("Server running at http://localhost:3000");
});