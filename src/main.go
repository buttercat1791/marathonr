package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Server entry point
func main() {
	// Handle WebSocket requests to root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade connection to WebSocket
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "Internal Error")

		// Set timeout
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		// Read request JSON
		var v interface{}
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			fmt.Println(err)
			return
		}

		// TODO: Handle "EVENT" request

		// TODO: Handle "REQ" request

		// TODO: Handle "CLOSE" request

		conn.Close(websocket.StatusNormalClosure, "")
	})
}
