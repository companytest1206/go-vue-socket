package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

// testing branch with new chnages

func main() {
	server, err := socketio.NewServer(nil)
	if(err != nil) {
		log.Fatal(err)
	}
	
	server.On("connection", func(so socketio.Socket) {
		log.Println("Someone Connected!!")
		so.Join("chat_room")
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat_room", "chat message", msg)
		})
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Serving at :5000 rapbeh...testing")
	log.Fatal(http.ListenAndServe(":5000", nil))
}