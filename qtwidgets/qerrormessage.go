package qtwidgets

// /usr/include/qt/QtWidgets/qerrormessage.h
// #include <qerrormessage.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// void done(int)
func (this *QErrorMessage) InheritDone(f func(arg0 int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "done", f)
}

// void changeEvent(class QEvent *)
func (this *QErrorMessage) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

type QErrorMessage struct {
	*QDialog
}

func (this *QErrorMessage) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QErrorMessage) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQErrorMessageFromPointer(cthis unsafe.Pointer) *QErrorMessage {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QErrorMessage{bcthis0}
}
func (*QErrorMessage) NewFromPointer(cthis unsafe.Pointer) *QErrorMessage {
	return NewQErrorMessageFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QErrorMessage) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QErrorMessage10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qerrormessage.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QErrorMessage(QWidget *)
func NewQErrorMessage(parent *QWidget /*777 QWidget **/) *QErrorMessage {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QErrorMessageC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQErrorMessageFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qerrormessage.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QErrorMessage()
func DeleteQErrorMessage(this *QErrorMessage) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QErrorMessageD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [8] QErrorMessage * qtHandler()
func (this *QErrorMessage) QtHandler() *QErrorMessage /*777 QErrorMessage **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QErrorMessage9qtHandlerEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQErrorMessageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QErrorMessage_QtHandler() *QErrorMessage /*777 QErrorMessage **/ {
	var nilthis *QErrorMessage
	rv := nilthis.QtHandler()
	return rv
}

// /usr/include/qt/QtWidgets/qerrormessage.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &)
func (this *QErrorMessage) ShowMessage(message string) {
	var tmpArg0 = qtcore.NewQString_5(message)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QErrorMessage11showMessageERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &)
func (this *QErrorMessage) ShowMessage_1(message string, type_ string) {
	var tmpArg0 = qtcore.NewQString_5(message)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(type_)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QErrorMessage11showMessageERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:68
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)
func (this *QErrorMessage) Done(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QErrorMessage4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QErrorMessage) ChangeEvent(e *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QErrorMessage11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
