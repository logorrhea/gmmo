package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
)

func socketHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade the HTTP connection to a websocket
	// @TODO: Check origin?
	log.Println("Attempting to upgrade connection to websocket....")
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Fatal("websocket.Upgrade:" + err.Error())
	}
	//_, ok := err.(websocket.HandshakeError)
	//if ok {
	//http.Error(w, "Not a websocket handshake", 400)
	//return
	//} else if err != nil {
	//return
	//}
	log.Println("Client connected")

	for {
		messageType, r, err := ws.NextReader()
		if err != nil {
			return
		}
		w, err := ws.NextWriter(messageType)
		if err != nil {
			log.Fatal("NextWriter:" + err.Error())
		}
		if _, err := io.Copy(w, r); err != nil {
			log.Fatal("io.Copy:" + err.Error())
		}
		if err := ws.Close(); err != nil {
			log.Fatal("Close: " + err.Error())
		}
	}
}

func main() {

	// 1. Load the initializers

	// 2. Load data models

	// 3. Load game maps data
	mapFolder, err := os.Open("maps")
	if err != nil {
		fmt.Println("Failed to open maps directory")
	} else {
		fmt.Printf("Opened folder %s\n", mapFolder.Name())
	}

	// 4. Initiate the server and listen to the internets
	http.HandleFunc("/", socketHandler)
	err = http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	// 5. All the server logic
}
