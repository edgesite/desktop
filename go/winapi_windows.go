// build +windows

package desktop

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

/*
https://msdn.microsoft.com/en-us/library/ms724832(VS.85).aspx

6.1               Windows 7     / Windows 2008 R2
6.0               Windows Vista / Windows 2008
5.2               Windows 2003
5.1               Windows XP
5.0               Windows 2000
*/
func GetVersion() (int, int, int) {
	v, err := syscall.GetVersion()
	if err != nil {
		panic(err)
	}

	return int(uint8(v)), int(uint8(v >> 8)), int(uint16(v >> 16))
}

func IsWindowsXP() bool {
	v1, v2, _ := GetVersion()

	return v1 == 5 && v2 == 1
}

type WString uintptr

func WStringPtr(r1, r2 uintptr, err error) WString {
	LastError = uintptr(err.(syscall.Errno))
	return WString(r1)
}

func WStringNew(s string) WString {
	u := utf16.Encode([]rune(s + "\x00"))

	l := len(u) * int(unsafe.Sizeof(u[0]))

	m := WStringPtr(GlobalAlloc.Call(Arg(GMEM_FIXED), Arg(l)))
	if m == 0 {
		panic(GetLastErrorString())
	}

	to := (*(*[1 << 30]byte)(unsafe.Pointer(m)))[:(l)]
	from := (*(*[1 << 30]byte)(unsafe.Pointer(&u[0])))[:(l)]

	for i := range to {
		to[i] = from[i]
	}

	return m
}

func (m WString) Size() int {
	return int(UINTPtr(lstrlen.Call(Arg(m))))
}

func (m WString) Close() {
	h := HRESULTPtr(GlobalFree.Call(Arg(m)))
	if h != 0 {
		panic(h.String())
	}
}

func WString2String(p uintptr) string {
	var rr []uint16 = make([]uint16, 0, MAX_PATH)
	for p := uintptr(unsafe.Pointer(p)); ; p += 2 {
		u := *(*uint16)(unsafe.Pointer(p))
		if u == 0 {
			return string(utf16.Decode(rr))
		}
		rr = append(rr, u)
	}
	panic("No zero at end of the string")
}

func WArray2String(rr []uint16) string {
	return string(utf16.Decode(rr))
}

var Bool2Int = map[bool]int{
	true:  1,
	false: 0,
}

func Arg(d interface{}) uintptr {
	switch d.(type) {
	case bool:
		return uintptr(Bool2Int[d.(bool)])
	}

	v := reflect.ValueOf(d)
	UIntPtr := reflect.TypeOf((uintptr)(0))

	if v.Type().ConvertibleTo(UIntPtr) {
		vv := v.Convert(UIntPtr)
		return vv.Interface().(uintptr)
	} else {
		return v.Pointer()
	}
}

var NULL = Arg(0)

// copy last error from last syscall
var LastError uintptr

type HMENU uintptr

func HMENUPtr(r1, r2 uintptr, err error) HMENU {
	LastError = uintptr(err.(syscall.Errno))
	return HMENU(r1)
}

func (m HMENU) Close() {
	if !BOOLPtr(DestroyMenu.Call(Arg(m))).Bool() {
		panic(GetLastErrorString())
	}
}

type HRESULT uintptr

func HRESULTPtr(r1, r2 uintptr, err error) HRESULT {
	LastError = uintptr(err.(syscall.Errno))
	return HRESULT(r1)
}

func (m HRESULT) String() string {
	msg := [1024]uint16{}
	FormatMessage.Call(Arg(FORMAT_MESSAGE_FROM_SYSTEM), NULL, Arg(m), NULL, Arg(&msg[0]), Arg(len(msg)), NULL)
	return fmt.Sprintf("HRESULT: 0x%08x [%s]", uintptr(m), strings.TrimSpace(WString2String(Arg(&msg[0]))))
}

var Int2Bool = map[int]bool{
	1: true,
	0: false,
}

type BOOL uint32

func BOOLPtr(r1, r2 uintptr, err error) BOOL {
	LastError = uintptr(err.(syscall.Errno))
	return BOOL(r1)
}

func (m BOOL) Bool() bool {
	return Int2Bool[int(m)]
}

type DWORD uint32
type UINT uint32
type ULONG uint32
type ULONG_PTR uintptr

func UINTPtr(r1, r2 uintptr, err error) UINT {
	LastError = uintptr(err.(syscall.Errno))
	return UINT(r1)
}

type HFONT uintptr

func HFONTPtr(r1, r2 uintptr, err error) HFONT {
	LastError = uintptr(err.(syscall.Errno))
	return HFONT(r1)
}

func (m HFONT) Close() {
	DeleteObject.Call(Arg(m))
}

type HWND uintptr

func HWNDPtr(r1, r2 uintptr, err error) HWND {
	LastError = uintptr(err.(syscall.Errno))
	return HWND(r1)
}

func (m HWND) Close() {
	if !BOOLPtr(DestroyWindow.Call(Arg(m))).Bool() {
		panic(GetLastErrorString())
	}
}

type TCHAR uint16

func GetLastErrorString() string {
	return HRESULT(LastError).String()
}

type WNDPROC uintptr

func WNDPROCNew(fn interface{}) WNDPROC {
	return WNDPROC(syscall.NewCallback(fn))
}

type HINSTANCE uintptr

func HINSTANCEPtr(r1, r2 uintptr, err error) HINSTANCE {
	LastError = uintptr(err.(syscall.Errno))
	return HINSTANCE(r1)
}

type BYTE byte
type WORD uint16
type LPCTSTR uintptr
type LPTSTR uintptr
type HCURSOR uintptr
type HBRUSH uintptr
type LONG uint32

type HANDLE uintptr

func HANDLEPtr(r1, r2 uintptr, err error) HANDLE {
	LastError = uintptr(err.(syscall.Errno))
	return HANDLE(r1)
}

type COLORREF uintptr

func COLORREFPtr(r1, r2 uintptr, err error) COLORREF {
	LastError = uintptr(err.(syscall.Errno))
	return COLORREF(r1)
}

type LRESULT uintptr

func LRESULTPtr(r1, r2 uintptr, err error) LRESULT {
	LastError = uintptr(err.(syscall.Errno))
	return LRESULT(r1)
}

type WPARAM uintptr
type LPARAM uintptr

type ATOM uintptr

func ATOMPtr(r1, r2 uintptr, err error) ATOM {
	LastError = uintptr(err.(syscall.Errno))
	return ATOM(r1)
}

type HDC uintptr

func HDCPtr(r1, r2 uintptr, err error) HDC {
	LastError = uintptr(err.(syscall.Errno))
	return HDC(r1)
}
