package main

import (
	"log"
	"net/http"
//	"os"
	"os/exec"
	"fmt"

	"github.com/googollee/go-socket.io"
)

// This is our "Streaming" object for streaming file writes
// To another function in an async manner
type Stream struct {
	Chan chan []byte
}

// Implement io.Writer
func (s Stream) Write(p []byte) (n int, err error) {
	// Write to our channel
	s.Chan <- p
	// For Debug:
//	n, err = os.Stdout.Write(p);
//	os.Stdout.Sync()
//	if err != nil {
//		return
//	}
	// Return how many bytes we sent to the channel,
	// and no error
	n = len(p)
	err = nil
	return
}

func socketServer() *socketio.Server {
	// Make our byte channel
	zoink := make(chan []byte)

	// Create our Socket.io server
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// On socket.io connection, do stuff
	server.On("connection", func(so socketio.Socket) {

		// Print that we connected and tell the world
		log.Println("on connection")
		so.Join("chat")

		// Create our Stream object and write the output of
		// proc_test.sh to it.
		s := Stream{Chan:zoink}
		cmd := exec.Command("./proc_test.sh")
		cmd.Stdout = s
		cmd.Start()

		// Go routine to listen for output from the command
		// and send it to all connected users
		go func(){
			for {
				tmp := fmt.Sprintf("%q", <-zoink)
				fmt.Printf("Printed: %s\n", tmp)
				so.BroadcastTo("chat", "chat message", tmp)
			}
		}()

		// On message, send message to everyone
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})

		// Echo on disconnection
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})

		// Emit "welcome" to the world after all initialization has completed
		so.Emit("welcome")
	})

	// When there is an error in socket.io, report it.
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	// Return the server
	return server
}
func main() {
	// Create the socket.io server and initialize it
	server := socketServer()

	// Set up http server and routing
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
