package libvirt

import (
	"testing"
	"time"
)

func init() {
	EventRegisterDefaultImpl()
}

func TestDomainEventRegister(t *testing.T) {

	callbackId := -1

	conn := buildTestConnection()
	defer func() {
		if callbackId >= 0 {
			ret := conn.DomainEventDeregister(callbackId)
			if ret != 0 {
				t.Errorf("got %d on DomainEventDeregister instead of 0", ret)
			}
		}
		conn.CloseConnection()
	}()

	nodom := VirDomain{}
	defName := time.Now().String()

	nbEvents := 0

	callback := DomainEventCallback(
		func(c *VirConnection, d *VirDomain, eventDetails interface{}, f func()) int {
			if lifecycleEvent, ok := eventDetails.(DomainLifecycleEvent); ok {
				if lifecycleEvent.Event == VIR_DOMAIN_EVENT_STARTED {
					domName, _ := d.GetName()
					if defName != domName {
						t.Fatalf("Name was not '%s': %s", defName, domName)
					}
				}
			} else {
				t.Fatalf("event details isn't DomainLifecycleEvent")
			}
			f()
			return 0
		},
	)

	callbackId = conn.DomainEventRegister(
		VirDomain{},
		VIR_DOMAIN_EVENT_ID_LIFECYCLE,
		&callback,
		func() {
			nbEvents++
		},
	)

	// Test a minimally valid xml
	xml := `<domain type="test">
		<name>` + defName + `</name>
		<memory unit="KiB">8192</memory>
		<os>
			<type>hvm</type>
		</os>
	</domain>`
	dom, err := conn.DomainCreateXML(xml, VIR_DOMAIN_NONE)
	if err != nil {
		t.Error(err)
		return
	}

	// This is blocking as long as there is no message
	EventRunDefaultImpl()
	if nbEvents == 0 {
		t.Fatal("At least one event was expected")
	}

	defer func() {
		if dom != nodom {
			dom.Destroy()
			dom.Free()
		}
	}()

	// Check that the internal context entry was added, and that there only is
	// one.
	goCallbackLock.Lock()
	if len(goCallbacks) != 1 {
		t.Error("goCallbacks should hold one entry")
	}
	goCallbackLock.Unlock()

	// Deregister the event
	if ret := conn.DomainEventDeregister(callbackId); ret < 0 {
		t.Fatal("Event deregistration failed")
	}
	callbackId = -1 // Don't deregister twice

	// Check that the internal context entries was removed
	goCallbackLock.Lock()
	if len(goCallbacks) > 0 {
		t.Error("goCallbacks entry wasn't removed")
	}
	goCallbackLock.Unlock()
}

func TestEventAddTimeout(t *testing.T) {

	callbackId := -1
	defer func() {
		if callbackId >= 0 {
			ret := EventRemoveTimeout(callbackId)
			if ret != 0 {
				t.Errorf("got %d on EventRemoveTimeout instead of 0", ret)
			}
			callbackId = -1
		}
	}()

	nbEvents := 0

	callback := EventCallback(
		func(eventDetails interface{}, f func()) {
			if timeoutEvent, ok := eventDetails.(TimeoutEvent); ok {
				if timeoutEvent.TimerId != callbackId {
					t.Errorf("Timer ID == %d, expected %d", timeoutEvent.TimerId, callbackId)
				} else {
					ret := EventRemoveTimeout(callbackId)
					if ret != 0 {
						t.Errorf("got %d on EventRemoveTimeout instead of 0", ret)
					}
					callbackId = -1
				}
			} else {
				t.Errorf("event details isn't TimeoutEvent")
			}
			f()
		},
	)

	callbackId = EventAddTimeout(10, // milliseconds
		&callback,
		func() {
			nbEvents++
		},
	)

	// This is blocking as long as there is no message
	done := make(chan struct{})
	timeout := time.After(100 * time.Millisecond)
	go func() {
		EventRunDefaultImpl()
		close(done)
		EventRunDefaultImpl()
	}()
	select {
	case <-done: // OK!
	case <-timeout:
		t.Fatalf("timeout reached while waiting timeout event")
		return
	}
	time.Sleep(100 * time.Millisecond)
	if nbEvents != 1 {
		t.Errorf("Exactly one event expected, got %d", nbEvents)
	}
}
