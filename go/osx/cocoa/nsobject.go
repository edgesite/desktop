package cocoa

import (
  "unsafe"

  "../objc"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSObject_Class/index.html#//apple_ref/occ/cl/NSObject

var NSObjectClass unsafe.Pointer = objc.Runtime_objc_lookUpClass("NSObject")
var NSObjectAlloc unsafe.Pointer = objc.Runtime_sel_getUid("alloc")
var NSObjectRetain unsafe.Pointer = objc.Runtime_sel_getUid("retain")
var NSObjectRelease unsafe.Pointer = objc.Runtime_sel_getUid("release")

type NSObject struct {
  unsafe.Pointer
}

func NSObjectNew() NSObject {
  return NSObjectPointer(objc.Runtime_objc_msgSend(NSObjectClass, NSObjectAlloc))
}

func NSObjectPointer(p unsafe.Pointer) NSObject {
  var m NSObject = NSObject{p}
  m.Retain()
  return m
}

func (m NSObject) Retain() {
  objc.Runtime_objc_msgSend(m.Pointer, NSObjectRetain)
}

func (m NSObject) Release() {
  objc.Runtime_objc_msgSend(m.Pointer, NSObjectRelease)
}
