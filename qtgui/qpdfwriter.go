package qtgui

// /usr/include/qt/QtGui/qpdfwriter.h
// #include <qpdfwriter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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

// QPaintEngine * paintEngine()
func (this *QPdfWriter) InheritPaintEngine(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "paintEngine", f)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPdfWriter) InheritMetric(f func(id int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

/*

 */
type QPdfWriter struct {
	*qtcore.QObject
	*QPagedPaintDevice
}
type QPdfWriter_ITF interface {
	qtcore.QObject_ITF
	QPagedPaintDevice_ITF
	QPdfWriter_PTR() *QPdfWriter
}

func (ptr *QPdfWriter) QPdfWriter_PTR() *QPdfWriter { return ptr }

func (this *QPdfWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QPdfWriter) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QPagedPaintDevice = NewQPagedPaintDeviceFromPointer(cthis)
}
func NewQPdfWriterFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQPagedPaintDeviceFromPointer(cthis)
	return &QPdfWriter{bcthis0, bcthis1}
}
func (*QPdfWriter) NewFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	return NewQPdfWriterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpdfwriter.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPdfWriter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPdfWriter10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpdfwriter.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPdfWriter(const QString &)

/*
Constructs a PDF writer that will write the pdf to filename.
*/
func NewQPdfWriter(filename string) *QPdfWriter {
	var tmpArg0 = qtcore.NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriterC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPdfWriter")
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPdfWriter(QIODevice *)

/*
Constructs a PDF writer that will write the pdf to filename.
*/
func NewQPdfWriter_1(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QPdfWriter {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriterC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPdfWriter")
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPdfWriter()

/*

 */
func DeleteQPdfWriter(this *QPdfWriter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpdfwriter.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPdfVersion(enum QPagedPaintDevice::PdfVersion)

/*
Sets the PDF version for this writer to version.

If version is the same value as currently set then no change will be made.

This function was introduced in  Qt 5.10.

See also pdfVersion().
*/
func (this *QPdfWriter) SetPdfVersion(version int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriter13setPdfVersionEN17QPagedPaintDevice10PdfVersionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), version)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] QPagedPaintDevice::PdfVersion pdfVersion() const

/*
Returns the PDF version for this writer. The default is PdfVersion_1_4.

This function was introduced in  Qt 5.10.

See also setPdfVersion().
*/
func (this *QPdfWriter) PdfVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPdfWriter10pdfVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*
Returns the title of the document.

See also setTitle().
*/
func (this *QPdfWriter) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPdfWriter5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpdfwriter.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)

/*
Sets the title of the document being created to title.

See also title().
*/
func (this *QPdfWriter) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriter8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString creator() const

/*
Returns the creator of the document.

See also setCreator().
*/
func (this *QPdfWriter) Creator() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPdfWriter7creatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpdfwriter.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCreator(const QString &)

/*
Sets the creator of the document to creator.

See also creator().
*/
func (this *QPdfWriter) SetCreator(creator string) {
	var tmpArg0 = qtcore.NewQString_5(creator)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriter10setCreatorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool newPage()

/*
Reimplemented from QPagedPaintDevice::newPage().
*/
func (this *QPdfWriter) NewPage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriter7newPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpdfwriter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolution(int)

/*
Sets the PDF resolution in DPI.

This setting affects the coordinate system as returned by, for example QPainter::viewport().

This function was introduced in  Qt 5.3.

See also resolution().
*/
func (this *QPdfWriter) SetResolution(resolution int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriter13setResolutionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int resolution() const

/*
Returns the resolution of the PDF in DPI.

This function was introduced in  Qt 5.3.

See also setResolution().
*/
func (this *QPdfWriter) Resolution() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPdfWriter10resolutionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpdfwriter.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSize(enum QPagedPaintDevice::PageSize)

/*
Sets the PDF page size to pageSize.

To get the current QPageSize use pageLayout().pageSize().

You should call this before calling QPainter::begin(), or immediately before calling newPage() to apply the new page size to a new page. You should not call any painting methods between a call to setPageSize() and newPage() as the wrong paint metrics may be used.

Returns true if the page size was successfully set to pageSize.

This function was introduced in  Qt 5.3.

See also pageLayout().
*/
func (this *QPdfWriter) SetPageSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriter11setPageSizeEN17QPagedPaintDevice8PageSizeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSizeMM(const QSizeF &)

/*

 */
func (this *QPdfWriter) SetPageSizeMM(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPdfWriter13setPageSizeMMERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*
Reimplemented from QPaintDevice::paintEngine().
*/
func (this *QPdfWriter) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPdfWriter11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpdfwriter.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric) const

/*

 */
func (this *QPdfWriter) Metric(id int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
