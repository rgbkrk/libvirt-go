// +build libvirt.1.2.14

package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

const (
	VIR_STORAGE_VOL_CREATE_REFLINK = C.VIR_STORAGE_VOL_CREATE_REFLINK
)
