package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

type msg struct {
	Bartosz int
	Louis   string
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go echo(conn)
}

func echo(conn *websocket.Conn) {
	for {
		//m := msg{}

		a, b, c := conn.ReadMessage()
		// if err != nil {
		// 	fmt.Println("Error reading json.", err)
		// }

		fmt.Println("Got message: ", a, b, c)

		ma := struct {
			haha string
			bebe string
			omg  float64
			nooo int
		}{
			"heee",
			"haaaa",
			3.14,
			69,
		}

		// maa := msg{
		// 	5,
		// 	"louis",
		// }

		if err := conn.WriteJSON(ma); err != nil {
			fmt.Println(err)
		}
	}
}
