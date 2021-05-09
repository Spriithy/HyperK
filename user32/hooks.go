package user32

import (
	"syscall"

	"github.com/Spriithy/gkl/wintypes"
)

var (
	setWindowsHookExA = user32DLL.NewProc("SetWindowsHookExA")
	callNextHookEx    = user32DLL.NewProc("CallNextHookEx")
)

func SetWindowsHookExA(idHook int, lpfn wintypes.HOOKPROC, hMod wintypes.HINSTANCE, dwThreadId wintypes.DWORD) wintypes.HHOOK {
	ret, _, _ := setWindowsHookExA.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return wintypes.HHOOK(ret)
}

func CallNextHookEx(hhk wintypes.HHOOK, nCode int, wParam wintypes.WPARAM, lParam wintypes.LPARAM) wintypes.LRESULT {
	ret, _, _ := callNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return wintypes.LRESULT(ret)
}
