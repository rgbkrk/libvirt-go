package libvirt

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

// virDomainState
const (
	VIR_DOMAIN_NOSTATE     = C.VIR_DOMAIN_NOSTATE
	VIR_DOMAIN_RUNNING     = C.VIR_DOMAIN_RUNNING
	VIR_DOMAIN_BLOCKED     = C.VIR_DOMAIN_BLOCKED
	VIR_DOMAIN_PAUSED      = C.VIR_DOMAIN_PAUSED
	VIR_DOMAIN_SHUTDOWN    = C.VIR_DOMAIN_SHUTDOWN
	VIR_DOMAIN_CRASHED     = C.VIR_DOMAIN_CRASHED
	VIR_DOMAIN_PMSUSPENDED = C.VIR_DOMAIN_PMSUSPENDED
	VIR_DOMAIN_SHUTOFF     = C.VIR_DOMAIN_SHUTOFF
)

//virConnectListAllDomainsFlags
const (
	VIR_CONNECT_LIST_DOMAINS_ACTIVE         = C.VIR_CONNECT_LIST_DOMAINS_ACTIVE
	VIR_CONNECT_LIST_DOMAINS_INACTIVE       = C.VIR_CONNECT_LIST_DOMAINS_INACTIVE
	VIR_CONNECT_LIST_DOMAINS_PERSISTENT     = C.VIR_CONNECT_LIST_DOMAINS_PERSISTENT
	VIR_CONNECT_LIST_DOMAINS_TRANSIENT      = C.VIR_CONNECT_LIST_DOMAINS_TRANSIENT
	VIR_CONNECT_LIST_DOMAINS_RUNNING        = C.VIR_CONNECT_LIST_DOMAINS_RUNNING
	VIR_CONNECT_LIST_DOMAINS_PAUSED         = C.VIR_CONNECT_LIST_DOMAINS_PAUSED
	VIR_CONNECT_LIST_DOMAINS_SHUTOFF        = C.VIR_CONNECT_LIST_DOMAINS_SHUTOFF
	VIR_CONNECT_LIST_DOMAINS_OTHER          = C.VIR_CONNECT_LIST_DOMAINS_OTHER
	VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE    = C.VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE
	VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE = C.VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE
	VIR_CONNECT_LIST_DOMAINS_AUTOSTART      = C.VIR_CONNECT_LIST_DOMAINS_AUTOSTART
	VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART   = C.VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART
	VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT   = C.VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT
	VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT    = C.VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT
)

// virDomainMetadataType
const (
	VIR_DOMAIN_METADATA_DESCRIPTION = C.VIR_DOMAIN_METADATA_DESCRIPTION
	VIR_DOMAIN_METADATA_TITLE       = C.VIR_DOMAIN_METADATA_TITLE
	VIR_DOMAIN_METADATA_ELEMENT     = C.VIR_DOMAIN_METADATA_ELEMENT
)

// virDomainVcpuFlags
const (
	VIR_DOMAIN_VCPU_CONFIG  = C.VIR_DOMAIN_VCPU_CONFIG
	VIR_DOMAIN_VCPU_CURRENT = C.VIR_DOMAIN_VCPU_CURRENT
	VIR_DOMAIN_VCPU_LIVE    = C.VIR_DOMAIN_VCPU_LIVE
	VIR_DOMAIN_VCPU_MAXIMUM = C.VIR_DOMAIN_VCPU_MAXIMUM
	VIR_DOMAIN_VCPU_GUEST   = C.VIR_DOMAIN_VCPU_GUEST
)

// virDomainMemoryModFlags
const (
	VIR_DOMAIN_MEM_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG
	VIR_DOMAIN_MEM_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_MEM_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
	VIR_DOMAIN_MEM_MAXIMUM = C.VIR_DOMAIN_MEM_MAXIMUM
)

// virStoragePoolState
const (
	VIR_STORAGE_POOL_INACTIVE     = C.VIR_STORAGE_POOL_INACTIVE     // Not running
	VIR_STORAGE_POOL_BUILDING     = C.VIR_STORAGE_POOL_BUILDING     // Initializing pool,not available
	VIR_STORAGE_POOL_RUNNING      = C.VIR_STORAGE_POOL_RUNNING      // Running normally
	VIR_STORAGE_POOL_DEGRADED     = C.VIR_STORAGE_POOL_DEGRADED     // Running degraded
	VIR_STORAGE_POOL_INACCESSIBLE = C.VIR_STORAGE_POOL_INACCESSIBLE // Running,but not accessible
)

