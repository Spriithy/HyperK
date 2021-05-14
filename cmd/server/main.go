package kbdserv

import (
	"net/http"
)

func main() {
	err := http.ListenAndServe("0.0.0.0:8443", nil)
	if err != nil {
		panic(err)
	}
}
