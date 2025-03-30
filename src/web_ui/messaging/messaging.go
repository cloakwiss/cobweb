package messaging

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var activeConn *websocket.Conn

type websocketWriter struct{}

func (w *websocketWriter) Write(p []byte) (n int, err error) {
	if activeConn != nil {
		err := activeConn.WriteMessage(websocket.TextMessage, p)
		if err != nil {
			log.Println("WebSocket error:", err)
			activeConn.Close()
			activeConn = nil
		}
	}
	return len(p), nil
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	activeConn = conn
	log.Println("WebSocket connected.")

	// Keep connection alive until closed
	for {
		if _, _, err := conn.NextReader(); err != nil {
			log.Println("WebSocket disconnected.")
			activeConn = nil
			break
		}
	}
}

func StartMessaging(w http.ResponseWriter) {
	log.Println("Waiting for WebSocket...")

	// Wait up to 5 seconds for WebSocket connection
	for range 50 {
		if activeConn != nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	if activeConn == nil {
		http.Error(w, "No WebSocket connection", http.StatusBadRequest)
		return
	}

	log.Println("WebSocket active. Starting...")

	// Redirect stdout to WebSocket & terminal
	// neccesary to use log.Prints to send messages
	multiWriter := io.MultiWriter(os.Stdout, &websocketWriter{})
	log.SetOutput(multiWriter)
}

func EndMessaging() {
	activeConn.Close()
	activeConn = nil
}
