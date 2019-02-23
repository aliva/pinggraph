package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type node struct {
	Name     string
	Host     string
	Port     int
	User     string
	IsRemote bool
}

func (n node) Ping(r node) {
	sshConfig := &ssh.ClientConfig{
		User: n.User,
		Auth: []ssh.AuthMethod{
			getPrivateKey(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", n.Host, n.Port), sshConfig)
	if err != nil {
		fmt.Println("client", err)
		return
	}

	cmd := fmt.Sprintf(
		"ping -c 1 %s -W 1 | tail -1 | awk '{print $4}' | cut -d '/' -f 2",
		r.Host,
	)

	resultItem := pingResult{
		HostName:   n.Name,
		RemoteName: r.Name,
		Counter:    0,
		Result:     -1,
	}

	for {
		pingResultChan <- resultItem
		resultItem.Counter++

		session, err := client.NewSession()
		if err != nil {
			fmt.Println("session", err)
			continue
		}
		outputBytes, err := session.CombinedOutput(cmd)
		if err != nil {
			fmt.Println("Err", err)
			continue

		}
		output := string(outputBytes)
		output = strings.Trim(output, " \n")
		f, err := strconv.ParseFloat(output, 64)

		if err != nil {
			fmt.Println("Err", output)
			continue
		}
		resultItem.Success++
		fmt.Printf("%s => %s, %f - %d/%d\n", n.Name, r.Name, f, resultItem.Success, resultItem.Counter)
		resultItem.Result = f

		time.Sleep(time.Second)
	}
}

func getPrivateKey() ssh.AuthMethod {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	key, err := ioutil.ReadFile(fmt.Sprintf("%s/.ssh/id_rsa", user.HomeDir))
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	return ssh.PublicKeys(signer)
}
