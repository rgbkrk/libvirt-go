package libvirt

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
)

type VirNWFilter struct {
	ptr C.virNWFilterPtr
}

func (f *VirNWFilter) Free() error {
	if result := C.virNWFilterFree(f.ptr); result != 0 {
		return errors.New(GetLastError())
	}
	return nil
}

func (f *VirNWFilter) GetName() (string, error) {
	name := C.virNWFilterGetName(f.ptr)
	if name == nil {
		return "", errors.New(GetLastError())
	}
	return C.GoString(name), nil
}

func (f *VirNWFilter) Undefine() error {
	result := C.virNWFilterUndefine(f.ptr)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}


