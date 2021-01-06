package pkg

import (
	"testing"
	"time"
)

func TestNewSlideWindowCounter(t *testing.T) {
	counter := NewSlideWindowCounter(60, 60*time.Second)

	go counter.run()

	go func() {
		for {
			counter.Increment()
		}
	}()

	go func() {
		for {
			time.Sleep(2000 * time.Millisecond)
			t.Logf("avg=%+v", counter.Avg())
		}
	}()

	select {}
}
