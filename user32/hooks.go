package user32

import (
	"syscall"

	"github.com/Spriithy/gkl/types"
)

var (
	setWindowsHookExA = user32DLL.NewProc("SetWindowsHookExA")
	setWindowsHookExW = user32DLL.NewProc("SetWindowsHookExW")
	callNextHookEx    = user32DLL.NewProc("CallNextHookEx")
)

func SetWindowsHookExA(idHook int, lpfn types.HOOKPROC, hMod types.HINSTANCE, dwThreadId types.DWORD) types.HHOOK {
	ret, _, _ := setWindowsHookExA.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return types.HHOOK(ret)
}

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
