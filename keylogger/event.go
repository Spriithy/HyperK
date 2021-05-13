package keylogger

import (
	"unsafe"

	"github.com/Spriithy/gkl/types"
)

type KeyboardEvent struct {
	WParam     types.WPARAM
	LParam     types.LPARAM
	Layout     types.HKL
	HookStruct *types.KBDLLHOOKSTRUCT
}

func NewKeyboardEvent(wParam types.WPARAM, lParam types.LPARAM, layout types.HKL) *KeyboardEvent {
	return &KeyboardEvent{
		WParam:     wParam,
		LParam:     lParam,
		Layout:     layout,
		HookStruct: (*types.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
	}
}

func (ke *KeyboardEvent) IsUp() bool {
	return ke.IsKeyUp() || ke.IsSysKeyUp()
}

func (ke *KeyboardEvent) IsDown() bool {
	return ke.IsKeyDown() || ke.IsSysKeyDown()
}

func (ke *KeyboardEvent) IsKeyUp() bool {
	return ke.WParam == types.WM_KEYUP
}

func (ke *KeyboardEvent) IsKeyDown() bool {
	return ke.WParam == types.WM_KEYDOWN
}

func (ke *KeyboardEvent) IsSysKeyUp() bool {
	return ke.WParam == types.WM_SYSKEYUP
}

func (ke *KeyboardEvent) IsSysKeyDown() bool {
	return ke.WParam == types.WM_SYSKEYDOWN
}

func (ke *KeyboardEvent) IsMode(mode types.WPARAM) bool {
	return ke.WParam == mode
}

func (ke *KeyboardEvent) IsVk(vkCode types.DWORD) bool {
	return ke.HookStruct.VkCode == vkCode
}

func (ke *KeyboardEvent) IsVkMode(vkCode types.DWORD, mode types.WPARAM) bool {
	return ke.IsMode(mode) && ke.IsVk(vkCode)
}

func (ke *KeyboardEvent) IsShift() bool {
	return ke.IsVk(types.VK_SHIFT) || ke.IsVk(types.VK_LSHIFT) || ke.IsVk(types.VK_RSHIFT)
}

func (ke *KeyboardEvent) IsCaps() bool {
	return ke.IsVk(types.VK_CAPITAL)
}

func (ke *KeyboardEvent) IsReturn() bool {
	return ke.IsVk(types.VK_RETURN)
}

func (ke *KeyboardEvent) IsBackspace() bool {
	return ke.IsVk(types.VK_BACK)
}

func (ke *KeyboardEvent) IsControl() bool {
	return ke.IsVk(types.VK_CONTROL)
}

func (ke *KeyboardEvent) IsEscape() bool {
	return ke.IsVk(types.VK_ESCAPE)
}
