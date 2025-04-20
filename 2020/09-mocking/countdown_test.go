package countdown

import (
	"bytes"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("counts down", func(t *testing.T) {
		// previous way - create a buffer, passing in a pointer to the value
		//buffer := bytes.Buffer{}
		//Countdown(&buffer)

		// new way, initialise the buffer as a pointer
		buffer := &bytes.Buffer{}
		spySleeper := &CountdownOperationsSpy{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if len(spySleeper.Calls) != 4 {
			t.Errorf("incorrect number of calls to sleeper; got %d, wanted 4", len(spySleeper.Calls))
		}
	})

	t.Run("sleeps before each count", func(t *testing.T) {
		countdownSpy := &CountdownOperationsSpy{}

		Countdown(countdownSpy, countdownSpy)

		for i, v := range countdownSpy.Calls {
			if i%2 == 0 && v != sleep {
				t.Errorf("got %q for call %d, want %q", v, i, sleep)
			}

			if i%2 == 1 && v != write {
				t.Errorf("got %q for call %d, want %q", v, i, write)
			}
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	timeSpy := &TimeSpy{}
	sleeper := ConfigurableSleeper{duration: sleepTime, sleep: timeSpy.Sleep}
	sleeper.Sleep()

	if timeSpy.durationSlept != sleepTime {
		t.Errorf("got %v, want %v", timeSpy.durationSlept, sleepTime)
	}
}

const write = "write"
const sleep = "sleep"

// create a struct that will hold a list of each call
// - this is similar to jest.fn().mock.calls
type CountdownOperationsSpy struct {
	Calls []string
}

// implement Sleeper.Sleep on CountdownOperationsSpy
func (s *CountdownOperationsSpy) Sleep() {
	// append a simple string to the calls
	s.Calls = append(s.Calls, sleep)
}

// implement io.Writer.Write on CountdownOperationsSpy
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)

	// without this return, and the returned values in the signature, this method
	// would fail to implement io.Writer.Write
	return
}

// create a struct that we can use to mock calls to ConfigurableSleeper.sleep
type TimeSpy struct {
	durationSlept time.Duration
}

// assign the passed in duration to the struct so that it can be asserted on later.
// We don't want the timer to actually sleep, we just want to ensure that the
// interface for time.Sleep is implemented
func (s *TimeSpy) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
