// build +windows

package desktop

import (
	"unicode/utf16"
	"syscall"
	"unsafe"
	"reflect"
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
	
	return int(uint8(v)), int(uint8(v>>8)), int(uint16(v>>16))
}

func IsWindowsXP() bool {
	v1, v2, _ := GetVersion()

	return v1 == 5 && v2 == 1
}

func String2WString(s string) uintptr {
  return uintptr(unsafe.Pointer(&utf16.Encode([]rune(s + "\x00"))[0]))
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

func Ptr(d interface{}) uintptr {
	v := reflect.ValueOf(d)
	UIntPtr := reflect.TypeOf((uintptr)(0))
	
	if v.Type().ConvertibleTo(UIntPtr) {
		vv := v.Convert(UIntPtr)
		return vv.Interface().(uintptr)
	}else {
		return v.Pointer()
	}
}
