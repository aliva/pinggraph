package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type pingResult struct {
	HostIP   string
	RemoteIP string
	ID       int
	Result   float64
}

type webClient struct {
	conn *websocket.Conn
}

var webClientSlice = make([]webClient, 0)
var pingResultChan = make(chan pingResult)

func webClientHandler() {
	for {
		result := <-pingResultChan
		fmt.Println(result)

		for i, client := range webClientSlice {
			err := client.conn.WriteJSON(result)
			if err != nil {
				fmt.Println("Connection closed")
				client.conn.Close()
				webClientSlice = append(webClientSlice[:i], webClientSlice[i+1:]...)
			}
		}
		time.Sleep(time.Second)
	}
}
