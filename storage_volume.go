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

type VirStorageVol struct {
	ptr C.virStorageVolPtr
}

type VirStorageVolInfo struct {
	ptr C.virStorageVolInfo
}

func (v *VirStorageVol) Delete(flags uint32) error {
	result := C.virStorageVolDelete(v.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (v *VirStorageVol) Free() error {
	if result := C.virStorageVolFree(v.ptr); result != 0 {
		return errors.New(GetLastError())
	}
	return nil
}

func (v *VirStorageVol) GetInfo() (VirStorageVolInfo, error) {
	vi := VirStorageVolInfo{}
	var ptr C.virStorageVolInfo
	result := C.virStorageVolGetInfo(v.ptr, (*C.virStorageVolInfo)(unsafe.Pointer(&ptr)))
	if result == -1 {
		return vi, errors.New(GetLastError())
	}
	vi.ptr = ptr
	return vi, nil
}
