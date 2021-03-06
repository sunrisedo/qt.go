package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QRegularExpressionMatch struct {
	*qtrt.CObject
}
type QRegularExpressionMatch_ITF interface {
	QRegularExpressionMatch_PTR() *QRegularExpressionMatch
}

func (ptr *QRegularExpressionMatch) QRegularExpressionMatch_PTR() *QRegularExpressionMatch { return ptr }

func (this *QRegularExpressionMatch) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegularExpressionMatch) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegularExpressionMatchFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatch {
	return &QRegularExpressionMatch{&qtrt.CObject{cthis}}
}
func (*QRegularExpressionMatch) NewFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatch {
	return NewQRegularExpressionMatchFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregularexpression.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionMatch()

/*

 */
func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatchC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpressionMatch)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegularExpressionMatch()

/*

 */
func DeleteQRegularExpressionMatch(this *QRegularExpressionMatch) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatchD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregularexpression.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch & operator=(const QRegularExpressionMatch &)

/*

 */
func (this *QRegularExpressionMatch) Operator_equal(match_ QRegularExpressionMatch_ITF) *QRegularExpressionMatch {
	var convArg0 unsafe.Pointer
	if match_ != nil && match_.QRegularExpressionMatch_PTR() != nil {
		convArg0 = match_.QRegularExpressionMatch_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatchaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:184
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QRegularExpressionMatch & operator=(QRegularExpressionMatch &&)

/*

 */
func (this *QRegularExpressionMatch) Operator_equal_1(match_ unsafe.Pointer /*333*/) *QRegularExpressionMatch {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatchaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), match_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegularExpressionMatch &)

/*
Swaps the regular expression other with this regular expression. This operation is very fast and never fails.
*/
func (this *QRegularExpressionMatch) Swap(other QRegularExpressionMatch_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRegularExpressionMatch_PTR() != nil {
		convArg0 = other.QRegularExpressionMatch_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatch4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression regularExpression() const

/*

 */
func (this *QRegularExpressionMatch) RegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17regularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:190
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchType matchType() const

/*

 */
func (this *QRegularExpressionMatch) MatchType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch9matchTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:191
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchOptions matchOptions() const

/*

 */
func (this *QRegularExpressionMatch) MatchOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12matchOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:193
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasMatch() const

/*

 */
func (this *QRegularExpressionMatch) HasMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8hasMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:194
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasPartialMatch() const

/*

 */
func (this *QRegularExpressionMatch) HasPartialMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch15hasPartialMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:196
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the regular expression is a valid regular expression (that is, it contains no syntax errors, etc.), or false otherwise. Use errorString() to obtain a textual description of the error.

See also errorString() and patternErrorOffset().
*/
func (this *QRegularExpressionMatch) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:198
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastCapturedIndex() const

/*

 */
func (this *QRegularExpressionMatch) LastCapturedIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17lastCapturedIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QString captured(int) const

/*

 */
func (this *QRegularExpressionMatch) Captured(nth int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QString captured(int) const

/*

 */
func (this *QRegularExpressionMatch) Captured__() string {
	// arg: 0, int=Int, =Invalid,
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:205
// index:1
// Public Visibility=Default Availability=Available
// [8] QString captured(const QString &) const

/*

 */
func (this *QRegularExpressionMatch) Captured_1(name string) string {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:209
// index:2
// Public Visibility=Default Availability=Available
// [8] QString captured(QStringView) const

/*

 */
func (this *QRegularExpressionMatch) Captured_2(name QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if name != nil && name.QStringView_PTR() != nil {
		convArg0 = name.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:201
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef capturedRef(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedRef(nth int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:201
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef capturedRef(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedRef__() *QStringRef /*123*/ {
	// arg: 0, int=Int, =Invalid,
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:206
// index:1
// Public Visibility=Default Availability=Available
// [16] QStringRef capturedRef(const QString &) const

/*

 */
func (this *QRegularExpressionMatch) CapturedRef_1(name string) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:210
// index:2
// Public Visibility=Default Availability=Available
// [16] QStringRef capturedRef(QStringView) const

/*

 */
func (this *QRegularExpressionMatch) CapturedRef_2(name QStringView_ITF /*123*/) *QStringRef /*123*/ {
	var convArg0 unsafe.Pointer
	if name != nil && name.QStringView_PTR() != nil {
		convArg0 = name.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:202
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringView capturedView(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedView(nth int) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:202
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringView capturedView(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedView__() *QStringView /*123*/ {
	// arg: 0, int=Int, =Invalid,
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:211
// index:1
// Public Visibility=Default Availability=Available
// [16] QStringView capturedView(QStringView) const

/*

 */
func (this *QRegularExpressionMatch) CapturedView_1(name QStringView_ITF /*123*/) *QStringView /*123*/ {
	var convArg0 unsafe.Pointer
	if name != nil && name.QStringView_PTR() != nil {
		convArg0 = name.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList capturedTexts() const

/*

 */
func (this *QRegularExpressionMatch) CapturedTexts() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedTextsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:215
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedStart(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedStart(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:215
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedStart(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedStart__() int {
	// arg: 0, int=Int, =Invalid,
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:220
// index:1
// Public Visibility=Default Availability=Available
// [4] int capturedStart(const QString &) const

/*

 */
func (this *QRegularExpressionMatch) CapturedStart_1(name string) int {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:225
// index:2
// Public Visibility=Default Availability=Available
// [4] int capturedStart(QStringView) const

/*

 */
func (this *QRegularExpressionMatch) CapturedStart_2(name QStringView_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if name != nil && name.QStringView_PTR() != nil {
		convArg0 = name.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:216
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedLength(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedLength(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:216
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedLength(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedLength__() int {
	// arg: 0, int=Int, =Invalid,
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:221
// index:1
// Public Visibility=Default Availability=Available
// [4] int capturedLength(const QString &) const

/*

 */
func (this *QRegularExpressionMatch) CapturedLength_1(name string) int {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:226
// index:2
// Public Visibility=Default Availability=Available
// [4] int capturedLength(QStringView) const

/*

 */
func (this *QRegularExpressionMatch) CapturedLength_2(name QStringView_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if name != nil && name.QStringView_PTR() != nil {
		convArg0 = name.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:217
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedEnd(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedEnd(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:217
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedEnd(int) const

/*

 */
func (this *QRegularExpressionMatch) CapturedEnd__() int {
	// arg: 0, int=Int, =Invalid,
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:222
// index:1
// Public Visibility=Default Availability=Available
// [4] int capturedEnd(const QString &) const

/*

 */
func (this *QRegularExpressionMatch) CapturedEnd_1(name string) int {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:227
// index:2
// Public Visibility=Default Availability=Available
// [4] int capturedEnd(QStringView) const

/*

 */
func (this *QRegularExpressionMatch) CapturedEnd_2(name QStringView_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if name != nil && name.QStringView_PTR() != nil {
		convArg0 = name.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
}

//  keep block end
