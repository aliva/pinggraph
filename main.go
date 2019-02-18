package main

var hosts = make([]Host, 0)

func pinger() {
	for _, h := range hosts {
		for _, r := range hosts {
			if r.Name != h.Name {
				go h.Ping(r)
			}
		}
	}
}

func main() {
	done := make(chan bool)
	hosts = LoadConfig("hosts.yml")
	pinger()
	<-done
}
