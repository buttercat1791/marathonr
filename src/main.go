package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nhooyr.io/websocket"
)

func handleMessage(msg []interface{}) (err error) {
	msgVerb, ok := msg[0].(string)
	if !ok {
		return fmt.Errorf("expected string as first element of message, got %T", msg[0])
	}

	switch msgVerb {
	case "EVENT":
		// TODO
		break
	case "REQ":
		// TODO
		break
	case "CLOSE":
		// TODO
		break
	default:
		break
	}

	return nil
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection to WebSocket
	conn, wsErr := websocket.Accept(w, r, nil)
	if wsErr != nil {
		defer conn.Close(websocket.StatusInternalError, "Internal Error")
		return
	}

	for {
		// Read a NIP-01 spec-compliant message from the connection.
		msgType, msg, readErr := conn.Read(r.Context())
		if readErr != nil {
			defer conn.Close(websocket.StatusInternalError, "Internal Error")
			return
		}
		if msgType != websocket.MessageText {
			defer conn.Close(websocket.StatusUnsupportedData, "Unsupported Message Type")
			return
		}

		// Deserialize the message into a JSON array.
		var data []interface{}
		jsonErr := json.Unmarshal(msg, &data)
		if jsonErr != nil {
			defer conn.Close(websocket.StatusInternalError, "JSON Deserialization Error")
			return
		}

		// Handle the message.
		handleErr := handleMessage(data)
		if handleErr != nil {
			defer conn.Close(websocket.StatusInternalError, "Nostr Message Handling Error")
			return
		}
	}
}

// Server entry point
func main() {
	// Handle WebSocket requests to root path
	http.HandleFunc("/", handleRoot)
}
