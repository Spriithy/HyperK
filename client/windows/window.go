package windows

import (
	"unsafe"
)

type POINT struct {
	X LONG
	Y LONG
}

type MSG struct {
	Hwnd    HWND
	Message UINT
	WParam  WPARAM
	LParam  LPARAM
	Time    DWORD
	Pt      POINT
}

var (
	getForegroundWindow      = user32DLL.NewProc("GetForegroundWindow")
	getWindowTextW           = user32DLL.NewProc("GetWindowTextW")
	getWindowThreadProcessId = user32DLL.NewProc("GetWindowThreadProcessId")
)

func GetForegroundWindow() HWND {
	ret, _, _ := getForegroundWindow.Call()
	return HWND(ret)
}

func GetWindowTextW(hWnd HWND, buf LPCWSTR, maxCount int) int {
	ret, _, _ := getWindowTextW.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(buf)),
		uintptr(maxCount),
	)
	return int(ret)
}

func GetWindowThreadProcessId(hWnd HWND, lpdwProcessId LPDWORD) HWND {
	ret, _, _ := getWindowThreadProcessId.Call(
		uintptr(hWnd),
		uintptr(lpdwProcessId),
	)
	return HWND(ret)
}
