package server

import "time"

type TimeFunctions struct{}

func (f TimeFunctions) Now() time.Time { return time.Now() }

func (f TimeFunctions) Hour(i int) time.Duration { return time.Hour * time.Duration(i) }

func (f TimeFunctions) Minute(i int) time.Duration { return time.Minute * time.Duration(i) }

func (f TimeFunctions) Second(i int) time.Duration { return time.Second * time.Duration(i) }

func (f TimeFunctions) Millisecond(i int) time.Duration { return time.Millisecond * time.Duration(i) }

func (f TimeFunctions) Microsecond(i int) time.Duration { return time.Microsecond * time.Duration(i) }

func (f TimeFunctions) Nanosecond(i int) time.Duration { return time.Nanosecond * time.Duration(i) }

func (f TimeFunctions) Unix(sec int64, nsec int64) time.Time { return time.Unix(sec, nsec) }

func (f TimeFunctions) UnixSeconds(sec int64) time.Time { return time.Unix(sec, 0) }

func (f TimeFunctions) UnixMilli(msec int64) time.Time { return time.UnixMilli(msec) }

func (f TimeFunctions) UnixMicro(usec int64) time.Time { return time.UnixMicro(usec) }

func (f TimeFunctions) Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func (f TimeFunctions) ParseInLocation(layout, value string, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(layout, value, loc)
}

func (f TimeFunctions) Layout() string { return time.Layout }

func (f TimeFunctions) ANSIC() string { return time.ANSIC }

func (f TimeFunctions) UnixDate() string { return time.UnixDate }

func (f TimeFunctions) RubyDate() string { return time.RubyDate }

func (f TimeFunctions) RFC822() string { return time.RFC822 }

func (f TimeFunctions) RFC822Z() string { return time.RFC822Z }

func (f TimeFunctions) RFC850() string { return time.RFC850 }

func (f TimeFunctions) RFC1123() string { return time.RFC1123 }

func (f TimeFunctions) RFC1123Z() string { return time.RFC1123Z }

func (f TimeFunctions) RFC3339() string { return time.RFC3339 }

func (f TimeFunctions) RFC3339Nano() string { return time.RFC3339Nano }

func (f TimeFunctions) Kitchen() string { return time.Kitchen }

func (f TimeFunctions) Stamp() string { return time.Stamp }

func (f TimeFunctions) StampMilli() string { return time.StampMilli }

func (f TimeFunctions) StampMicro() string { return time.StampMicro }

func (f TimeFunctions) StampNano() string { return time.StampNano }

func (f TimeFunctions) ParseDuration(s string) (time.Duration, error) { return time.ParseDuration(s) }

func (f TimeFunctions) Add(c time.Time, d time.Duration) time.Time { return c.Add(d) }

func (f TimeFunctions) Sub(c time.Time, d time.Time) time.Duration { return c.Sub(d) }

func (f TimeFunctions) Format(t time.Time, layout string) string { return t.Format(layout) }

func (f TimeFunctions) Until(t time.Time) time.Duration { return time.Until(t) }

func (f TimeFunctions) Since(t time.Time) time.Duration { return time.Since(t) }
