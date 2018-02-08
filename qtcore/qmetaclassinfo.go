package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QMetaClassInfo struct {
	*qtrt.CObject
}

func (this *QMetaClassInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaClassInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMetaClassInfoFromPointer(cthis unsafe.Pointer) *QMetaClassInfo {
	return &QMetaClassInfo{&qtrt.CObject{cthis}}
}
func (*QMetaClassInfo) NewFromPointer(cthis unsafe.Pointer) *QMetaClassInfo {
	return NewQMetaClassInfoFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:303
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMetaClassInfo()
func NewQMetaClassInfo() *QMetaClassInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMetaClassInfoC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQMetaClassInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMetaClassInfo)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:304
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * name()
func (this *QMetaClassInfo) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMetaClassInfo4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:305
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * value()
func (this *QMetaClassInfo) Value() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMetaClassInfo5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:306
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMetaObject * enclosingMetaObject()
func (this *QMetaClassInfo) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QMetaClassInfo19enclosingMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQMetaClassInfo(this *QMetaClassInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMetaClassInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
