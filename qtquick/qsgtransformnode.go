package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
import "qt.go/qtqml"

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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSGTransformNode struct {
	*QSGNode
}

func (this *QSGTransformNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGNode.GetCthis()
	}
}
func (this *QSGTransformNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGNode = NewQSGNodeFromPointer(cthis)
}
func NewQSGTransformNodeFromPointer(cthis unsafe.Pointer) *QSGTransformNode {
	bcthis0 := NewQSGNodeFromPointer(cthis)
	return &QSGTransformNode{bcthis0}
}
func (*QSGTransformNode) NewFromPointer(cthis unsafe.Pointer) *QSGTransformNode {
	return NewQSGTransformNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGTransformNode()
func NewQSGTransformNode() *QSGTransformNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGTransformNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGTransformNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:285
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGTransformNode()
func DeleteQSGTransformNode(this *QSGTransformNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 216)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix4x4 &)
func (this *QSGTransformNode) SetMatrix(matrix *qtgui.QMatrix4x4) {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNode9setMatrixERK10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:288
// index:0
// Public inline Visibility=Default Availability=Available
// [68] const QMatrix4x4 & matrix()
func (this *QSGTransformNode) Matrix() *qtgui.QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSGTransformNode6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:290
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCombinedMatrix(const QMatrix4x4 &)
func (this *QSGTransformNode) SetCombinedMatrix(matrix *qtgui.QMatrix4x4) {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNode17setCombinedMatrixERK10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:291
// index:0
// Public inline Visibility=Default Availability=Available
// [68] const QMatrix4x4 & combinedMatrix()
func (this *QSGTransformNode) CombinedMatrix() *qtgui.QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSGTransformNode14combinedMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix4x4)
	return rv2
}

//  body block end
