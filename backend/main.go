package main

import (
	"fmt"
	"github.com/mxstrv/go-react-chat/pkg/websocket"
	"net/http"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat application v0.2 starting")
	setupRoutes()
	http.ListenAndServe("localhost:8080", nil)
}
