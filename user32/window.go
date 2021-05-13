package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/wintypes"
)

var (
	getForegroundWindow      = user32DLL.NewProc("GetForegroundWindow")
	getKeyboardLayout        = user32DLL.NewProc("GetKeyboardLayout")
	getKeyboardState         = user32DLL.NewProc("GetKeyboardState")
	mapVirtualKeyExW         = user32DLL.NewProc("MapVirtualKeyW")
	getWindowTextA           = user32DLL.NewProc("GetWindowTextA")
	getWindowThreadProcessId = user32DLL.NewProc("GetWindowThreadProcessId")
)

func GetForegroundWindow() wintypes.HWND {
	ret, _, _ := getForegroundWindow.Call()
	return wintypes.HWND(ret)
}

func GetKeyboardLayout(hWnd wintypes.HWND) wintypes.HKL {
	ret, _, _ := getKeyboardLayout.Call(uintptr(hWnd))
	return wintypes.HKL(ret)
}

func GetKeyboardState(lbKeyState wintypes.PBYTE) wintypes.BOOL {
	ret, _, _ := getKeyboardState.Call(
		uintptr(unsafe.Pointer(lbKeyState)),
	)
	return wintypes.BOOL(ret)
}

func MapVirtualKeyExW(uCode, uMapType wintypes.UINT, dwhkl wintypes.HKL) wintypes.UINT {
	ret, _, _ := mapVirtualKeyExW.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl),
	)
	return wintypes.UINT(ret)
}

func GetWindowTextA(hWnd wintypes.HWND, buf wintypes.LPCSTR, maxCount int) int {
	ret, _, _ := getWindowTextA.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(buf)),
		uintptr(maxCount),
	)
	return int(ret)
}

func GetWindowThreadProcessId(hWnd wintypes.HWND, lpdwProcessId wintypes.LPDWORD) wintypes.HWND {
	ret, _, _ := getWindowThreadProcessId.Call(
		uintptr(hWnd),
		uintptr(lpdwProcessId),
	)
	return wintypes.HWND(ret)
}
