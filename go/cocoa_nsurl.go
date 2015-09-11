// +build darwin

package desktop

import "C"

import (
  "unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSURL_Class/

var NSURLClass unsafe.Pointer = Runtime_objc_lookUpClass("NSURL")
var NSURLAbsoluteString unsafe.Pointer = Runtime_sel_getUid("absoluteString")
var NSURLPath unsafe.Pointer = Runtime_sel_getUid("path")

type NSURL struct {
  NSObject
}

func NSURLPointer(p unsafe.Pointer) NSURL {
  var m NSURL = NSURL{NSObjectPointer(p)}
  return m
}

func (m NSObject) AbsoluteString() string {
  s := NSStringPointer(Runtime_objc_msgSend(m.Pointer, NSURLAbsoluteString))
  defer s.Release()
  return s.String()
}

func (m NSObject) Path() string {
  s := NSStringPointer(Runtime_objc_msgSend(m.Pointer, NSURLPath))
  defer s.Release()
  return s.String()
}
