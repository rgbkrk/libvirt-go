// +build integration
// Stream tests can't run on the test driver

package libvirt

import (
	"testing"
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
    defer func () {
        dom.Destroy()
        dom.Undefine()
    } ()

	// Create the tested stream
	stream, err := NewVirStream(&conn, VIR_STREAM_NONBLOCK)
	if err != nil {
		t.Fatalf("Failed to create stream: %s", err)
	}
	defer func() {
		stream.Close()
		stream.Free()
	} ()

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

		nbEvents   int = 0
		nbReadable int = 0
		nbWritable int = 0
		nbError    int = 0
		nbHangup   int = 0
	)
	callback = VirStreamEventCallback(
		func(s *VirStream, events int, f func()) int {
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
		nbReadable = 0
		nbWritable = 0
		nbError = 0
		nbHangup = 0
		nbEvents = 0
	}

	quit := false
	lock := make(chan int)
	go func() {
		for !quit {
			EventRunDefaultImpl()
			if nbEvents == 1 {
				lock <- 1
			}
		}
	}()

	<-lock
	if nbReadable == 0 || nbWritable != 0 ||
		nbError != 0 || nbHangup != 0 {
		t.Fatalf("Expected only readable events, got [%d, %d, %d, %d] "+
			"read, write, error, hang up events",
			nbReadable, nbWritable, nbError, nbHangup)
	}
	clearEvents()

	if err = stream.EventUpdateCallback(VIR_EVENT_HANDLE_WRITABLE); err != nil {
		t.Fatalf("Failed to update callback: %s", err)
	}

	<-lock
	// Don't check for readable events, we may get some after the update
	if nbWritable == 0 || nbError != 0 || nbHangup != 0 {
		t.Fatalf("Expected writable events, got [%d, %d, %d, %d] "+
			"read, write, error, hang up events",
			nbReadable, nbWritable, nbError, nbHangup)
	}
	clearEvents()
	quit = true
	close(lock)

}
