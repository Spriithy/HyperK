package user32

import (
	"unsafe"

	"github.com/Spriithy/gkl/types"
)

var (
	getKeyboardLayout = user32DLL.NewProc("GetKeyboardLayout")
	getKeyboardState  = user32DLL.NewProc("GetKeyboardState")
	toUnicodeEx       = user32DLL.NewProc("ToUnicodeEx")
)

func GetKeyboardLayout(hWnd types.HWND) types.HKL {
	ret, _, _ := getKeyboardLayout.Call(uintptr(hWnd))
	return types.HKL(ret)
}

func GetKeyboardState(lbKeyState types.PBYTE) types.BOOL {
	ret, _, _ := getKeyboardState.Call(
		uintptr(unsafe.Pointer(lbKeyState)),
	)
	return types.BOOL(ret)
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
