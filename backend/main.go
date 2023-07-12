package main

import (
	"fmt"
	"github.com/mxstrv/go-react-chat/pkg/websocket"
	"net/http"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket endpoint hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Realtime chat v0.3")
	setupRoutes()
	http.ListenAndServe("localhost:8080", nil)
}
