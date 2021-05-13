package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/types"
)

var (
	getKeyboardLayout = user32DLL.NewProc("GetKeyboardLayout")
	getAsyncKeyState  = user32DLL.NewProc("GetAsyncKeyState")
	getKeyboardState  = user32DLL.NewProc("GetKeyboardState")
	mapVirtualKeyExW  = user32DLL.NewProc("MapVirtualKeyW")
	toUnicodeEx       = user32DLL.NewProc("ToUnicodeEx")
)

func GetKeyboardLayout(hWnd types.HWND) types.HKL {
	ret, _, _ := getKeyboardLayout.Call(uintptr(hWnd))
	return types.HKL(ret)
}

func GetAsyncKeyState(vkCode int) uint16 {
	ret, _, _ := getAsyncKeyState.Call(
		uintptr(vkCode),
	)
	return uint16(ret)
}

func GetKeyboardState(lbKeyState types.PBYTE) types.BOOL {
	ret, _, _ := getKeyboardState.Call(
		uintptr(unsafe.Pointer(lbKeyState)),
	)
	return types.BOOL(ret)
}

func MapVirtualKeyExW(uCode, uMapType types.UINT, dwhkl types.HKL) types.UINT {
	ret, _, _ := mapVirtualKeyExW.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl),
	)
	return types.UINT(ret)
}

func ToUnicodeEx(vkCode, scanCode types.UINT, lpKeyState types.PBYTE, pwsqzBuff types.LPCWSTR, cchBuff int, wFlags types.UINT, hkl types.HKL) int {
	ret, _, _ := toUnicodeEx.Call(
		uintptr(vkCode),
		uintptr(scanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(pwsqzBuff)),
		uintptr(cchBuff),
		uintptr(wFlags),
		uintptr(hkl),
	)
	return int(ret)
}
