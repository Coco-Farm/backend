// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	branchOff chan struct {
		client    *Client
		jsonBytes []byte
	}

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	frameInterval := time.Second / 60
	ticker := time.NewTicker(frameInterval)
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			players, _ := json.Marshal(&Message{
				Action: "players",
				Data:   g.Players,
			})
			h.broadcast <- players
		}
	}()

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case branch := <-h.branchOff:
			msg := new(Message)
			if err := json.Unmarshal(branch.jsonBytes, msg); err != nil {
				log.Println("json parsing error: ", err.Error())
				continue
			}

			switch msg.Action {
			case "new":
				gCnt++
				p := &Player{ID: gCnt, X: 1, Y: 1}
				g.AddPlayer(p)
				marshaled, _ := json.Marshal(&Message{
					Action: "registered-client",
					Data:   p,
				})

				branch.client.send <- marshaled

			case "event":
				p, ok := msg.Data.(Player)
				marshaled, _ := json.Marshal(&Message{
					Action: "error",
					Data:   "잘못된 json값 입니다.",
				})
				if !ok {
					branch.client.send <- marshaled
					continue
				}

				g.MovePlayer(p.ID, p.X, p.Y)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
