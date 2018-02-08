package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

type QAssociativeIterable struct {
	*qtrt.CObject
}

func (this *QAssociativeIterable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAssociativeIterable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAssociativeIterableFromPointer(cthis unsafe.Pointer) *QAssociativeIterable {
	return &QAssociativeIterable{&qtrt.CObject{cthis}}
}
func (*QAssociativeIterable) NewFromPointer(cthis unsafe.Pointer) *QAssociativeIterable {
	return NewQAssociativeIterableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:680
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator begin()
func (this *QAssociativeIterable) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:681
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator end()
func (this *QAssociativeIterable) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:682
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator find(const QVariant &)
func (this *QAssociativeIterable) Find(key *QVariant) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable4findERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:684
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QVariant &)
func (this *QAssociativeIterable) Value(key *QVariant) *QVariant /*123*/ {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable5valueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:686
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QAssociativeIterable) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQAssociativeIterable(this *QAssociativeIterable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAssociativeIterableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
