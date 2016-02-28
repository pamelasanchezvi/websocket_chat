package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func RegisterAndListen(clientMsg string) {
	ServerUsername := "administrator"
	ServerPassword := "piopio"

	config, err := websocket.NewConfig("ws://192.168.1.105:9000", "ws://172.16.162.131")
	if err != nil {
		fmt.Println("NewConfig failed ! :(")

	}
	message := ServerUsername + ":" + ServerPassword
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
		//	msgMarshalled, err := proto.Marshal(clientMsg)
		//		err != nil{
		//fmt.Println("marshalling error")
		//		}
		fmt.Println("No error, message sending")
		websocket.Message.Send(webs, clientMsg)
	}
}

func main() {
	RegisterAndListen("I am Sam")
}
