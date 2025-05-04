package racer

import (
	"fmt"
	"net/http"
	"time"
)

func RacerNaive(a, b string) string {
	durationA := measureResponseTime(a)
	durationB := measureResponseTime(b)

	if durationA < durationB {
		return a
	}

	return b
}

const tenSecondTimeout = time.Second * 10

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	// select allows for waiting on multiple channels
	// The first channel to send a value 'wins', and the case statement
	// is executed
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)

	return time.Since(startTime)
}

func ping(url string) chan struct{} {
	// 1. always `make` channels - using var will use the zero value for
	// 		the channel, which is nil, and a nil channel cannot be sent to
	// 2. struct{} has the smallest memory footprint, smaller than the
	// 		allocation a bool would produce
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
