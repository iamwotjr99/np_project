package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/iamwotjr99/np_project/back/chat_server/Dto"
)

var clients = make(map[*websocket.Conn] bool)
var broadcast = make(chan dto.Message)

var upgrader = websocket.Upgrader{}


func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg dto.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <- broadcast
		
		for client := range clients {
			err := clients.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}