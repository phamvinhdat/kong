package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	nConnectionKey  = "n-connection"
	nCConnectionKey = "n-closed-connection"
)

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/ws", wsHandler).Methods(http.MethodGet)
	r.HandleFunc("/ws-info", getConnectionInfo).Methods(http.MethodGet)
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
		defer func() {
			log.Println("connection closed!!!")
			conn.Close()
		}()

		err = increaseNConn(nConnectionKey)
		if err != nil {
			log.Println(err)
			return
		}

		defer func() {
			_ = decreaseNConn()
			_ = increaseNConn(nCConnectionKey)
		}()

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

func getConnectionInfo(w http.ResponseWriter, req *http.Request) {
	nc, err := rdb.Get(context.Background(), nConnectionKey).Int()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: " + err.Error()))
		return
	}

	ncc, err := rdb.Get(context.Background(), nCConnectionKey).Int()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: " + err.Error()))
		return
	}

	info := struct {
		Connection       int `json:"connection"`
		ClosedConnection int `json:"closed_connection"`
	}{
		Connection:       nc,
		ClosedConnection: ncc,
	}
	b, _ := json.Marshal(info)

	_, _ = w.Write(b)
}

func increaseNConn(key string) error {
	_, err := rdb.Incr(context.Background(), key).Result()
	if err != nil {
		return errors.New("increase failed: " + err.Error())
	}

	return nil
}

func decreaseNConn() error {
	_, err := rdb.Decr(context.Background(), nConnectionKey).Result()
	if err != nil {
		return errors.New("decrease failed: " + err.Error())
	}

	return nil
}
