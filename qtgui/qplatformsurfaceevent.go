package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QPlatformSurfaceEvent struct {
	*qtcore.QEvent
}

func (this *QPlatformSurfaceEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QPlatformSurfaceEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQPlatformSurfaceEventFromPointer(cthis unsafe.Pointer) *QPlatformSurfaceEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QPlatformSurfaceEvent{bcthis0}
}
func (*QPlatformSurfaceEvent) NewFromPointer(cthis unsafe.Pointer) *QPlatformSurfaceEvent {
	return NewQPlatformSurfaceEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:451
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlatformSurfaceEvent(enum QPlatformSurfaceEvent::SurfaceEventType)
func NewQPlatformSurfaceEvent(surfaceEventType int) *QPlatformSurfaceEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPlatformSurfaceEventC2ENS_16SurfaceEventTypeE", qtrt.FFI_TYPE_POINTER, surfaceEventType)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlatformSurfaceEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPlatformSurfaceEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:452
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPlatformSurfaceEvent()
func DeleteQPlatformSurfaceEvent(this *QPlatformSurfaceEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPlatformSurfaceEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:454
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QPlatformSurfaceEvent::SurfaceEventType surfaceEventType()
func (this *QPlatformSurfaceEvent) SurfaceEventType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

type QPlatformSurfaceEvent__SurfaceEventType = int

const QPlatformSurfaceEvent__SurfaceCreated QPlatformSurfaceEvent__SurfaceEventType = 0
const QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed QPlatformSurfaceEvent__SurfaceEventType = 1

//  body block end
