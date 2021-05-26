package ebest

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	kernel32, _      = syscall.LoadLibrary("kernel32.dll")
	user32, _        = syscall.LoadLibrary("User32.dll")
	pCreateThread, _ = syscall.GetProcAddress(kernel32, "CreateThread")
	pPeekMessage, _  = syscall.GetProcAddress(user32, "PeekMessageW")
)

type threadFunc func(uintptr) uintptr

// createThread : go-scheduler를 통하지 않고 직접 쓰레드를 생성하는 API
//
// COM에서는 이벤트를 맺은 윈도우 쓰레드에 메시지를 전달,
//
// go-scheduler은 M개의 OS-threads와 N개의 go-routines 이 존재하기 때문에 이벤스 수신이 중간에 끊김
func createThread(f threadFunc, arg1 uintptr) (int32, error) {
	ret, _, err := syscall.Syscall6(uintptr(pCreateThread), 6,
		0, 0,
		syscall.NewCallback(f),
		uintptr(arg1), 0, 0)

	return int32(ret), err
}

func peekMessage(msg *ole.Msg, hwnd uint32, MsgFilterMin uint32, MsgFilterMax uint32, RemoveMsg uint32) (ret int32, err error) {
	r0, _, err := syscall.Syscall6(uintptr(pPeekMessage), 5,
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(MsgFilterMin),
		uintptr(MsgFilterMax),
		uintptr(RemoveMsg),
		0)

	ret = int32(r0)
	return
}

func pumpWaitingMessages() int32 {
	ret := int32(0)

	var msg ole.Msg

	runtime.LockOSThread()
	for {
		r, _ := peekMessage(&msg, 0, 0, 0, 1)
		if r == 0 {
			break
		}
		if msg.Message == 0x0012 { // WM_QUIT
			ret = int32(1)
			break
		}
		ole.DispatchMessage(&msg)
	}
	runtime.UnlockOSThread()
	return ret
}
