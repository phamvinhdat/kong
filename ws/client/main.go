package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	nConn := 0
	for {
		nConn++
		err := makeConnection()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		fmt.Printf("\x0c%d", nConn)
		time.Sleep(time.Millisecond * 10)
	}
}

func makeConnection() error {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8000/v3/websocket", nil)
	if err != nil {
		log.Println("connection err: ", err)
		return err
	}

	go func() {
		for {
			msg := fmt.Sprintf("message: %s", time.Now())
			err = ws.WriteMessage(websocket.BinaryMessage, []byte(msg))
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			printMsg(string(message), false)
		}
	}()

	return nil
}

func printMsg(msg string, print bool) {
	if print {
		log.Printf("recv: %s\n", msg)
	}
}
