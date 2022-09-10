package server

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func CreateServer() {
	server := socketio.NewServer(nil)

	server.OnConnect("/chat", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Emit("joined")
		return nil
	})

	server.OnEvent("/chat", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) {
		s.SetContext(msg)
		s.Emit("msg", "recv "+msg)
		fmt.Println("recv " + msg)
	})

	server.OnEvent("/chat", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnEvent("/chat", "chujem", func(s socketio.Conn) {
		s.SetContext("")
		s.Emit("wdupe", "recv")
		fmt.Println("ping")
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/chat", func(s socketio.Conn, reason string) {
		fmt.Printf("Client %s left: %s", s.ID(), reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
