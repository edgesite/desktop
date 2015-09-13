// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSEvent_Class

const (
	NSAnyEventMask = -1
)

const (
	NSLeftMouseDown         = 1
	NSLeftMouseUp           = 2
	NSRightMouseDown        = 3
	NSRightMouseUp          = 4
	NSMouseMoved            = 5
	NSLeftMouseDragged      = 6
	NSRightMouseDragged     = 7
	NSMouseEntered          = 8
	NSMouseExited           = 9
	NSKeyDown               = 10
	NSKeyUp                 = 11
	NSFlagsChanged          = 12
	NSAppKitDefined         = 13
	NSSystemDefined         = 14
	NSApplicationDefined    = 15
	NSPeriodic              = 16
	NSCursorUpdate          = 17
	NSScrollWheel           = 22
	NSTabletPoint           = 23
	NSTabletProximity       = 24
	NSOtherMouseDown        = 25
	NSOtherMouseUp          = 26
	NSOtherMouseDragged     = 27
	NSEventTypeGesture      = 29
	NSEventTypeMagnify      = 30
	NSEventTypeSwipe        = 31
	NSEventTypeRotate       = 18
	NSEventTypeBeginGesture = 19
	NSEventTypeEndGesture   = 20
	NSEventTypeSmartMagnify = 32
	NSEventTypeQuickLook    = 33
	NSEventTypePressure     = 34
)

var NSEventClass unsafe.Pointer = Runtime_objc_lookUpClass("NSEvent")
var NSEventType unsafe.Pointer = Runtime_sel_getUid("type")

type NSEvent struct {
	NSObject
}

func NSEventPointer(p unsafe.Pointer) NSEvent {
	return NSEvent{NSObjectPointer(p)}
}

func (m NSEvent) Type() int {
	return Pointer2Int(Runtime_objc_msgSend(m.Pointer, NSEventType))
}
