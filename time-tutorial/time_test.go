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
	d, _ := time.ParseDuration("100.10ms")
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
	fmt.Println(time.UTC)

	loc := time.FixedZone("UTC-8", -8*60*60)
	t1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t1.Format(time.RFC822))

	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC, timeInUTC.In(location))
}

func Test_type_Month(t *testing.T) {
	fmt.Println(time.January)
	fmt.Println(time.February)
	fmt.Println(time.March)
	fmt.Println(time.April)
	fmt.Println(time.May)
	fmt.Println(time.June)
	fmt.Println(time.July)
	fmt.Println(time.August)
	fmt.Println(time.September)
	fmt.Println(time.October)
	fmt.Println(time.November)
	fmt.Println(time.December)
}

func Test_type_Ticker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

// Programs using times should typically store and pass them as values,
// not pointers. That is, time variables and struct fields should be of type time.Time, not *time.Time.
func Test_type_Time(t *testing.T) {
	t1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t1.Local())

	t2, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t2)

	loc, _ := time.LoadLocation("Europe/Berlin")

	// This will look for the name CEST in the Europe/Berlin time zone.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t3, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println("t3", t3)

	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t3, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println("t3", t3)

	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	t4 := time.Unix(unixTime.Unix(), 0).UTC()
	fmt.Println("t4", t4)

	fmt.Println("t4", time.UnixMicro(t4.UnixMicro()).UTC())
	fmt.Println("t4", time.UnixMilli(t4.UnixMilli()).UTC())

	//https://pkg.go.dev/time@go1.20#Time.Add

	t5 := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t5.AppendFormat(text, time.Kitchen)
	fmt.Println("t5", string(text))

	te, _ := t5.MarshalText()
	fmt.Println(string(te))
}

func Test_type_timer(t *testing.T) {
	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.AfterFunc(time.Second, func() {
		fmt.Println("Timer 3 expired")
	})
}
