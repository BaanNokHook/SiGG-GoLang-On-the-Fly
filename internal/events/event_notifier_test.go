// SiGG-GoLang-On-the-Fly //

package events

import (
	"context"
	"testing"
)

func TestEventNotifier(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	en := newEventNotifier(ctx, "ut")
	var mySeq int64 = 1000000
	events := make(chan bool)
	go func() {
		defer close(events)
		for {
			err := en.waitNext(mySeq)
			if err != nil {
				return
			}
			events <- true
			mySeq++
		}
	}()

	en.newEvents <- 1000001
	<-events
	en.newEvents <- 1000002
	<-events

	cancel()
	<-events
}

func TestEventNotifierClosedChannel(t *testing.T) {
	en := newEventNotifier(context.Background(), "ut")
	var mySeq int64 = 1000000
	events := make(chan bool)
	go func() {
		defer close(events)
		for {
			err := en.waitNext(mySeq)
			if err != nil {
				return
			}
			events <- true
			mySeq++
		}
	}()

	en.newEvents <- 1000001
	<-events
	en.newEvents <- 1000002
	<-events

	close(en.newEvents)
	<-events
}
