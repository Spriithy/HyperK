package windows

import (
	"unsafe"
)

var (
	getMessageW = user32DLL.NewProc("GetMessageW")
)

func GetMessageW(lpMsg *MSG, hWnd HWND, wMsgFilterMin, wMsgFilterMax UINT) BOOL {
	ret, _, _ := getMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return BOOL(ret)
}
