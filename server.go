package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	username string
	room     string
}

var clients = make(map[*websocket.Conn]Client)
var passwords = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	fmt.Println("Launching server...")

	http.HandleFunc("/ws", handleConnections)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	var username, password, room string
	err = conn.ReadJSON(&username)
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	err = conn.ReadJSON(&password)
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	if validPassword, ok := passwords[username]; !ok || validPassword != password {
		fmt.Println("Invalid username or password")
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid username or password"))
		return
	}

	err = conn.ReadJSON(&room)
	if err != nil {
		fmt.Println("Error reading room:", err)
		return
	}

	client := Client{conn, username, room}
	mu.Lock()
	clients[conn] = client
	mu.Unlock()
	fmt.Printf("User %s connected to room %s\n", username, room)

	for {
		var message string
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("Error reading message:", err)
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		formattedMessage := fmt.Sprintf("%s: %s\n", username, message)
		fmt.Print(formattedMessage)

		saveMessageToFile(room, formattedMessage)

		mu.Lock()
		for _, cl := range clients {
			if cl.conn != conn && cl.room == room {
				err := cl.conn.WriteMessage(websocket.TextMessage, []byte(formattedMessage))
				if err != nil {
					fmt.Println("Error sending message:", err)
				}
			}
		}
		mu.Unlock()
	}
}

func saveMessageToFile(roomName, message string) {
	filePath := fmt.Sprintf("messages/%s.txt", roomName)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(message); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
