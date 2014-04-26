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
	"unsafe"
)

type VirSecret struct {
	ptr C.virSecretPtr
}

func (s *VirSecret) Free() error {
	if result := C.virSecretFree(s.ptr); result != 0 {
		return errors.New(GetLastError())
	}
	return nil
}

func (s *VirSecret) Undefine() error {
	result := C.virSecretUndefine(s.ptr)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}


