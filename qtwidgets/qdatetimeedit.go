package qtwidgets

// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
// void keyPressEvent(class QKeyEvent *)
func (this *QDateTimeEdit) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QDateTimeEdit) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QDateTimeEdit) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QDateTimeEdit) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// QValidator::State validate(class QString &, int &)
func (this *QDateTimeEdit) InheritValidate(f func(input string, pos int) int) {
	qtrt.SetAllInheritCallback(this, "validate", f)
}

// void fixup(class QString &)
func (this *QDateTimeEdit) InheritFixup(f func(input string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "fixup", f)
}

// QDateTime dateTimeFromText(const class QString &)
func (this *QDateTimeEdit) InheritDateTimeFromText(f func(text string) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "dateTimeFromText", f)
}

// QString textFromDateTime(const class QDateTime &)
func (this *QDateTimeEdit) InheritTextFromDateTime(f func(dt *qtcore.QDateTime) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "textFromDateTime", f)
}

// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QDateTimeEdit) InheritStepEnabled(f func() int) {
	qtrt.SetAllInheritCallback(this, "stepEnabled", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QDateTimeEdit) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QDateTimeEdit) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QDateTimeEdit) InheritInitStyleOption(f func(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QDateTimeEdit struct {
	*QAbstractSpinBox
}

func (this *QDateTimeEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSpinBox.GetCthis()
	}
}
func (this *QDateTimeEdit) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSpinBox = NewQAbstractSpinBoxFromPointer(cthis)
}
func NewQDateTimeEditFromPointer(cthis unsafe.Pointer) *QDateTimeEdit {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QDateTimeEdit{bcthis0}
}
func (*QDateTimeEdit) NewFromPointer(cthis unsafe.Pointer) *QDateTimeEdit {
	return NewQDateTimeEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QDateTimeEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(QWidget *)
func NewQDateTimeEdit(parent *QWidget /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QDateTime &, QWidget *)
func NewQDateTimeEdit_1(dt *qtcore.QDateTime, parent *QWidget /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 = dt.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK9QDateTimeP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:97
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QDate &, QWidget *)
func NewQDateTimeEdit_2(d *qtcore.QDate, parent *QWidget /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 = d.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QDateP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:98
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QTime &, QWidget *)
func NewQDateTimeEdit_3(t *qtcore.QTime, parent *QWidget /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 = t.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QTimeP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:194
// index:4
// Protected Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QVariant &, QVariant::Type, QWidget *)
func NewQDateTimeEdit_4(val *qtcore.QVariant, parserType int, parent *QWidget /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 = val.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK8QVariantNS0_4TypeEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, parserType, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDateTimeEdit()
func DeleteQDateTimeEdit(this *QDateTimeEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime dateTime()
func (this *QDateTimeEdit) DateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8dateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate date()
func (this *QDateTimeEdit) Date() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit4dateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime time()
func (this *QDateTimeEdit) Time() *qtcore.QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit4timeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime minimumDateTime()
func (this *QDateTimeEdit) MinimumDateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit15minimumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMinimumDateTime()
func (this *QDateTimeEdit) ClearMinimumDateTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMinimumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDateTime(const QDateTime &)
func (this *QDateTimeEdit) SetMinimumDateTime(dt *qtcore.QDateTime) {
	var convArg0 = dt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime maximumDateTime()
func (this *QDateTimeEdit) MaximumDateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit15maximumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMaximumDateTime()
func (this *QDateTimeEdit) ClearMaximumDateTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMaximumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumDateTime(const QDateTime &)
func (this *QDateTimeEdit) SetMaximumDateTime(dt *qtcore.QDateTime) {
	var convArg0 = dt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateTimeRange(const QDateTime &, const QDateTime &)
func (this *QDateTimeEdit) SetDateTimeRange(min *qtcore.QDateTime, max *qtcore.QDateTime) {
	var convArg0 = min.GetCthis()
	var convArg1 = max.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate minimumDate()
func (this *QDateTimeEdit) MinimumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDate(const QDate &)
func (this *QDateTimeEdit) SetMinimumDate(min *qtcore.QDate) {
	var convArg0 = min.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMinimumDate()
func (this *QDateTimeEdit) ClearMinimumDate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate maximumDate()
func (this *QDateTimeEdit) MaximumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumDate(const QDate &)
func (this *QDateTimeEdit) SetMaximumDate(max *qtcore.QDate) {
	var convArg0 = max.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMaximumDate()
func (this *QDateTimeEdit) ClearMaximumDate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateRange(const QDate &, const QDate &)
func (this *QDateTimeEdit) SetDateRange(min *qtcore.QDate, max *qtcore.QDate) {
	var convArg0 = min.GetCthis()
	var convArg1 = max.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit12setDateRangeERK5QDateS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime minimumTime()
func (this *QDateTimeEdit) MinimumTime() *qtcore.QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumTime(const QTime &)
func (this *QDateTimeEdit) SetMinimumTime(min *qtcore.QTime) {
	var convArg0 = min.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMinimumTime()
func (this *QDateTimeEdit) ClearMinimumTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:129
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime maximumTime()
func (this *QDateTimeEdit) MaximumTime() *qtcore.QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumTime(const QTime &)
func (this *QDateTimeEdit) SetMaximumTime(max *qtcore.QTime) {
	var convArg0 = max.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMaximumTime()
func (this *QDateTimeEdit) ClearMaximumTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeRange(const QTime &, const QTime &)
func (this *QDateTimeEdit) SetTimeRange(min *qtcore.QTime, max *qtcore.QTime) {
	var convArg0 = min.GetCthis()
	var convArg1 = max.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] QDateTimeEdit::Sections displayedSections()
func (this *QDateTimeEdit) DisplayedSections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit17displayedSectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] QDateTimeEdit::Section currentSection()
func (this *QDateTimeEdit) CurrentSection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit14currentSectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] QDateTimeEdit::Section sectionAt(int)
func (this *QDateTimeEdit) SectionAt(index int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit9sectionAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetCurrentSection(section int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCurrentSectionENS_7SectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentSectionIndex()
func (this *QDateTimeEdit) CurrentSectionIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit19currentSectionIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentSectionIndex(int)
func (this *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit22setCurrentSectionIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QCalendarWidget * calendarWidget()
func (this *QDateTimeEdit) CalendarWidget() *QCalendarWidget /*777 QCalendarWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit14calendarWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCalendarWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCalendarWidget(QCalendarWidget *)
func (this *QDateTimeEdit) SetCalendarWidget(calendarWidget *QCalendarWidget /*777 QCalendarWidget **/) {
	var convArg0 = calendarWidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionCount()
func (this *QDateTimeEdit) SectionCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit12sectionCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectedSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetSelectedSection(section int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18setSelectedSectionENS_7SectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sectionText(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SectionText(section int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11sectionTextENS_7SectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayFormat()
func (this *QDateTimeEdit) DisplayFormat() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit13displayFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDisplayFormat(const QString &)
func (this *QDateTimeEdit) SetDisplayFormat(format string) {
	var tmpArg0 = qtcore.NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDisplayFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:155
// index:0
// Public Visibility=Default Availability=Available
// [1] bool calendarPopup()
func (this *QDateTimeEdit) CalendarPopup() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit13calendarPopupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCalendarPopup(_Bool)
func (this *QDateTimeEdit) SetCalendarPopup(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16setCalendarPopupEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:158
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TimeSpec timeSpec()
func (this *QDateTimeEdit) TimeSpec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8timeSpecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeSpec(Qt::TimeSpec)
func (this *QDateTimeEdit) SetTimeSpec(spec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11setTimeSpecEN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:161
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QDateTimeEdit) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:163
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()
func (this *QDateTimeEdit) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:164
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void stepBy(int)
func (this *QDateTimeEdit) StepBy(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit6stepByEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QDateTimeEdit) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dateTimeChanged(const QDateTime &)
func (this *QDateTimeEdit) DateTimeChanged(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit15dateTimeChangedERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void timeChanged(const QTime &)
func (this *QDateTimeEdit) TimeChanged(time *qtcore.QTime) {
	var convArg0 = time.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11timeChangedERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dateChanged(const QDate &)
func (this *QDateTimeEdit) DateChanged(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11dateChangedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateTime(const QDateTime &)
func (this *QDateTimeEdit) SetDateTime(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11setDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDate(const QDate &)
func (this *QDateTimeEdit) SetDate(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit7setDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTime(const QTime &)
func (this *QDateTimeEdit) SetTime(time *qtcore.QTime) {
	var convArg0 = time.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit7setTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:178
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QDateTimeEdit) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:180
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QDateTimeEdit) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:182
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QDateTimeEdit) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QDateTimeEdit) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:184
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QDateTimeEdit) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &)
func (this *QDateTimeEdit) Fixup(input string) {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QDateTime dateTimeFromText(const QString &)
func (this *QDateTimeEdit) DateTimeFromText(text string) *qtcore.QDateTime /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QString textFromDateTime(const QDateTime &)
func (this *QDateTimeEdit) TextFromDateTime(dt *qtcore.QDateTime) string {
	var convArg0 = dt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:189
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QDateTimeEdit) StepEnabled() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11stepEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:190
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QDateTimeEdit) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QDateTimeEdit) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:192
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSpinBox *)
func (this *QDateTimeEdit) InitStyleOption(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QDateTimeEdit__Section = int

const QDateTimeEdit__NoSection QDateTimeEdit__Section = 0
const QDateTimeEdit__AmPmSection QDateTimeEdit__Section = 1
const QDateTimeEdit__MSecSection QDateTimeEdit__Section = 2
const QDateTimeEdit__SecondSection QDateTimeEdit__Section = 4
const QDateTimeEdit__MinuteSection QDateTimeEdit__Section = 8
const QDateTimeEdit__HourSection QDateTimeEdit__Section = 16
const QDateTimeEdit__DaySection QDateTimeEdit__Section = 256
const QDateTimeEdit__MonthSection QDateTimeEdit__Section = 512
const QDateTimeEdit__YearSection QDateTimeEdit__Section = 1024
const QDateTimeEdit__TimeSections_Mask QDateTimeEdit__Section = 31
const QDateTimeEdit__DateSections_Mask QDateTimeEdit__Section = 1792

//  body block end
