package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	responseBody = []byte("OK")
	responseType = "text/plain; charset=utf-8"
)

const (
	certFile = "/autoops-data/tls/tls.crt"
	keyFile  = "/autoops-data/tls/tls.key"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	count := uint64(0)
	locker := &sync.Mutex{}

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		locker.Lock()
		defer locker.Unlock()
		count++

		// request id
		title := fmt.Sprintf("================== %s ==== #%d ==================", time.Now().Format(time.RFC3339), count)
		log.Printf(title)
		// proto / method / url
		log.Println("")
		log.Printf("%s %s %s", req.Proto, req.Method, req.URL.String())
		// headers
		log.Println("")
		// fix for golang Host header
		log.Printf("Host: %s", req.Host)
		for k, vs := range req.Header {
			for _, v := range vs {
				log.Printf("%s: %s", k, v)
			}
		}
		// body
		log.Println("")
		br := bufio.NewReader(req.Body)
		for {
			line, err := br.ReadString('\n')
			log.Println(line)
			if err != nil {
				break
			}
		}
		endLine := &strings.Builder{}
		for range title {
			endLine.WriteRune('=')
		}
		log.Println(endLine.String())

		// response with OK
		rw.Header().Set("Content-Type", responseType)
		rw.Header().Set("Content-Length", strconv.Itoa(len(responseBody)))
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(responseBody)
	})

	log.Println("listening at :443")
	log.Fatal(http.ListenAndServeTLS(":443", certFile, keyFile, nil))
}
