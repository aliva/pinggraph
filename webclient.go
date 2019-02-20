package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type webClient struct {
	conn *websocket.Conn
}

var webClients = make([]webClient, 0)

func webClientHandler() {
	for {
		for i, client := range webClients {
			data := map[string]string{
				"key": time.Now().String(),
			}
			err := client.conn.WriteJSON(data)
			if err != nil {
				fmt.Println("Connection closed")
				client.conn.Close()
				webClients = append(webClients[:i], webClients[i+1:]...)
			}
		}
		time.Sleep(time.Second)
	}
}
