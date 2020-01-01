package main

import (
	"log"
	"net/http"
	"github.com/googolle/go-socket.io"
)

func main() {
	socketioServer, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Sockets logic
	socketioServer.On("connection", func (so socketio.Socket) {
		log.Println("A new user connected.")

		so.Join("chat_room")

		so.On("chat message", func (msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadCastTo("chat_room", "chat message", msg)
		})
	})

	// Start all the server
	http.Handle("/socket.io/", socketioServer)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}