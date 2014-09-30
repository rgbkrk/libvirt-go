package libvirt

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>

int streamEventCallback_cgo(virStreamPtr stream, int events, void *opaque);
*/
import "C"
import (
	"io"
	"unsafe"
)

type VirStream struct {
	ptr C.virStreamPtr
}

func NewVirStream(c *VirConnection, flags uint) (*VirStream, error) {
	virStream := C.virStreamNew(c.ptr, C.uint(flags))
	if virStream == nil {
		return nil, GetLastError()
	}

	return &VirStream{
		ptr: virStream,
	}, nil
}

func (v *VirStream) Abort() error {
	result := C.virStreamAbort(v.ptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (v *VirStream) Close() error {
	result := C.virStreamFinish(v.ptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (v *VirStream) Free() error {
	result := C.virStreamFree(v.ptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (v *VirStream) Read(p []byte) (int, error) {
	n := C.virStreamRecv(v.ptr, (*C.char)(unsafe.Pointer(&p[0])), C.size_t(len(p)))
	if n < 0 {
		return 0, GetLastError()
	}
	if n == 0 {
		return 0, io.EOF
	}

	return int(n), nil
}

func (v *VirStream) Write(p []byte) (int, error) {
	n := C.virStreamSend(v.ptr, (*C.char)(unsafe.Pointer(&p[0])), C.size_t(len(p)))
	if n < 0 {
		return 0, GetLastError()
	}
	if n == 0 {
		return 0, io.EOF
	}

	return int(n), nil
}

type VirStreamEventCallback func(s *VirStream, events int, f func()) int

type streamEventContext struct {
	cb *VirStreamEventCallback
	f  func()
}

func (v *VirStream) EventAddCallback(events int, cb *VirStreamEventCallback, f func()) error {
	context := streamEventContext{cb: cb, f: f}

	callbackPtr := unsafe.Pointer(C.streamEventCallback_cgo)
	ret := C.virStreamEventAddCallback(v.ptr, C.int(events),
		C.virStreamEventCallback(callbackPtr), unsafe.Pointer(&context), nil)
	if int(ret) < 0 {
		return GetLastError()
	}
	return nil
}

func (v *VirStream) EventUpdateCallback(events int) error {
	ret := C.virStreamEventUpdateCallback(v.ptr, C.int(events))
	if int(ret) < 0 {
		return GetLastError()
	}
	return nil
}

func (v *VirStream) EventRemoveCallback() error {
	ret := C.virStreamEventRemoveCallback(v.ptr)
	if int(ret) < 0 {
		return GetLastError()
	}
	return nil
}
