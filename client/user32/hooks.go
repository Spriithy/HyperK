package user32

import (
	"syscall"

	"github.com/Spriithy/gkl/client/types"
)

var (
	setWindowsHookExW = user32DLL.NewProc("SetWindowsHookExW")
	callNextHookEx    = user32DLL.NewProc("CallNextHookEx")
)

func SetWindowsHookExW(idHook int, lpfn types.HOOKPROC, hMod types.HINSTANCE, dwThreadId types.DWORD) types.HHOOK {
	ret, _, _ := setWindowsHookExW.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return types.HHOOK(ret)
}

func CallNextHookEx(hhk types.HHOOK, nCode int, wParam types.WPARAM, lParam types.LPARAM) types.LRESULT {
	ret, _, _ := callNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return types.LRESULT(ret)
}