// virStoragePoolBuildFlags
const (
	VIR_STORAGE_POOL_BUILD_NEW          = C.VIR_STORAGE_POOL_BUILD_NEW          // Regular build from scratch
	VIR_STORAGE_POOL_BUILD_REPAIR       = C.VIR_STORAGE_POOL_BUILD_REPAIR       // Repair / reinitialize
	VIR_STORAGE_POOL_BUILD_RESIZE       = C.VIR_STORAGE_POOL_BUILD_RESIZE       // Extend existing pool
	VIR_STORAGE_POOL_BUILD_NO_OVERWRITE = C.VIR_STORAGE_POOL_BUILD_NO_OVERWRITE // Do not overwrite existing pool
	VIR_STORAGE_POOL_BUILD_OVERWRITE    = C.VIR_STORAGE_POOL_BUILD_OVERWRITE    // Overwrite data
)

// virDomainDestroyFlags
const (
	VIR_DOMAIN_DESTROY_DEFAULT  = C.VIR_DOMAIN_DESTROY_DEFAULT
	VIR_DOMAIN_DESTROY_GRACEFUL = C.VIR_DOMAIN_DESTROY_GRACEFUL
)

// virDomainShutdownFlags
const (
	VIR_DOMAIN_SHUTDOWN_DEFAULT        = C.VIR_DOMAIN_SHUTDOWN_DEFAULT
	VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN = C.VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN
	VIR_DOMAIN_SHUTDOWN_GUEST_AGENT    = C.VIR_DOMAIN_SHUTDOWN_GUEST_AGENT
	VIR_DOMAIN_SHUTDOWN_INITCTL        = C.VIR_DOMAIN_SHUTDOWN_INITCTL
	VIR_DOMAIN_SHUTDOWN_SIGNAL         = C.VIR_DOMAIN_SHUTDOWN_SIGNAL
)

// virDomainAttachDeviceFlags
const (
	VIR_DOMAIN_DEVICE_MODIFY_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG
	VIR_DOMAIN_DEVICE_MODIFY_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_DEVICE_MODIFY_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
	VIR_DOMAIN_DEVICE_MODIFY_FORCE   = C.VIR_DOMAIN_DEVICE_MODIFY_FORCE
)

// virStorageVolCreateFlags
const (
	VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA = C.VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA
)

// virStorageVolDeleteFlags
const (
	VIR_STORAGE_VOL_DELETE_NORMAL = C.VIR_STORAGE_VOL_DELETE_NORMAL // Delete metadata only (fast)
	VIR_STORAGE_VOL_DELETE_ZEROED = C.VIR_STORAGE_VOL_DELETE_ZEROED // Clear all data to zeros (slow)
)

// virStorageVolResizeFlags
const (
	VIR_STORAGE_VOL_RESIZE_ALLOCATE = C.VIR_STORAGE_VOL_RESIZE_ALLOCATE // force allocation of new size
	VIR_STORAGE_VOL_RESIZE_DELTA    = C.VIR_STORAGE_VOL_RESIZE_DELTA    // size is relative to current
	VIR_STORAGE_VOL_RESIZE_SHRINK   = C.VIR_STORAGE_VOL_RESIZE_SHRINK   // allow decrease in capacity
)

// virStorageVolType
const (
	VIR_STORAGE_VOL_FILE    = C.VIR_STORAGE_VOL_FILE    // Regular file based volumes
	VIR_STORAGE_VOL_BLOCK   = C.VIR_STORAGE_VOL_BLOCK   // Block based volumes
	VIR_STORAGE_VOL_DIR     = C.VIR_STORAGE_VOL_DIR     // Directory-passthrough based volume
	VIR_STORAGE_VOL_NETWORK = C.VIR_STORAGE_VOL_NETWORK //Network volumes like RBD (RADOS Block Device)
	VIR_STORAGE_VOL_NETDIR  = C.VIR_STORAGE_VOL_NETDIR  // Network accessible directory that can contain other network volumes
)

// virStorageVolWipeAlgorithm
const (
	VIR_STORAGE_VOL_WIPE_ALG_ZERO       = C.VIR_STORAGE_VOL_WIPE_ALG_ZERO       // 1-pass, all zeroes
	VIR_STORAGE_VOL_WIPE_ALG_NNSA       = C.VIR_STORAGE_VOL_WIPE_ALG_NNSA       // 4-pass NNSA Policy Letter NAP-14.1-C (XVI-8)
	VIR_STORAGE_VOL_WIPE_ALG_DOD        = C.VIR_STORAGE_VOL_WIPE_ALG_DOD        // 4-pass DoD 5220.22-M section 8-306 procedure
	VIR_STORAGE_VOL_WIPE_ALG_BSI        = C.VIR_STORAGE_VOL_WIPE_ALG_BSI        // 9-pass method recommended by the German Center of Security in Information Technologies
	VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    = C.VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    // The canonical 35-pass sequence
	VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   = C.VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   // 7-pass method described by Bruce Schneier in "Applied Cryptography" (1996)
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  = C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  // 7-pass random
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 = C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 // 33-pass random
	VIR_STORAGE_VOL_WIPE_ALG_RANDOM     = C.VIR_STORAGE_VOL_WIPE_ALG_RANDOM     // 1-pass random
)

