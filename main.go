package main

import (
	"fmt"
	"time"
)

var hosts = make([]Host, 0)

func pinger(h *Host) {
	counter :=0
	for {
		for _, r := range hosts {
			if r.Name != h.Name {
				counter ++
				ping := h.Ping(r.Host)
				fmt.Printf("%s => %s, %d, %f\n", h.Name, r.Name, counter , ping)
			}
			time.Sleep(1)
		}
	}
}

func main() {
	done := make(chan bool)
	hosts = LoadConfig("hosts.yml")
	for _, h := range hosts {
		go pinger(&h)
	}
	<-done
}
