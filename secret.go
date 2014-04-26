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

func (s *VirSecret) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virSecretGetUUID(s.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, errors.New(GetLastError())
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (s *VirSecret) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virSecretGetUUIDString(s.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", errors.New(GetLastError())
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}

