package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"example.com/schema"
	"github.com/gorilla/websocket"
)

func main() {
	// address := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/v1/CLASSICTOKEN/ws"}
	address := url.URL{Scheme: "ws", Host: "localhost:8081", Path: "/ws"}

	// channel to exchange info between proxyClient and proxyServer
	data := make(chan []byte, 1024)

	go wsProxy(address, 1, data)

	http.HandleFunc("/ws", wsListener(data))
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func wsProxy(address url.URL, ttl time.Duration, data chan<- []byte) {
	for { // no channel end or another condition
		ws, _, err := websocket.DefaultDialer.Dial(address.String(), nil)
		if err != nil {
			log.Println("Dial:", err)
			<-time.After(ttl * time.Second)
		} else {
			for {
				_, message, err := ws.ReadMessage()
				if err != nil {
					log.Println("Read:", err)
					<-time.After(ttl * time.Second)
					ws.Close()
					break
				}

				var info schema.WsValues
				json.Unmarshal(message, &info)

				// process data
				log.Println("Proxy:", info)
				data <- message
			}
		}
	}
}

var upgrader = websocket.Upgrader{}
var no int64 = 1

func wsListener(data <-chan []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mutex := sync.Mutex{}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer ws.Close()

		go func() {
			for {
				<-time.After(5 * time.Second)
				mutex.Lock()
				ws.WriteMessage(websocket.BinaryMessage, []byte(`{"topic": "HeartBeat"}`))
				mutex.Unlock()
			}
		}()

		for {
			message := <-data
			
			mutex.Lock()
			err = ws.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
			mutex.Unlock()

			log.Println("written:", string(message))
			<-time.After(time.Second)
			no += 1
		}
	}
}
