package countdown

import (
	"fmt"
	"io"
	"os"
	"time"
)

/*
Countdown's 2nd parameter must implement this interface

Sleep() can be any function with this signature
*/
type Sleeper interface {
	Sleep()
}

/*
We can define a default sleeper that implements Sleeper.Sleep(), and pass this value
to Countdown
*/
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

/*
We can define a configurable sleeper, too. As long as this struct implements,
Sleeper.Sleep(), it can be passed as a parameter to Countdown
*/
type ConfigurableSleeper struct {
	duration time.Duration
	/*
		sleep has the same signature as time.Sleep - this allows us to override
		this function in our tests, while allowing us to use time.Sleep when called
		in the application
	*/
	sleep func(time.Duration)
}

/*
implement time.Sleep on ConfigurableSleeper

When a configurable sleeper is passed to Countdown, it will find (and expect) this
implementation. We can define ConfigurableSleeper.sleep to be anything we
want, all that Countdown requires is that Sleep is implemented, as per the
Sleeper interface
*/
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const start = 3
const finalWord = "Go!"

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := start; i > 0; i-- {
		sleeper.Sleep()
		// print i to w
		fmt.Fprintln(w, fmt.Sprint(i))
	}

	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	// use a default sleeper, which uses time.Sleep(1)
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)

	// use a configurable sleeper, where we can configure the sleep function
	configurableSleeper := &ConfigurableSleeper{duration: 5, sleep: time.Sleep}
	Countdown(os.Stdout, configurableSleeper)
}
