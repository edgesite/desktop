package cocoa

import (
  "unsafe"

  "../objc"
)

// http://developer.apple.com/library/mac/#documentation/Cocoa/Reference/Foundation/Classes/NSData_Class/Reference/Reference.html#//apple_ref/doc/c_ref/NSData

var NSArrayClass unsafe.Pointer = objc.Runtime_objc_lookUpClass("NSArray")
var NSArrayCount unsafe.Pointer = objc.Runtime_sel_getUid("count")
var NSArrayObjectAtIndex unsafe.Pointer = objc.Runtime_sel_getUid("objectAtIndex:")

type NSArray struct {
  NSObject
}

func NSArrayPointer(p unsafe.Pointer) NSArray {
  var m NSArray = NSArray{NSObjectPointer(p)}
  return m
}

func (m NSArray) Count() int {
  u := (uintptr)(objc.Runtime_objc_msgSend(m.Pointer, NSArrayCount))
  return (int)(u)
}

func (m NSArray) ObjectAtIndex(i int) unsafe.Pointer {
  return objc.Runtime_objc_msgSend(m.Pointer, NSArrayObjectAtIndex, unsafe.Pointer(uintptr(i)))
}
