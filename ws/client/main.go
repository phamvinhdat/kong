package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	nConn := 0
	err := makeConnection()
	if err != nil {
		time.Sleep(time.Second)
		return
	}

	for {
		nConn++


		fmt.Printf("\x0c%d", nConn)
		time.Sleep(time.Millisecond * 1000)
	}
}

func makeConnection() error {
	ws, _, err := websocket.DefaultDialer.Dial("ws://192.168.64.2/v3/ws", nil)
	if err != nil {
		log.Println("connection err: ", err)
		return err
	}

	go func() {
		for {
			if randInt() == 50 {
				msg := fmt.Sprintf("message: %s", time.Now())
				err = ws.WriteMessage(websocket.BinaryMessage, []byte(msg))
				if err != nil {
					log.Fatal(err)
				}
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

func randInt() int {
	return rand.Intn(100)
}
