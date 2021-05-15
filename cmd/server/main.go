package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(base64.NewDecoder(base64.RawStdEncoding, r.Body))
		fmt.Println(string(data))
	})

	http.ListenAndServe("localhost:8000", nil)
}
