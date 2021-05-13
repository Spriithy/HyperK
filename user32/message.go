package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/wintypes"
)

var (
	getMessageW      = user32DLL.NewProc("GetMessageW")
	translateMessage = user32DLL.NewProc("TranslateMessage")
	dispatchMessage  = user32DLL.NewProc("DispatchMessage")
)

func GetMessageW(lpMsg *wintypes.MSG, hWnd wintypes.HWND, wMsgFilterMin, wMsgFilterMax wintypes.UINT) wintypes.BOOL {
	ret, _, _ := getMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return wintypes.BOOL(ret)
}

func TranslateMessage(lpMsg *wintypes.MSG) wintypes.BOOL {
	ret, _, _ := translateMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return wintypes.BOOL(ret)
}

func DispatchMessage(lpMsg *wintypes.MSG) wintypes.LRESULT {
	ret, _, _ := dispatchMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return wintypes.LRESULT(ret)
}
