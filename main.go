package main

import (
	"flag"
	"log"
	"net/http"
)

var hosts = make([]Host, 0)

func main() {
	fileFlag := flag.String("f", "hosts.yml", "hosts file")
	flag.Parse()

	hosts = LoadConfig(*fileFlag)

	for _, h := range hosts {
		for _, r := range hosts {
			if h.Name != r.Name && h.IsRemote == false {
				go h.Ping(r)
			}
		}
	}

	go webClientHandler()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
