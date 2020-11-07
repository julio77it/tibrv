package tibrv

import (
	"os"
	"testing"
	"time"
)

func TestFtMember(t *testing.T) {
	groupName := "TEST.GRP"
	subject := "UNIT.TEST.FT"

	var queue RvQueue
	if err := queue.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer queue.Destroy()

	var transport RvNetTransport
	err := transport.Create(
		Service(os.Getenv("TEST_SERVICE")),
		Network(os.Getenv("TEST_NETWORK")),
		Daemon(os.Getenv("TEST_DAEMON")),
	)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer transport.Destroy()

	var msg RvMessage
	if err := msg.Create(); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	defer msg.Destroy()

	if err := msg.SetInt32("Integer32bit", -25); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if err := msg.SetSendSubject(subject); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	var output string
	var listener RvListener

	var callback FtCallback = func(s *string, q RvQueue, t RvNetTransport, l *RvListener, test *testing.T) func(groupName string, ftAction uint) {
		return func(groupName string, ftAction uint) {
			*s = msg.String()

			if ftAction == FtActivate {
				var callback RvCallback = func(s *string) func(msg *RvMessage) {
					return func(msg *RvMessage) {
						*s = msg.String()
					}
				}(s)

				err := l.Create(
					q,
					callback,
					t,
					subject,
				)
				if err != nil {
					test.Fatalf("Expected nil, got %v", err)
				}
			}
		}
	}(&output, queue, transport, &listener, t)

	var member FtMember
	err = member.Create(
		queue,
		callback,
		transport,
		groupName,
		2, 2, 1, 0, 3,
	)
	time.Sleep(time.Second * 2)

	if err := transport.Send(msg); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	for i := 0; i < 5; i++ {
		queue.Dispatch()
	}
	input := msg.String()

	if output != input {
		t.Fatalf("Expected %s, got %s", input, output)
	}
}
