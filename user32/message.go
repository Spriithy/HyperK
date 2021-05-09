package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/wintypes"
)

var (
	getMessageA = user32DLL.NewProc("GetMessageA")
)

func GetMessageA(lpMsg *wintypes.MSG, hWnd wintypes.HWND, wMsgFilterMin, wMsgFilterMax wintypes.UINT) wintypes.BOOL {
	ret, _, _ := getMessageA.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return wintypes.BOOL(ret)
}
