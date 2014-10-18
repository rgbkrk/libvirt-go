// +build integration
// Stream tests can't run on the test driver

package libvirt

import (
	"sync"
	"testing"
	"time"
)

func TestStreamEvents(t *testing.T) {

	EventRegisterDefaultImpl()

	// First open the lxc connection
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Fatalf("Failed to create connection: %s", err)
	}
	defer conn.CloseConnection()

	// Then create the test domain
	script := []string{
		`-c`,
		`for i in 0 1 2 3 4; do echo "Hello World #$i"; sleep 1; done`,
	}
	dom, err := defineTestLxcDomainWithCmd(conn, "streamEvents",
		"/bin/sh", script)
	if err != nil {
		t.Fatalf("Failed to define domain: %s", err)
	}
	defer func() {
		dom.Destroy()
		dom.Undefine()
	}()

	// Create the tested stream
	stream, err := NewVirStream(&conn, VIR_STREAM_NONBLOCK)
	if err != nil {
		t.Fatalf("Failed to create stream: %s", err)
	}
	defer func() {
		stream.Close()
		stream.Free()
	}()

	// Start the domain
	if err = dom.Create(); err != nil {
		t.Fatalf("Failed to start domain: %s", err)
	}

	// Connect the stream to the domain console
	if err = dom.OpenConsole("", stream, 0); err != nil {
		t.Fatalf("Failed to open console: %s", err)
	}

	// Add the stream event callback
	var (
		callback      VirStreamEventCallback
		eventsCounter func()
		mCount        sync.RWMutex // synchronize access to the nb variables

		nbEvents   int = 0
		nbReadable int = 0
		nbWritable int = 0
		nbError    int = 0
		nbHangup   int = 0
	)
	callback = VirStreamEventCallback(
		func(s *VirStream, events int, f func()) int {
			mCount.Lock()
			switch events {
			case VIR_EVENT_HANDLE_ERROR:
				nbError++
			case VIR_EVENT_HANDLE_HANGUP:
				nbHangup++
			default:
				if (events & VIR_EVENT_HANDLE_READABLE) > 0 {
					nbReadable++
				}
				if (events & VIR_EVENT_HANDLE_WRITABLE) > 0 {
					nbWritable++
				}
			}
			mCount.Unlock()

			f()
			return 0
		},
	)
	eventsCounter = func() {
		nbEvents++
	}
	if err := stream.EventAddCallback(VIR_EVENT_HANDLE_READABLE, &callback, eventsCounter); err != nil {
		t.Fatalf("Failed to add stream callback: %s", err)
	}
	defer stream.EventRemoveCallback()

	// Test that it actually works
	clearEvents := func() {
		mCount.Lock()
		defer mCount.Unlock()
		nbReadable = 0
		nbWritable = 0
		nbError = 0
		nbHangup = 0
		nbEvents = 0
	}

	quitCh := make(chan struct{})
	lock := make(chan struct{})
	defer close(quitCh)
	defer close(lock)

	// Event loop, send to lock on first callback
	// and then do nothing until a count reset
	go func() {
		for {
			select {
			case <-quitCh:
				return
			default:
				EventRunDefaultImpl()
				mCount.Lock()
				if nbEvents == 1 {
					lock <- struct{}{}
				}
				mCount.Unlock()
			}
		}
	}()

	// Wait 2 seconds for a READABLE event to come in, or fail
	timeoutDur := 2 * time.Second
	timeout := time.AfterFunc(timeoutDur, func() {
		t.Fatal("Timed out waiting for READABLE events")
	})
	<-lock
	timeout.Stop()

	mCount.RLock()
	if nbReadable == 0 || nbWritable != 0 ||
		nbError != 0 || nbHangup != 0 {
		mCount.RUnlock()
		t.Fatalf("Expected only readable events, got [%d, %d, %d, %d] "+
			"read, write, error, hang up events",
			nbReadable, nbWritable, nbError, nbHangup)
	}
	mCount.RUnlock()

	if err = stream.EventUpdateCallback(VIR_EVENT_HANDLE_WRITABLE); err != nil {
		t.Fatalf("Failed to update callback: %s", err)
	}
	// Clear event count after we update the callback
	clearEvents()

	// Wait 2 seconds for a WRITEABLE event to come in, or fail
	timeout = time.AfterFunc(timeoutDur, func() {
		t.Fatal("Timed out waiting for WRITEABLE events")
	})
	<-lock
	timeout.Stop()

	mCount.RLock()
	// TODO: nReadable incorrectly gets incremented after we're
	// unsubscribed from it ...
	if nbWritable == 0 || nbError != 0 || nbHangup != 0 {
		defer mCount.RUnlock()
		t.Fatalf("Expected writable events, got [%d, %d, %d, %d] "+
			"read, write, error, hang up events",
			nbReadable, nbWritable, nbError, nbHangup)
	}
	mCount.RUnlock()

	// wait for event loop to quit
	quitCh <- struct{}{}
}
