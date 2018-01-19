//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		ffiqt.KeepMe()
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
type QGraphicsSceneWheelEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:139
// index:0
// void QGraphicsSceneWheelEvent(enum QEvent::Type)
func NewQGraphicsSceneWheelEvent(type_ int) *QGraphicsSceneWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	return &QGraphicsSceneWheelEvent{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:140
// index:0
// virtual
// void ~QGraphicsSceneWheelEvent()
func DeleteQGraphicsSceneWheelEvent(*QGraphicsSceneWheelEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:142
// index:0
// QPointF pos()
func (this *QGraphicsSceneWheelEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:143
// index:0
// void setPos(const class QPointF &)
func (this *QGraphicsSceneWheelEvent) SetPos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:145
// index:0
// QPointF scenePos()
func (this *QGraphicsSceneWheelEvent) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent8scenePosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:146
// index:0
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneWheelEvent) SetScenePos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:148
// index:0
// QPoint screenPos()
func (this *QGraphicsSceneWheelEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:149
// index:0
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneWheelEvent) SetScreenPos(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:151
// index:0
// Qt::MouseButtons buttons()
func (this *QGraphicsSceneWheelEvent) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent7buttonsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:154
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneWheelEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:157
// index:0
// int delta()
func (this *QGraphicsSceneWheelEvent) Delta() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent5deltaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:158
// index:0
// void setDelta(int)
func (this *QGraphicsSceneWheelEvent) SetDelta(delta int) {
	// 0: (, int delta), (&delta)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent8setDeltaEi", ffiqt.FFI_TYPE_VOID, this.cthis, &delta)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:160
// index:0
// Qt::Orientation orientation()
func (this *QGraphicsSceneWheelEvent) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent11orientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:161
// index:0
// void setOrientation(Qt::Orientation)
func (this *QGraphicsSceneWheelEvent) SetOrientation(orientation int) {
	// 0: (, Qt::Orientation orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation)
	gopp.ErrPrint(err, rv)
}

//  body block end