package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/types"
)

var (
	getForegroundWindow      = user32DLL.NewProc("GetForegroundWindow")
	getWindowTextA           = user32DLL.NewProc("GetWindowTextA")
	getWindowThreadProcessId = user32DLL.NewProc("GetWindowThreadProcessId")
)

func GetForegroundWindow() types.HWND {
	ret, _, _ := getForegroundWindow.Call()
	return types.HWND(ret)
}

func GetWindowTextA(hWnd types.HWND, buf types.LPCSTR, maxCount int) int {
	ret, _, _ := getWindowTextA.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(buf)),
		uintptr(maxCount),
	)
	return int(ret)
}

func GetWindowThreadProcessId(hWnd types.HWND, lpdwProcessId types.LPDWORD) types.HWND {
	ret, _, _ := getWindowThreadProcessId.Call(
		uintptr(hWnd),
		uintptr(lpdwProcessId),
	)
	return types.HWND(ret)
}
