package libvirt

import (
	"fmt"
	"sync"
	"unsafe"
)

/*
 * Golang 1.6 doesn't support C pointers to go memory.
 * A hacky-solution might be some multi-threaded approach to support domain events, but let's make it work
 * without domain events for now.
 */

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>

int domainEventLifecycleCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     int event, int detail, void* data);

int domainEventGenericCallback_cgo(virConnectPtr c, virDomainPtr d, void* data);

int domainEventRTCChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     long long utcoffset, void* data);

int domainEventWatchdogCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int action, void* data);

int domainEventIOErrorCallback_cgo(virConnectPtr c, virDomainPtr d,
                                   const char *srcPath, const char *devAlias,
                                   int action, void* data);

int domainEventGraphicsCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int phase, const virDomainEventGraphicsAddress *local,
                                    const virDomainEventGraphicsAddress *remote,
                                    const char *authScheme,
                                    const virDomainEventGraphicsSubject *subject, void* data);

int domainEventIOErrorReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *srcPath, const char *devAlias,
                                         int action, const char *reason, void* data);

int domainEventBlockJobCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    const char *disk, int type, int status, void* data);

int domainEventDiskChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *oldSrcPath, const char *newSrcPath,
                                      const char *devAlias, int reason, void* data);

int domainEventTrayChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *devAlias, int reason, void* data);

int domainEventReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                  int reason, void* data);

int domainEventBalloonChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         unsigned long long actual, void* data);

int domainEventDeviceRemovedCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *devAlias, void* data);

int virConnectDomainEventRegisterAny_cgo(virConnectPtr c,  virDomainPtr d,
						                             int eventID, virConnectDomainEventGenericCallback cb,
                                         int goCallbackId);
*/
import "C"

type DomainLifecycleEvent struct {
	Event  int
	Detail int
}

type DomainRTCChangeEvent struct {
	Utcoffset int64
}

type DomainWatchdogEvent struct {
	Action int
}

type DomainIOErrorEvent struct {
	SrcPath  string
	DevAlias string
	Action   int
}

type DomainEventGraphicsAddress struct {
	Family  int
	Node    string
	Service string
}

type DomainEventGraphicsSubjectIdentity struct {
	Type string
	Name string
}

type DomainGraphicsEvent struct {
	Phase      int
	Local      DomainEventGraphicsAddress
	Remote     DomainEventGraphicsAddress
	AuthScheme string
	Subject    []DomainEventGraphicsSubjectIdentity
}

type DomainIOErrorReasonEvent struct {
	DomainIOErrorEvent
	Reason string
}

type DomainBlockJobEvent struct {
	Disk   string
	Type   int
	Status int
}

type DomainDiskChangeEvent struct {
	OldSrcPath string
	NewSrcPath string
	DevAlias   string
	Reason     int
}

type DomainTrayChangeEvent struct {
	DevAlias string
	Reason   int
}

type DomainReasonEvent struct {
	Reason int
}

type DomainBalloonChangeEvent struct {
	Actual uint64
}

type DomainDeviceRemovedEvent struct {
	DevAlias string
}

