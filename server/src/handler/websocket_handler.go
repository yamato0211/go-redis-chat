package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/yamato0211/go-redis-chat/src/domain"
)

type WebSocketHandler struct {
	hub *domain.Hub
}

func NewWebSocketHandler(hub *domain.Hub) *WebSocketHandler {
	return &WebSocketHandler{
		hub: hub,
	}
}

func (h *WebSocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := domain.NewClient(ws)
	go client.ReadLoop(h.hub.BroadcastCh, h.hub.UnRegisterCh)
	go client.WriteLoop()
	h.hub.RegisterCh <- client
}
