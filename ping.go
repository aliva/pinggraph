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

// Ping docs
func (host Host) Ping(remote Host) {
	var success int

	sshConfig := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			getPrivateKey(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host.Host, host.Port), sshConfig)
	if err != nil {
		fmt.Println("client", err)
		return
	}

	cmd := fmt.Sprintf(
		"ping -c 1 %s -W 1 | tail -1 | awk '{print $4}' | cut -d '/' -f 2",
		remote.Host,
	)

	resultItem := pingResult{
		HostName:   host.Name,
		RemoteName: remote.Name,
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
		success++
		fmt.Printf("%s => %s, %f - %d/%d\n", host.Name, remote.Name, f, success, counter)
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
