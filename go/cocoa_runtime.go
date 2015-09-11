// +build darwin

package desktop

/*
#cgo LDFLAGS: -lobjc -framework AppKit

#include <stdlib.h>
#include <objc/objc-runtime.h>

id objc_msgSend0(id to, SEL sel) {
  return objc_msgSend(to, sel);
}

id objc_msgSend1(id to, SEL sel, void* arg1) {
  return objc_msgSend(to, sel, arg1);
}

id objc_msgSend2(id to, SEL sel, void* arg1, void* arg2) {
  return objc_msgSend(to, sel, arg1, arg2);
}

id objc_msgSend3(id to, SEL sel, void* arg1, void* arg2, void* arg3) {
  return objc_msgSend(to, sel, arg1, arg2, arg3);
}
*/
import "C"

import (
  "fmt"
  "unsafe"
)

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/ObjCRuntimeRef/Reference/reference.html

func Runtime_objc_lookUpClass(s string) unsafe.Pointer {
  c := unsafe.Pointer(C.CString(s))
  defer C.free(c)
  return unsafe.Pointer(C.objc_lookUpClass((*C.char)(c)))
}

func Runtime_class_getName(m unsafe.Pointer) string {
  return C.GoString(C.class_getName((*C.struct_objc_class)(m)))
}

func Runtime_sel_getUid(s string) unsafe.Pointer {
  var c *C.char = C.CString(s)
  defer C.free(unsafe.Pointer(c))
  return unsafe.Pointer(C.sel_getUid(c))
}

func Runtime_sel_getName(sel unsafe.Pointer) string {
  return C.GoString(C.sel_getName((*C.struct_objc_selector)(sel)))
}

func Runtime_objc_msgSend(self unsafe.Pointer, sel unsafe.Pointer, args... unsafe.Pointer) unsafe.Pointer {
  switch len(args) {
  case 0:
    return unsafe.Pointer(C.objc_msgSend0((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel)))
  case 1:
    return unsafe.Pointer(C.objc_msgSend1((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel), args[0]))
  case 2:
    return unsafe.Pointer(C.objc_msgSend2((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel), args[0], args[1]))
  case 3:
    return unsafe.Pointer(C.objc_msgSend3((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel), args[0], args[1], args[2]))
  default:
    panic(fmt.Sprint("Unsupported number of arguments ", len(args)))
  }
}

//Pointer class_createInstance(Pointer cls, int extraBytes);

//boolean class_addMethod(Pointer cls, Pointer selector, StdCallCallback imp, String types);

//Pointer sel_registerName(String name);

//Pointer objc_allocateClassPair(Pointer superClass, String name, long extraBytes);

//void objc_registerClassPair(Pointer cls);

//Pointer class_getInstanceMethod(Pointer cls, Pointer selecter);

//Pointer method_setImplementation(Pointer method, StdCallCallback imp);

//Pointer objc_getProtocol(String protocol);

//boolean class_addProtocol(Pointer cls, Pointer protocol);
