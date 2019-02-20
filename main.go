package main

import (
	"log"
	"net/http"
)

var hosts = make([]Host, 0)

func pinger() {
	for _, h := range hosts {
		for _, r := range hosts {
			if h.Name != r.Name && h.IsRemote == false {
				go h.Ping(r)
			}
		}
	}
}

func main() {
	hosts = LoadConfig("hosts.yml")
	// pinger()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
