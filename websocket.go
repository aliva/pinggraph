package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

var webSocketConnections = make([]*websocket.Conn, 0)

func webSocketHandler() {
	for {
		r := <-pingResultChan

		for i, c := range webSocketConnections {
			err := c.WriteJSON(r)
			if err != nil {
				fmt.Println("Connection closed")
				c.Close()
				webSocketConnections = append(
					webSocketConnections[:i],
					webSocketConnections[i+1:]...,
				)
			}
		}
	}
}
