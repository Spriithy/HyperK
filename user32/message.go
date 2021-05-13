package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/types"
)

var (
	getMessageW      = user32DLL.NewProc("GetMessageW")
	translateMessage = user32DLL.NewProc("TranslateMessage")
	dispatchMessage  = user32DLL.NewProc("DispatchMessage")
)

func GetMessageW(lpMsg *types.MSG, hWnd types.HWND, wMsgFilterMin, wMsgFilterMax types.UINT) types.BOOL {
	ret, _, _ := getMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return types.BOOL(ret)
}

func TranslateMessage(lpMsg *types.MSG) types.BOOL {
	ret, _, _ := translateMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return types.BOOL(ret)
}

func DispatchMessage(lpMsg *types.MSG) types.LRESULT {
	ret, _, _ := dispatchMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return types.LRESULT(ret)
}
