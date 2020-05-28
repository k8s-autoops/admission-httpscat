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
	optPort         string
	optResponse     []byte
	optResponseType string
	optResponseCode int
)

func init() {
	optPort = os.Getenv("PORT")
	if optPort == "" {
		optPort = "80"
	}
	optResponse = []byte(os.Getenv("RESPONSE"))
	if len(optResponse) == 0 {
		optResponse = []byte("OK")
	}
	optResponseType = os.Getenv("RESPONSE_TYPE")
	if optResponseType == "" {
		optResponseType = "text/plain; charset=utf-8"
	}
	optResponseCode, _ = strconv.Atoi(os.Getenv("RESPONSE_CODE"))
	if optResponseCode == 0 {
		optResponseCode = http.StatusOK
	}
}

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
		rw.Header().Set("Content-Type", optResponseType)
		rw.Header().Set("Content-Length", strconv.Itoa(len(optResponse)))
		rw.WriteHeader(optResponseCode)
		_, _ = rw.Write(optResponse)
	})

	log.Printf("listening at %s", optPort)
	log.Fatal(http.ListenAndServe(":"+optPort, nil))
}
