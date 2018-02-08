package qtcore

// /usr/include/qt/QtCore/qsavefile.h
// #include <qsavefile.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
// qint64 writeData(const char *, qint64)
func (this *QSaveFile) InheritWriteData(f func(data string, len int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

type QSaveFile struct {
	*qtrt.CObject
}

func (this *QSaveFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSaveFile) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSaveFileFromPointer(cthis unsafe.Pointer) *QSaveFile {
	return &QSaveFile{&qtrt.CObject{cthis}}
}
func (*QSaveFile) NewFromPointer(cthis unsafe.Pointer) *QSaveFile {
	return NewQSaveFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsavefile.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSaveFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSaveFile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsavefile.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(const QString &)
func NewQSaveFile(name string) *QSaveFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(QObject *)
func NewQSaveFile_1(parent *QObject /*777 QObject **/) *QSaveFile {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:71
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(const QString &, QObject *)
func NewQSaveFile_2(name string, parent *QObject /*777 QObject **/) *QSaveFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSaveFile()
func DeleteQSaveFile(this *QSaveFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsavefile.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QSaveFile) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSaveFile8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsavefile.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QSaveFile) SetFileName(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QSaveFile) Open(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool commit()
func (this *QSaveFile) Commit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile6commitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancelWriting()
func (this *QSaveFile) CancelWriting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile13cancelWritingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirectWriteFallback(_Bool)
func (this *QSaveFile) SetDirectWriteFallback(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile22setDirectWriteFallbackEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool directWriteFallback()
func (this *QSaveFile) DirectWriteFallback() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSaveFile19directWriteFallbackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:87
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QSaveFile) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

//  body block end
