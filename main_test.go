package main

import (
	"testing"
	//	"time"
	"fmt"
	//	"os"
	"bufio"
	"bytes"
	. "github.com/franela/goblin"
	"os/exec"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	//	"github.com/kvz/logstreamer"
	//	"github.com/oguzbilgic/socketio"
	//	"net/http"
	//	"net/http/httptest"
	//	"net/url"
	//	"strings"
	//	sio "github.com/googollee/go-socket.io"
)

func dotWriter(c <-chan []byte, buf *bufio.Writer) {
	var nc []byte = []byte{' '}
	for {
		nc = <-c
		//if string(nc[:]) == "1." {
		//	os.Stdout.Write([]byte{'*'})
		//}
		buf.Write(nc)
		buf.Flush()
	}
}

func TestChannels(t *testing.T) {
	g := Goblin(t)
	setupLoggers()
	g.Describe("Channels", func() {
		g.It("Channels should send data successfully", func() {
			cmd := exec.Command("./proc_test.sh")
			//stdout, err := cmd.StdoutPipe()
			Chan := make(chan []byte)
			Chan2 := make(chan []byte)
			var buf *bytes.Buffer = new(bytes.Buffer)
			var buf2 *bytes.Buffer = new(bytes.Buffer)
			go dotWriter(Chan, bufio.NewWriter(buf))
			go dotWriter(Chan2, bufio.NewWriter(buf2))
			sout := Stream{Chan: Chan}
			serr := Stream{Chan: Chan2}
			cmd.Stdout = sout
			cmd.Stderr = serr
			cmd.Start()
			cmd.Wait()
			scanner := bufio.NewReader(buf)
			text, err := scanner.ReadString('\n')
			if text != "1.2.3.4.5.\n" {
				g.Fail(fmt.Sprintf("Captured text doesn't match expected value: %s", text))
			}
			if err != nil {
				g.Fail(fmt.Sprintf("Error occurred while reading bytestring: %s", err))
			}
			scanner = bufio.NewReader(buf2)
			text, err = scanner.ReadString('\n')
			if text != "1.2.\n" {
				g.Fail(fmt.Sprintf("Captured stderr text doesn't match expected value: %s", text))
			}
			if err != nil {
				g.Fail(fmt.Sprintf("Error occurred while reading stderr bytestring: %s", err))
			}
			//r := bufio.NewReader(str)
			//line, _, err := r.ReadLine()
			//if err != nil {
			//	t.Error("Failed to read line")
			//}
			//fmt.Printf("|%s", line)
		})
	})
}
func TestMain(t *testing.T) {
	g := Goblin(t)
	g.Describe("Main Purpose", func() {
		g.It("Should begin build command on cgi request", func() {
			req, err := http.NewRequest("GET", "http://example.com/jobs/TestProj?start", nil)
			if err != nil { g.Fail("Unable to setup request object") }
			w := httptest.NewRecorder()
			APIJobStart(w, req)
			if w.Code != 200 {
				g.Fail("Job Runner didn't return HTTP 200")
			}
			var msg StatusResponse
			err = json.Unmarshal(w.Body.Bytes(), &msg)
			if err != nil {
				g.Fail(fmt.Sprintf("Failed to parse JSON response: %s", err))
			}
			g.Assert(msg.Started).Equal(true)
			//fmt.Printf("%d - %s", w.Code, w.Body.String())
		})
		g.It("Should use token system")
		g.It("Should add to repository after completed build")
		g.It("Should work with git hooks")
		g.It("Should email on completed build")
		g.It("Should create repository after done building")
	})
	g.Describe("Web", func() {
		g.It("Build Status Monitor should not error on request")
	})
}

/*
DISABLED: until I can test socket.io
type server struct {
	_server *sio.Server
	pattern string
}

func createServer(pattern string) *server {
	return &server{
		_server: socketServer(),
		pattern: pattern,
	}
}

func (self *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Index(r.URL.Path, self.pattern) == 0 {
		self._server.ServeHTTP(w, r)
	}
}

func TestConnection(t *testing.T) {
	go tester()
	//ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	fmt.Fprintln(w, `{"fake twitter json string"}`)
	//}))
	ts := httptest.NewServer(createServer("/socket.io/"))
	pUrl, _ := url.Parse(ts.URL)
	fmt.Println(pUrl.Host)
	defer ts.Close()

	//twitterUrl = ts.URL
	//c := make(chan *twitterResult)
	//go retrieveTweets(c)

	//tweet := <-c
	//if tweet != expected1 {
	//	t.Fail()
	//}
	//tweet = <-c
	//if tweet != expected2 {
	//	t.Fail()
	//}
	time.Sleep(30)
	//t.Skip()
	_, err := socketio.Dial(pUrl.Host)
	if err != nil {
		t.Error(fmt.Sprintf("Failed to connect to server: %s", err))
	}
}
*/
