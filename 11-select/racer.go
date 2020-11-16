package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimer = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimer)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	/*
		select is similar to Promise.race in Javascript - the first channel that is closed
		or has anything sent to it "wins"

		select is specifically for waiting on multiple channels
	*/
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %q and %q", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		/*
			make the request
		*/

		http.Get(url)
		/*
			close the channel when we are done

			There's no need to send anything to the channel, closing it is enough to
			determine whether we are done or not
		*/
		close(ch)
	}()

	return ch
}

func measureResponseTime(url string) time.Duration {
	/*
		`time` works similarly to `performance` in Javascript
	*/
	start := time.Now()
	duration := time.Since(start)

	return duration
}
