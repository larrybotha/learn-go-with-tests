package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

// an iterator
func countdownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	// imperative loop
	//for i := countdownStart; i > 0; i-- {
	//  fmt.Fprintln(out, i)
	//  sleeper.Sleep()
	//}

	for i := range countdownFrom(3) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{
		duration: 250 * time.Millisecond,
		sleep:    time.Sleep,
	}

	Countdown(os.Stdout, sleeper)
}
