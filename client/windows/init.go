package windows

import "syscall"

var (
	user32DLL = syscall.NewLazyDLL("user32.dll")
)
