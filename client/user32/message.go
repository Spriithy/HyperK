package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/client/types"
)

var (
	getMessageW = user32DLL.NewProc("GetMessageW")
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
