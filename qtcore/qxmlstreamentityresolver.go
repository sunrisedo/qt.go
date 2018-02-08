package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QXmlStreamEntityResolver struct {
	*qtrt.CObject
}

func (this *QXmlStreamEntityResolver) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamEntityResolver) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamEntityResolverFromPointer(cthis unsafe.Pointer) *QXmlStreamEntityResolver {
	return &QXmlStreamEntityResolver{&qtrt.CObject{cthis}}
}
func (*QXmlStreamEntityResolver) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamEntityResolver {
	return NewQXmlStreamEntityResolverFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:336
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QXmlStreamEntityResolver()
func DeleteQXmlStreamEntityResolver(this *QXmlStreamEntityResolver) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolverD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:337
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString resolveEntity(const QString &, const QString &)
func (this *QXmlStreamEntityResolver) ResolveEntity(publicId string, systemId string) string {
	var tmpArg0 = NewQString_5(publicId)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(systemId)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qxmlstream.h:338
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString resolveUndeclaredEntity(const QString &)
func (this *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name string) string {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

//  body block end