// virSecretUsageType
const (
	VIR_SECRET_USAGE_TYPE_NONE   = C.VIR_SECRET_USAGE_TYPE_NONE
	VIR_SECRET_USAGE_TYPE_VOLUME = C.VIR_SECRET_USAGE_TYPE_VOLUME
	VIR_SECRET_USAGE_TYPE_CEPH   = C.VIR_SECRET_USAGE_TYPE_CEPH
	VIR_SECRET_USAGE_TYPE_ISCSI  = C.VIR_SECRET_USAGE_TYPE_ISCSI
)

// virConnectListAllNetworksFlags
const (
	VIR_CONNECT_LIST_NETWORKS_INACTIVE     = C.VIR_CONNECT_LIST_NETWORKS_INACTIVE
	VIR_CONNECT_LIST_NETWORKS_ACTIVE       = C.VIR_CONNECT_LIST_NETWORKS_ACTIVE
	VIR_CONNECT_LIST_NETWORKS_PERSISTENT   = C.VIR_CONNECT_LIST_NETWORKS_PERSISTENT
	VIR_CONNECT_LIST_NETWORKS_TRANSIENT    = C.VIR_CONNECT_LIST_NETWORKS_TRANSIENT
	VIR_CONNECT_LIST_NETWORKS_AUTOSTART    = C.VIR_CONNECT_LIST_NETWORKS_AUTOSTART
	VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART = C.VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART
)

// virConnectListAllStoragePoolsFlags
const (
	VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE     = C.VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE
	VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE       = C.VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE
	VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT   = C.VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT
	VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT    = C.VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT
	VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART    = C.VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART
	VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART = C.VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART
	VIR_CONNECT_LIST_STORAGE_POOLS_DIR          = C.VIR_CONNECT_LIST_STORAGE_POOLS_DIR
	VIR_CONNECT_LIST_STORAGE_POOLS_FS           = C.VIR_CONNECT_LIST_STORAGE_POOLS_FS
	VIR_CONNECT_LIST_STORAGE_POOLS_NETFS        = C.VIR_CONNECT_LIST_STORAGE_POOLS_NETFS
	VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL      = C.VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL
	VIR_CONNECT_LIST_STORAGE_POOLS_DISK         = C.VIR_CONNECT_LIST_STORAGE_POOLS_DISK
	VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI        = C.VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI
	VIR_CONNECT_LIST_STORAGE_POOLS_SCSI         = C.VIR_CONNECT_LIST_STORAGE_POOLS_SCSI
	VIR_CONNECT_LIST_STORAGE_POOLS_MPATH        = C.VIR_CONNECT_LIST_STORAGE_POOLS_MPATH
	VIR_CONNECT_LIST_STORAGE_POOLS_RBD          = C.VIR_CONNECT_LIST_STORAGE_POOLS_RBD
	VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG     = C.VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG
	VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER      = C.VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER
)

// virStreamFlags
const (
	VIR_STREAM_NONBLOCK = C.VIR_STREAM_NONBLOCK
)

// virKeycodeSet
const (
	VIR_KEYCODE_SET_LINUX  = C.VIR_KEYCODE_SET_LINUX
	VIR_KEYCODE_SET_XT     = C.VIR_KEYCODE_SET_XT
	VIR_KEYCODE_SET_ATSET1 = C.VIR_KEYCODE_SET_ATSET1
	VIR_KEYCODE_SET_ATSET2 = C.VIR_KEYCODE_SET_ATSET2
	VIR_KEYCODE_SET_ATSET3 = C.VIR_KEYCODE_SET_ATSET3
	VIR_KEYCODE_SET_OSX    = C.VIR_KEYCODE_SET_OSX
	VIR_KEYCODE_SET_XT_KBD = C.VIR_KEYCODE_SET_XT_KBD
	VIR_KEYCODE_SET_USB    = C.VIR_KEYCODE_SET_USB
	VIR_KEYCODE_SET_WIN32  = C.VIR_KEYCODE_SET_WIN32
	VIR_KEYCODE_SET_RFB    = C.VIR_KEYCODE_SET_RFB
)

// virDomainCreateFlags
const (
    VIR_DOMAIN_NONE = C.VIR_DOMAIN_NONE
    VIR_DOMAIN_START_PAUSED = C.VIR_DOMAIN_START_PAUSED
    VIR_DOMAIN_START_AUTODESTROY = C.VIR_DOMAIN_START_AUTODESTROY
    VIR_DOMAIN_START_BYPASS_CACHE = C.VIR_DOMAIN_START_BYPASS_CACHE
    VIR_DOMAIN_START_FORCE_BOOT = C.VIR_DOMAIN_START_FORCE_BOOT
)

const VIR_DOMAIN_MEMORY_PARAM_UNLIMITED = C.VIR_DOMAIN_MEMORY_PARAM_UNLIMITED
