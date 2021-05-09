package wintypes

type (
	BOOL          uint32
	BOOLEAN       byte
	BYTE          byte
	DWORD         uint32
	DWORD64       uint64
	HANDLE        uintptr
	HLOCAL        uintptr
	LARGE_INTEGER int64
	LONG          int32
	LPVOID        uintptr
	SIZE_T        uintptr
	UINT          uint32
	ULONG_PTR     uintptr
	ULONGLONG     uint64
	WORD          uint16
	WPARAM        uintptr
	LPARAM        uintptr
	LRESULT       uintptr
	HINSTANCE     HANDLE
	HHOOK         HANDLE
	HWND          HANDLE
	LPCSTR        *uint8
	LPCWSTR       *uint16
)

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT
