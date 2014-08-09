package main

import (
	"log"
	"net/http"
//	"os"
	"os/exec"
	"fmt"

	"github.com/googollee/go-socket.io"
)

type Stream struct {
	Chan chan []byte
}

func (s Stream) Write(p []byte) (n int, err error) {
	//pipe := []byte{'|'}
	//_, err = os.Stdout.Write(pipe);
	//if err != nil {
	//	return
	//}
	//b := append(pipe[:], p[:]...)
	//for i := 0; i < len(p); i++ {
	//}
	s.Chan <- p
//	n, err = os.Stdout.Write(p);
//	os.Stdout.Sync()
//	if err != nil {
//		return
//	}
	n = len(p)
	err = nil
	return
}

func tester() {
	log.Println("Eep!")
}

func socketServer() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	zoink := make(chan []byte)
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Emit("welcome")
		so.Join("chat")
		s := Stream{Chan:zoink}
		cmd := exec.Command("./proc_test.sh")
		cmd.Stdout = s
		cmd.Start()

		go func(){
			for {
				tmp := fmt.Sprintf("%q", <-zoink)
				fmt.Println("Printed: %s", tmp)
				so.BroadcastTo("chat", "chat message", tmp)
			}
		}()
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
		so.Emit("welcome")
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	return server
}
func main() {
	server := socketServer()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
