// This package is just a thin wrapper around libao.
// Defer to its [API documentation](http://www.xiph.org/ao/doc/overview.html).
package ao

/*
#cgo LDFLAGS: -lao
#include <ao/ao.h>
*/
import "C"
import "fmt"
import "unsafe"

func Initialize() {
	C.ao_initialize()
}

func Shutdown() {
	C.ao_shutdown()
}

type ByteFormat int

const (
	FormatNative ByteFormat = C.AO_FMT_NATIVE
)

type SampleFormat struct {
	BitsPerSample int
	Rate          int
	Channels      int
	ByteFormat    ByteFormat
	Matrix        []byte
}

type Device struct {
	inner *C.ao_device
}

func initSampleFormat(aoFormat *C.ao_sample_format, format *SampleFormat) {
	aoFormat.bits = C.int(format.BitsPerSample)
	aoFormat.rate = C.int(format.Rate)
	aoFormat.channels = C.int(format.Channels)
	aoFormat.byte_format = C.AO_FMT_NATIVE
	// aoFormat.matrix = format.Matrix // this is not supported on the libao version in Ubuntu 10.04 repos
}

// Wraps `ao_open_live()`
func NewDevice(format *SampleFormat, aoDriver int) *Device {
	var aoFormat C.ao_sample_format
	initSampleFormat(&aoFormat, format)
	return &Device{C.ao_open_live(C.int(aoDriver), &aoFormat, nil)}
}

// Wraps `ao_open_live()`, but uses the default driver id
// via `ao_default_driver_id()`
func NewLiveDevice(format *SampleFormat) *Device {
	return NewDevice(format, int(C.ao_default_driver_id()))
}

// `ao_close()` frees the memory that is used for the device;
// not calling this will leak memory
func (dev *Device) Close() error {
	if 0 == C.ao_close(dev.inner) {
		return fmt.Errorf("Error closing ao_device")
	}
	return nil
}

// Wraps `ao_play()`; satisfies `io.Writer`
func (dev *Device) Write(p []byte) (int, error) {
	err := C.ao_play(dev.inner, (*C.char)(unsafe.Pointer(&p[0])), C.uint_32(len(p)))
	if err == 0 {
		return 0, fmt.Errorf("ao_play error")
	}
	return len(p), nil
}
