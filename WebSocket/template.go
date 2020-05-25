package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type connection struct {
	ws *websocket.Conn
	send chan []byte
	h *hub
}

type hub struct {
	connections map[*connection]bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func main() {
	
	log.SetPrefix("( superlog )== ")
	
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8080", nil)
	
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")
}

func wsHandler(response http.ResponseWriter, request *http.Request) {
	
	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	
	for {
		messageType, byteMessage, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		
		output := processMessage(byteMessage)
		
		if err := conn.WriteMessage(messageType, output); err != nil {
			log.Println(err)
			return
		}
	}
	
}


func processMessage(input []byte) []byte {

	fmt.Println("byteMessage: ")
	fmt.Println(input)
	
	fmt.Println("\nstring(byteMessage)")
	fmt.Println(string(input))
	
	return input

}
