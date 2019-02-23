package main

type pingResult struct {
	HostName   string
	RemoteName string
	Counter    int
	Success    int
	Result     float64
}

var pingResultChan = make(chan pingResult)
