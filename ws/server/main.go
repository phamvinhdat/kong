package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", wsHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":9501", r))
}

var upgrader websocket.Upgrader

func wsHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("upgrader err: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go func() {
		defer conn.Close()
		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read msg err: ", err)
				return
			}

			log.Println("receive msg, echo: ", string(msg))
			err = conn.WriteMessage(t, msg)
			if err != nil {
				log.Println("write msg err: ", err)
				return
			}
		}
	}()
}
