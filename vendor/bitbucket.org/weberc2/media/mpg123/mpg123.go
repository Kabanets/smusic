// This package is a thin wrapper around libmpg123.
// API documentation [here](http://www.mpg123.de/api/)
package mpg123

/*
#cgo LDFLAGS: -lmpg123
#include <mpg123.h>
*/
import "C"

import (
	"fmt"
	"io"
	"runtime"
	"unsafe"
)

func Initialize() {
	C.mpg123_init()
}

func Exit() {
	C.mpg123_exit()
}

type Handle struct {
	inner *C.mpg123_handle
}

func mpg123PlainStrError(errcode C.int) string {
	return C.GoString(C.mpg123_plain_strerror(errcode))
}

func isError(errcode C.int) bool {
	return errcode != C.MPG123_OK
}

func mpg123Error(errcode C.int) error {
	if isError(errcode) {

		// check for io.EOF
		if errcode == C.MPG123_DONE {
			return io.EOF
		}

		s := mpg123PlainStrError(errcode)
		return fmt.Errorf("MPG123 ERR %v: %s", errcode, s)
	}
	return nil
}

func (h *Handle) Close() error {
	return mpg123Error(C.mpg123_close(h.inner))
}

func newInnerHandle() (*C.mpg123_handle, error) {
	var err C.int
	handle := C.mpg123_new(nil, &err)
	if err := mpg123Error(err); err != nil {
		return nil, err
	}
	return handle, nil
}

func (h *Handle) free() {
	C.mpg123_delete(h.inner)
}

func newHandle() (*Handle, error) {
	inner, err := newInnerHandle()
	if err != nil {
		return nil, err
	}

	h := &Handle{inner}
	runtime.SetFinalizer(h, (*Handle).free)

	return h, nil
}

func openHandle(handle *C.mpg123_handle, path string) error {
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))
	return mpg123Error(C.mpg123_open(handle, cs))
}

func Open(path string) (*Handle, error) {
	h, err := newHandle()
	if err != nil {
		return nil, err
	}
	if err := openHandle(h.inner, path); err != nil {
		return nil, err
	}
	return h, nil
}

// This is mostly a thin wrapper around mpg123_read() with the notable
// exception that MPG_DONE return is interpreted as io.EOF
func (h *Handle) Read(p []byte) (int, error) {
	var bytesRead C.size_t
	buffer := (*C.uchar)(unsafe.Pointer(&p[0]))
	bufferSize := (C.size_t)(len(p))
	err := C.mpg123_read(h.inner, buffer, bufferSize, &bytesRead)
	if err := mpg123Error(err); err != nil {
		return int(bytesRead), err
	}
	return int(bytesRead), nil
}

func (h *Handle) Format() (rate int64, channels int, encoding int) {
	c_rate := (*C.long)(unsafe.Pointer(&rate))
	c_channels := (*C.int)(unsafe.Pointer(&channels))
	c_encoding := (*C.int)(unsafe.Pointer(&encoding))
	C.mpg123_getformat(h.inner, c_rate, c_channels, c_encoding)
	return
}

func (h *Handle) EncodingSize(encoding int) int {
	return int(C.mpg123_encsize(C.int(encoding)))
}

func (h *Handle) Outblock() int {
	return int(C.mpg123_outblock(h.inner))
}
