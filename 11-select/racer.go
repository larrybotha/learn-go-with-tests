package racer

import (
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

func Racer(a, b string) (winner string) {
	// select allows for waiting on multiple channels
	// The first channel to send a value 'wins', and the case statement
	// is executed
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
