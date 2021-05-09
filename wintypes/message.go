package wintypes

type POINT struct {
	X LONG
	Y LONG
}

type PPOINT = POINT

type MSG struct {
	Hwnd    HWND
	Message UINT
	WParam  WPARAM
	LParam  LPARAM
	Time    DWORD
	Pt      POINT
}

type PMSG = *MSG
type NPMSG = *MSG
type LPMSG = *MSG
