package time_tutorial

import (
	"fmt"
	"testing"
	"time"
)

func Test_time(t *testing.T) {

}

func Test_func_After(t *testing.T) {
	var c chan int
	d := 10 * time.Second
	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(d):
		fmt.Println("timed out")
	}

	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.NewTimer(d).C:
		fmt.Println("timed out")
	}
}

func Test_func_Sleep(t *testing.T) {
	time.Sleep(time.Minute * -1)
	fmt.Println("not sleep")
	time.Sleep(time.Minute)
	fmt.Println("sleep a minute")
}

func Test_func_Tick(t *testing.T) {
	c := time.Tick(-5 * time.Second)
	if c == nil {
		fmt.Println("negative")
	}
	c = time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, "p!")
	}
}

func Test_type_Duration(t *testing.T) {
	d, _ := time.ParseDuration("100ms")
	fmt.Println(d)

	// Since
	// Until
	fmt.Println(time.Since(time.Now().Add(-time.Hour)))
	fmt.Println(time.Now().Sub(time.Now().Add(-time.Hour)))

	fmt.Println(time.Until(time.Now().Add(-time.Hour)))
	fmt.Println(time.Now().Add(-time.Hour).Sub(time.Now()))

	fmt.Println(time.Until(time.Now().Add(-time.Hour)).Abs())

	fmt.Println(d.Hours())

	fmt.Println(d.Milliseconds())

	fmt.Println(d.Microseconds())

	fmt.Println(d.Minutes())

	fmt.Println(d.Nanoseconds())

	d2, _ := time.ParseDuration("21ms")
	fmt.Println(d.Round(d2))

	fmt.Println(d.String())

	d, _ = time.ParseDuration("10.5s")

	d2, _ = time.ParseDuration("1h25m10s760ms")
	fmt.Println(d2.Truncate(d))
}

func Test_type_Location(t *testing.T) {
	fmt.Println(time.Local)
}
