package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	win "github.com/lxn/win"
)

var defaultCursor win.HCURSOR

func main() {
	hInstance := win.GetModuleHandle((*uint16)(nil))
	log.Printf("hInstance: %#v", hInstance)

	wndClassName := win.SysAllocString("MyMainWindowClass")
	wndClass := win.WNDCLASSEX{
		LpszClassName: wndClassName,
		HInstance:     hInstance,
	}

	wndClass.CbSize = uint32(unsafe.Sizeof(wndClass))
	wndClass.LpfnWndProc = syscall.NewCallback(myMainWindowProc)

	defaultCursor = win.LoadCursor(win.HINSTANCE(0), (*uint16)(unsafe.Pointer(uintptr(win.IDC_ARROW))))
	wndClass.HCursor = defaultCursor
	log.Printf("Load default cursor: %v", defaultCursor)

	atom := win.RegisterClassEx(&wndClass)
	if atom == win.ATOM(0) {
		panic(errors.New("Failed to register class. " + getWin32LastError()))
	}

	hWin := win.CreateWindowEx(
		0,
		wndClassName,
		win.SysAllocString("Main Window"),
		win.WS_OVERLAPPEDWINDOW,
		win.CW_USEDEFAULT,
		win.CW_USEDEFAULT,
		win.CW_USEDEFAULT,
		win.CW_USEDEFAULT,
		0,
		0,
		hInstance,
		nil,
	)

	if hWin == win.HWND(0) {
		panic(errors.New("Failed to create a window! " + getWin32LastError()))
	}

	win.ShowWindow(hWin, win.SW_SHOWDEFAULT)

	win.UpdateWindow(hWin)

	win.SetCursor(defaultCursor)

	var msg win.MSG

	for win.GetMessage(&msg, win.HWND(0), 0, 0) > 0 {
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}

	os.Exit(int(msg.WParam))
}

func getWin32LastError() string {
	err := win.GetLastError()
	if err == 0 {
		return ""
	}

	return fmt.Sprintf("#%d", err)
}

func myMainWindowProc(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_PAINT:
		hDc := win.GetDC(hWnd)
		// win.MoveToEx(hDc, 10, 10, nil)
		str := "Hello world!"
		cStr := win.SysAllocString(str)
		win.TextOut(hDc, 0, 0, cStr, int32(win.SysStringLen(cStr)))
		win.SysFreeString(cStr)
		win.ReleaseDC(hWnd, hDc)
	case win.WM_CLOSE:
		win.DestroyWindow(hWnd)
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
	case win.WM_SETCURSOR:
		win.SetCursor(defaultCursor)
	case win.WM_SIZE:
		win.InvalidateRect(hWnd, nil, true)
		win.UpdateWindow(hWnd)
	default:
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}

	return 0
}
