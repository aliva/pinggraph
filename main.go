package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileFlag := flag.String("f", "nodes.yml", "hosts file")
	flag.Parse()

	nodes := loadNodes(*fileFlag)

	for _, n := range nodes {
		for _, r := range nodes {
			if n.Name != r.Name && n.IsRemote == false {
				go n.Ping(r)
			}
		}
	}

	go webSocketHandler()

	fmt.Println("Open http://127.0.0.1:8000 in browser")

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
