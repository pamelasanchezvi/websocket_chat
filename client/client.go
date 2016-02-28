package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
)

func CommandLine() (error, string) {
	bio := bufio.NewReader(os.Stdin)
	if line, hasMore, err := bio.ReadLine(); hasMore == false {
		if err != nil {
			return err, ""
		} else {
			return nil, string(line)
		}
	}
	return nil, ""
}

func Close(ws *websocket.Conn) {
	ws.Close()
}

func RegisterAndListen(dest, src, Username, Password string) {

	config, err := websocket.NewConfig(dest, src)
	if err != nil {
		fmt.Println("NewConfig failed ! :(")

	}
	message := Username + ":" + Password
	usrpasswd := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(usrpasswd, []byte(message))
	config.Header = make(http.Header)
	config.Header.Set("Authorization", "Basic "+string(usrpasswd))
	fmt.Println("Header set is: " + config.Header.Get("Authorization"))
	webs, err := websocket.DialConfig(config)
	if err != nil {
		if webs == nil {
			fmt.Println(" websocket is nil")

		} else {
			fmt.Println("some error")
		}
	} else {
		for {
			err, line := CommandLine()
			if err == nil && line != "quit" {
				fmt.Println("Sent")
				websocket.Message.Send(webs, line)
			} else if err != nil {
				fmt.Println("Error, please type again")
				continue
			}
			if line == "quit" {
				break
			}
		}
	}
	Close(webs)
}

func main() {
	// set IP:port of server and client here
	dest := "ws://192.168.1.105:9000"
	src := "ws://172.16.162.131"
	Username := "administrator"
	Password := "p"

	RegisterAndListen(dest, src, Username, Password)
}
