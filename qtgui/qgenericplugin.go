package qtgui

// /usr/include/qt/QtGui/qgenericplugin.h
// #include <qgenericplugin.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 37
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QGenericPlugin struct {
	*qtcore.QObject
}
type QGenericPlugin_ITF interface {
	qtcore.QObject_ITF
	QGenericPlugin_PTR() *QGenericPlugin
}

func (ptr *QGenericPlugin) QGenericPlugin_PTR() *QGenericPlugin { return ptr }

func (this *QGenericPlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGenericPlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGenericPluginFromPointer(cthis unsafe.Pointer) *QGenericPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGenericPlugin{bcthis0}
}
func (*QGenericPlugin) NewFromPointer(cthis unsafe.Pointer) *QGenericPlugin {
	return NewQGenericPluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qgenericplugin.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGenericPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGenericPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qgenericplugin.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGenericPlugin(QObject *)

/*
Constructs a plugin with the given parent.

Note that this constructor is invoked automatically by the moc generated code that exports the plugin, so there is no need for calling it explicitly.
*/
func NewQGenericPlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QGenericPlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGenericPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGenericPlugin")
	return gothis
}

// /usr/include/qt/QtGui/qgenericplugin.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGenericPlugin(QObject *)

/*
Constructs a plugin with the given parent.

Note that this constructor is invoked automatically by the moc generated code that exports the plugin, so there is no need for calling it explicitly.
*/
func NewQGenericPlugin__() *QGenericPlugin {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGenericPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGenericPlugin")
	return gothis
}

// /usr/include/qt/QtGui/qgenericplugin.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGenericPlugin()

/*

 */
func DeleteQGenericPlugin(this *QGenericPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGenericPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qgenericplugin.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QObject * create(const QString &, const QString &)

/*
Implement this function to create a driver matching the type specified by the given key and specification parameters. Note that keys are case-insensitive.
*/
func (this *QGenericPlugin) Create(name string, spec string) *qtcore.QObject /*777 QObject **/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(spec)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGenericPlugin6createERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