//export domainEventLifecycleCallback
func domainEventLifecycleCallback(c C.virConnectPtr, d C.virDomainPtr,
	event int, detail int,
	opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainLifecycleEvent{
		Event:  event,
		Detail: detail,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventRTCChangeCallback
func domainEventRTCChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	utcoffset int64, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainRTCChangeEvent{
		Utcoffset: utcoffset,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventWatchdogCallback
func domainEventWatchdogCallback(c C.virConnectPtr, d C.virDomainPtr,
	action int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainWatchdogEvent{
		Action: action,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventIOErrorCallback
func domainEventIOErrorCallback(c C.virConnectPtr, d C.virDomainPtr,
	srcPath *C.char, devAlias *C.char, action int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainIOErrorEvent{
		SrcPath:  C.GoString(srcPath),
		DevAlias: C.GoString(devAlias),
		Action:   action,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventGraphicsCallback
func domainEventGraphicsCallback(c C.virConnectPtr, d C.virDomainPtr,
	phase int,
	local C.virDomainEventGraphicsAddressPtr,
	remote C.virDomainEventGraphicsAddressPtr,
	authScheme *C.char,
	subject C.virDomainEventGraphicsSubjectPtr,
	opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	subjectGo := make([]DomainEventGraphicsSubjectIdentity, subject.nidentity)
	nidentities := int(subject.nidentity)
	identities := (*[1 << 30]C.virDomainEventGraphicsSubjectIdentity)(unsafe.Pointer(&subject.identities))[:nidentities:nidentities]
	for _, identity := range identities {
		subjectGo = append(subjectGo,
			DomainEventGraphicsSubjectIdentity{
				Type: C.GoString(identity._type),
				Name: C.GoString(identity.name),
			},
		)
	}

	eventDetails := DomainGraphicsEvent{
		Phase: phase,
		Local: DomainEventGraphicsAddress{
			Family:  int(local.family),
			Node:    C.GoString(local.node),
			Service: C.GoString(local.service),
		},
		Remote: DomainEventGraphicsAddress{
			Family:  int(remote.family),
			Node:    C.GoString(remote.node),
			Service: C.GoString(remote.service),
		},
		AuthScheme: C.GoString(authScheme),
		Subject:    subjectGo,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventIOErrorReasonCallback
func domainEventIOErrorReasonCallback(c C.virConnectPtr, d C.virDomainPtr,
	srcPath *C.char, devAlias *C.char, action int, reason *C.char,
	opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainIOErrorReasonEvent{
		DomainIOErrorEvent: DomainIOErrorEvent{
			SrcPath:  C.GoString(srcPath),
			DevAlias: C.GoString(devAlias),
			Action:   action,
		},
		Reason: C.GoString(reason),
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventBlockJobCallback
func domainEventBlockJobCallback(c C.virConnectPtr, d C.virDomainPtr,
	disk *C.char, _type int, status int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainBlockJobEvent{
		Disk:   C.GoString(disk),
		Type:   _type,
		Status: status,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventDiskChangeCallback
func domainEventDiskChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	oldSrcPath *C.char, newSrcPath *C.char, devAlias *C.char,
	reason int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainDiskChangeEvent{
		OldSrcPath: C.GoString(oldSrcPath),
		NewSrcPath: C.GoString(newSrcPath),
		DevAlias:   C.GoString(devAlias),
		Reason:     reason,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventTrayChangeCallback
func domainEventTrayChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	devAlias *C.char, reason int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainTrayChangeEvent{
		DevAlias: C.GoString(devAlias),
		Reason:   reason,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventReasonCallback
func domainEventReasonCallback(c C.virConnectPtr, d C.virDomainPtr,
	reason int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainReasonEvent{
		Reason: reason,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventBalloonChangeCallback
func domainEventBalloonChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	actual uint64, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainBalloonChangeEvent{
		Actual: actual,
	}

	return callCallback(opaque, &connection, &domain, eventDetails)
}

//export domainEventDeviceRemovedCallback
func domainEventDeviceRemovedCallback(c C.virConnectPtr, d C.virDomainPtr,
	devAlias *C.char, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainDeviceRemovedEvent{
		DevAlias: C.GoString(devAlias),
	}
	return callCallback(opaque, &connection, &domain, eventDetails)
}

type DomainEventCallback func(c *VirConnection, d *VirDomain,
	event interface{}, f func()) int

type domainCallbackContext struct {
	cb *DomainEventCallback
	f  func()
}

var goCallbackLock sync.RWMutex
var goCallbacks = make(map[int]*domainCallbackContext)
var nextGoCallbackId int

type storageDomainCallbackContext struct {
	sync.RWMutex
	Data  map[int]*domainCallbackContext
	Index int
}

var stDomCbCtx = storageDomainCallbackContext{
	Data:  make(map[int]*domainCallbackContext),
	Index: 0,
}

func (s *storageDomainCallbackContext) Find(key int) *domainCallbackContext {
	s.RLock()
	defer s.RUnlock()

	return s.Data[key]
}

func (s *storageDomainCallbackContext) Erase(key int) {
	s.Lock()
	defer s.Unlock()

	delete(s.Data, key)

	return
}

func (s *storageDomainCallbackContext) Insert(value *domainCallbackContext) int {
	s.Lock()
	defer s.Unlock()

	s.Index++
	s.Data[s.Index] = value

	return s.Index
}

func (s *storageDomainCallbackContext) Size() int {

	return len(s.Data)
}

//export freeCallbackId
func freeCallbackId(ID int) {
	stDomCbCtx.Erase(ID)

	return
}

func callCallback(ID int, c *VirConnection, d *VirDomain, event interface{}) int {
	ctx := stDomCbCtx.Find(ID)
	if ctx == nil {
		// If this happens there must be a bug in libvirt
		panic("Callback arrived after freeCallbackId was called")
	}

	return (*ctx.cb)(c, d, event, ctx.f)
}

func (c *VirConnection) DomainEventRegister(dom VirDomain, evID int, callback *DomainEventCallback, opaque func()) (cbID int, err error) {

	var cbPtr unsafe.Pointer

	ID := stDomCbCtx.Insert(&domainCallbackContext{
		cb: callback,
		f:  opaque,
	})

	switch evID {
	case VIR_DOMAIN_EVENT_ID_LIFECYCLE:
		cbPtr = unsafe.Pointer(C.domainEventLifecycleCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_LIFECYCLE
	case VIR_DOMAIN_EVENT_ID_RTC_CHANGE:
		cbPtr = unsafe.Pointer(C.domainEventRTCChangeCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_RTC_CHANGE
	case VIR_DOMAIN_EVENT_ID_WATCHDOG:
		cbPtr = unsafe.Pointer(C.domainEventWatchdogCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_WATCHDOG
	case VIR_DOMAIN_EVENT_ID_IO_ERROR:
		cbPtr = unsafe.Pointer(C.domainEventIOErrorCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_IO_ERROR
	case VIR_DOMAIN_EVENT_ID_GRAPHICS:
		cbPtr = unsafe.Pointer(C.domainEventGraphicsCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_GRAPHICS
	case VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON:
		cbPtr = unsafe.Pointer(C.domainEventIOErrorReasonCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON
	case VIR_DOMAIN_EVENT_ID_BLOCK_JOB:
		cbPtr = unsafe.Pointer(C.domainEventBlockJobCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_BLOCK_JOB
	case VIR_DOMAIN_EVENT_ID_DISK_CHANGE:
		cbPtr = unsafe.Pointer(C.domainEventDiskChangeCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_DISK_CHANGE
	case VIR_DOMAIN_EVENT_ID_TRAY_CHANGE:
		cbPtr = unsafe.Pointer(C.domainEventTrayChangeCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_TRAY_CHANGE
	case VIR_DOMAIN_EVENT_ID_BALLOON_CHANGE:
		cbPtr = unsafe.Pointer(C.domainEventBalloonChangeCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_BALLOON_CHANGE
	case VIR_DOMAIN_EVENT_ID_DEVICE_REMOVED:
		cbPtr = unsafe.Pointer(C.domainEventDeviceRemovedCallback_cgo)
		evID = C.VIR_DOMAIN_EVENT_ID_DEVICE_REMOVED
	default:
		cbPtr = nil
		evID = -1
	}

	if cbID = int(C.virConnectDomainEventRegisterAny_cgo(c.ptr, dom.ptr, C.int(evID), C.virConnectDomainEventGenericCallback(cbPtr), C.int(ID))); cbID == -1 {
		stDomCbCtx.Erase(ID)
		err = GetLastError()
	}

	return
}

/*
func (c *VirConnection) DomainEventDeregister(callbackId int) int {
	// Deregister the callback
	return int(C.virConnectDomainEventDeregisterAny(c.ptr, C.int(callbackId)))
}

func EventRegisterDefaultImpl() int {
	return int(C.virEventRegisterDefaultImpl())
}

func EventRunDefaultImpl() int {
	return int(C.virEventRunDefaultImpl())
}
*/
func (c *VirConnection) DomainEventDeregister(callbackId int) (err error) {

	if rc := int(C.virConnectDomainEventDeregisterAny(c.ptr, C.int(callbackId))); rc == -1 {
		err = GetLastError()
	}

	return
}

func EventRegisterDefaultImpl() (err error) {
	if rc := int(C.virEventRegisterDefaultImpl()); rc == -1 {
		err = GetLastError()
	}

	return
}

func EventRunDefaultImpl() (err error) {
	if rc := int(C.virEventRunDefaultImpl()); rc == -1 {
		err = GetLastError()
	}

	return
}

func (e *DomainLifecycleEvent) String() string {

	var detail, event string
	switch e.Event {

	case VIR_DOMAIN_EVENT_DEFINED:
		event = "defined"
		switch e.Detail {

		case VIR_DOMAIN_EVENT_DEFINED_ADDED:
			detail = "added"

		case VIR_DOMAIN_EVENT_DEFINED_UPDATED:
			detail = "updated"

		default:
			detail = "unknown detail description"
		}

	case VIR_DOMAIN_EVENT_UNDEFINED:
		event = "undefined"
		switch e.Detail {

		case VIR_DOMAIN_EVENT_UNDEFINED_REMOVED:
			detail = "removed"

		default:
			detail = "unknown detail description"
		}

	case VIR_DOMAIN_EVENT_STARTED:
		event = "started"
		switch e.Detail {

		case VIR_DOMAIN_EVENT_STARTED_BOOTED:
			detail = "booted"

		case VIR_DOMAIN_EVENT_STARTED_MIGRATED:
			detail = "migrated"

		case VIR_DOMAIN_EVENT_STARTED_RESTORED:
			detail = "restored"

		case VIR_DOMAIN_EVENT_STARTED_FROM_SNAPSHOT:
			detail = "snapshot"

		default:
			detail = "unknown detail description"
		}

	case VIR_DOMAIN_EVENT_SUSPENDED:
		event = "suspended"
		switch e.Detail {

		case VIR_DOMAIN_EVENT_SUSPENDED_PAUSED:
			detail = "paused"

		case VIR_DOMAIN_EVENT_SUSPENDED_MIGRATED:
			detail = "migrated"

		case VIR_DOMAIN_EVENT_SUSPENDED_IOERROR:
			detail = "i/o error"

		case VIR_DOMAIN_EVENT_SUSPENDED_WATCHDOG:
			detail = "watchdog"

		case VIR_DOMAIN_EVENT_SUSPENDED_RESTORED:
			detail = "restored"

		case VIR_DOMAIN_EVENT_SUSPENDED_FROM_SNAPSHOT:
			detail = "snapshot"

		default:
			detail = "unknown detail description"
		}

	case VIR_DOMAIN_EVENT_RESUMED:
		event = "resumed"
		switch e.Detail {

		case VIR_DOMAIN_EVENT_RESUMED_UNPAUSED:
			detail = "unpaused"

		case VIR_DOMAIN_EVENT_RESUMED_MIGRATED:
			detail = "migrated"

		case VIR_DOMAIN_EVENT_RESUMED_FROM_SNAPSHOT:
			detail = "snapshot"

		default:
			detail = "unknown detail description"
		}

	case VIR_DOMAIN_EVENT_STOPPED:
		event = "stopped"
		switch e.Detail {

		case VIR_DOMAIN_EVENT_STOPPED_SHUTDOWN:
			detail = "shutdown"

		case VIR_DOMAIN_EVENT_STOPPED_DESTROYED:
			detail = "destroyed"

		case VIR_DOMAIN_EVENT_STOPPED_CRASHED:
			detail = "crashed"

		case VIR_DOMAIN_EVENT_STOPPED_MIGRATED:
			detail = "migrated"

		case VIR_DOMAIN_EVENT_STOPPED_SAVED:
			detail = "saved"

		case VIR_DOMAIN_EVENT_STOPPED_FAILED:
			detail = "failed"

		case VIR_DOMAIN_EVENT_STOPPED_FROM_SNAPSHOT:
			detail = "snapshot"

		default:
			detail = "unknown detail description"
		}

	case VIR_DOMAIN_EVENT_SHUTDOWN:
		event = "shutdown"
		switch e.Detail {

		case VIR_DOMAIN_EVENT_SHUTDOWN_FINISHED:
			detail = "finished"

		default:
			detail = "unknown detail description"
		}

	default:
		event = "unknown event description"
	}

	return fmt.Sprintf("Domain event='%s' detail='%s'", event, detail)
}

func (e *DomainRTCChangeEvent) String() string {

	return fmt.Sprintf("RTC change %d", e.Utcoffset)
}

func (e *DomainWatchdogEvent) String() string {

	return fmt.Sprintf("Watchdog action='%d'", e.Action)
}

func (e *DomainIOErrorEvent) String() string {

	return fmt.Sprintf("IO error path='%s' alias='%s' action='%d'", e.SrcPath, e.DevAlias, e.Action)
}

func (e *DomainGraphicsEvent) String() string {

	var phase string
	switch e.Phase {

	case VIR_DOMAIN_EVENT_GRAPHICS_CONNECT:
		phase = "connected"

	case VIR_DOMAIN_EVENT_GRAPHICS_INITIALIZE:
		phase = "initialized"

	case VIR_DOMAIN_EVENT_GRAPHICS_DISCONNECT:
		phase = "disconnected"
	}

	return fmt.Sprintf("Graphics phase='%s'", phase)
}

func (e *DomainIOErrorReasonEvent) String() string {

	return fmt.Sprintf("IO error path='%s' alias='%s' action='%d' reason='%s'", e.SrcPath, e.DevAlias, e.Action, e.Reason)
}

func (e *DomainBlockJobEvent) String() string {

	return fmt.Sprintf("Block job disk='%s' status='%d' type='%d'", e.Disk, e.Status, e.Type)
}

func (e *DomainDiskChangeEvent) String() string {

	return fmt.Sprintf("Disk change old='%s' new='%s' alias='%s' reason='%d'", e.OldSrcPath, e.NewSrcPath, e.DevAlias, e.Reason)
}

func (e *DomainTrayChangeEvent) String() string {

	return fmt.Sprintf("Tray change dev='%s' reason='%d'", e.DevAlias, e.Reason)
}

func (e *DomainBalloonChangeEvent) String() string {

	return fmt.Sprintf("Ballon change '%d'", e.Actual)
}

func (e *DomainDeviceRemovedEvent) String() string {

	return fmt.Sprintf("Device '%s' removed ", e.DevAlias)
}
