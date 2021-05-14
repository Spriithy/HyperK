package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/client/types"
)

var (
	getForegroundWindow      = user32DLL.NewProc("GetForegroundWindow")
	getWindowTextW           = user32DLL.NewProc("GetWindowTextW")
	getWindowThreadProcessId = user32DLL.NewProc("GetWindowThreadProcessId")
)

func GetForegroundWindow() types.HWND {
	ret, _, _ := getForegroundWindow.Call()
	return types.HWND(ret)
}

func GetWindowTextW(hWnd types.HWND, buf types.LPCWSTR, maxCount int) int {
	ret, _, _ := getWindowTextW.Call(
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
