package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

// we want a spy that does two things:
//   - spies on the Sleeper interface
//   - spies on the Writer interface
//
// This way we can spy on when something was written, and when
// something was slept, and ensure that the order of execution is as
// we expect
type SpyCountdownOperations struct {
	Calls []string
}

// implement Sleeper
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// implement Writer
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)

	return
}

func TestCountdown(t *testing.T) {
	t.Run("generates correct output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyCountdownOperations{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("runs sleep before each print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		Countdown(spySleepPrinter, spySleepPrinter)

		got := spySleepPrinter.Calls
		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := &ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.Sleep,
	}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf(
			"should have slept for %v but slept for %v",
			sleepTime,
			spyTime.durationSlept,
		)
	}
}
