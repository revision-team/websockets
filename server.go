package main
 
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", chaoticWs)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

var upgrader = websocket.Upgrader{}
var no int64 = 1

func chaoticWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	for {
		message := fmt.Sprintf(`{"topic":"myTopic%d", "payload": {}}`, no)
		err = ws.WriteMessage(websocket.BinaryMessage, []byte(message))
		if err != nil {
			log.Println("write:", err)
			break
		}
		log.Println("written:", message)
		no += 1
		<-time.After(time.Second)
	}
}
