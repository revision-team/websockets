package main

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"example.com/schema"
	"github.com/gorilla/websocket"
)

func main() {
	// address := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/v1/CLASSICTOKEN/ws"}
	address := url.URL{Scheme: "ws", Host: "localhost:8082", Path: "/ws"}
	// NetWs(address)
	GorillaWs(address, 1)
}

func GorillaWs(address url.URL, ttl time.Duration) {
	for { // no channel end or another condition
		ws, _, err := websocket.DefaultDialer.Dial(address.String(), nil)
		if err != nil {
			log.Fatalln("Dial:", err)

		}

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Fatalln("Read:", err)
				ws.Close()
			}

			var info schema.WsValues
			json.Unmarshal(message, &info)

			// process data
			log.Println("Data:", info)
		}
	}

}
