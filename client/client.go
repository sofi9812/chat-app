package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run client.go <username> <password> <room>")
		return
	}

	username := os.Args[1]
	password := os.Args[2]
	room := os.Args[3]

	fmt.Println("Launching client...")

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	conn.WriteJSON(username)
	conn.WriteJSON(password)
	conn.WriteJSON(room)

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading from server:", err)
				return
			}
			fmt.Print(string(message))
			if string(message) == "Invalid username or password" {
				log.Fatal()
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("U are in the room ", room, " U can start chatting now!\n")
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		conn.WriteJSON(text)
	}
}
