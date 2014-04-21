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

// TODO: "type" is a builtin type in Go.Won't compile
//func (i *VirStorageVolInfo) GetType() int {
//	return int(i.ptr.type)
//}

func (i *VirStorageVolInfo) GetCapacityInBytes() uint64 {
	return uint64(i.ptr.capacity)
}

func (i *VirStorageVolInfo) GetAllocationInBytes() uint64 {
	return uint64(i.ptr.allocation)
}

func (v *VirStorageVol) GetKey() (string, error) {
	key := C.virStorageVolGetKey(v.ptr)
	if key == nil {
		return "", errors.New(GetLastError())
	}
	return C.GoString(key), nil
}

func (v *VirStorageVol) GetName() (string, error) {
	name := C.virStorageVolGetName(v.ptr)
	if name == nil {
		return "", errors.New(GetLastError())
	}
	return C.GoString(name), nil
}

func (v *VirStorageVol) GetPath() (string, error) {
	result := C.virStorageVolGetPath(v.ptr)
	if result == nil {
		return "", errors.New(GetLastError())
	}
	path := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return path, nil
}

func (v *VirStorageVol) GetXMLDesc(flags uint32) (string, error) {
	result := C.virStorageVolGetXMLDesc(v.ptr, C.uint(flags))
	if result == nil {
		return "", errors.New(GetLastError())
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}
