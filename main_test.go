package main

import (
	"testing"
//	"time"
	"fmt"
//	"os"
	"os/exec"
	"bytes"
	"bufio"
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
	cmd := exec.Command("./proc_test.sh")
	//stdout, err := cmd.StdoutPipe()
	Chan := make(chan []byte)
	Chan2 := make(chan []byte)
	var buf *bytes.Buffer = new(bytes.Buffer)
	var buf2 *bytes.Buffer = new(bytes.Buffer)
	go dotWriter(Chan, bufio.NewWriter(buf))
	go dotWriter(Chan2, bufio.NewWriter(buf2))
	sout := Stream{Chan:Chan}
	serr := Stream{Chan:Chan2}
	cmd.Stdout = sout
	cmd.Stderr = serr
	cmd.Start()
	cmd.Wait()
	scanner := bufio.NewReader(buf)
text, err:=scanner.ReadString('\n')
		if text != "1.2.3.4.5.\n" {
			t.Error(fmt.Sprintf("Captured text doesn't match expected value: %s", text))
		}
	if err != nil {
		t.Error(fmt.Sprintf("Error occurred while reading bytestring: %s", err))
	}
	scanner = bufio.NewReader(buf2)
text, err=scanner.ReadString('\n')
		if text != "1.2.\n" {
			t.Error(fmt.Sprintf("Captured stderr text doesn't match expected value: %s", text))
		}
	if err != nil {
		t.Error(fmt.Sprintf("Error occurred while reading stderr bytestring: %s", err))
	}
	//r := bufio.NewReader(str)
	//line, _, err := r.ReadLine()
	//if err != nil {
	//	t.Error("Failed to read line")
	//}
	//fmt.Printf("|%s", line)
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
