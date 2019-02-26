package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aliva/pinggraph/pinggraph"
)

func main() {
	fileFlag := flag.String("f", "nodes.yml", "hosts file")
	flag.Parse()

	nodes := pinggraph.LoadNodes(*fileFlag)

	for _, n := range nodes {
		for _, r := range nodes {
			if n.Name != r.Name && n.IsRemote == false {
				go n.Ping(r)
			}
		}
	}

	go pinggraph.WebSocketHandler()

	fmt.Println("Open http://127.0.0.1:8000 in browser")

	http.HandleFunc("/", pinggraph.ServeHome)
	http.HandleFunc("/ws", pinggraph.ServeWs)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
