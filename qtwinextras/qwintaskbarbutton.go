package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h
// #include <qwintaskbarbutton.h>
// #include <Qtwinextras>

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

type QWinTaskbarButton struct {
	*qtcore.QObject
}

func (this *QWinTaskbarButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWinTaskbarButton) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWinTaskbarButtonFromPointer(cthis unsafe.Pointer) *QWinTaskbarButton {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWinTaskbarButton{bcthis0}
}
func (*QWinTaskbarButton) NewFromPointer(cthis unsafe.Pointer) *QWinTaskbarButton {
	return NewQWinTaskbarButtonFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QWinTaskbarButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWinTaskbarButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinTaskbarButton(QObject *)
func NewQWinTaskbarButton(parent *qtcore.QObject /*777 QObject **/) *QWinTaskbarButton {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWinTaskbarButtonC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinTaskbarButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinTaskbarButton()
func DeleteQWinTaskbarButton(this *QWinTaskbarButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWinTaskbarButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindow(QWindow *)
func (this *QWinTaskbarButton) SetWindow(window *qtgui.QWindow /*777 QWindow **/) {
	var convArg0 = window.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWinTaskbarButton9setWindowEP7QWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * window()
func (this *QWinTaskbarButton) Window() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWinTaskbarButton6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon overlayIcon()
func (this *QWinTaskbarButton) OverlayIcon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWinTaskbarButton11overlayIconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QString overlayAccessibleDescription()
func (this *QWinTaskbarButton) OverlayAccessibleDescription() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWinTaskbarButton28overlayAccessibleDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QWinTaskbarProgress * progress()
func (this *QWinTaskbarButton) Progress() *QWinTaskbarProgress /*777 QWinTaskbarProgress **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWinTaskbarButton8progressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQWinTaskbarProgressFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QWinTaskbarButton) EventFilter(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWinTaskbarButton11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverlayIcon(const QIcon &)
func (this *QWinTaskbarButton) SetOverlayIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWinTaskbarButton14setOverlayIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverlayAccessibleDescription(const QString &)
func (this *QWinTaskbarButton) SetOverlayAccessibleDescription(description string) {
	var tmpArg0 = qtcore.NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWinTaskbarButton31setOverlayAccessibleDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearOverlayIcon()
func (this *QWinTaskbarButton) ClearOverlayIcon() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWinTaskbarButton16clearOverlayIconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
