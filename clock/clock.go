package clock

import "time"

const layoutFormat = "15:04"

type Clock struct {
	s string
}

func New(h, m int) Clock {
	date := time.Date(2009, time.November, 10, h, m, 0, 0, time.UTC)
	return Clock{
		s: date.Format(layoutFormat),
	}
}

func (c Clock) String() string {
	return c.s
}

func (c Clock) Add(m int) Clock {
	parsed, _ := time.Parse(layoutFormat, c.s)
	return Clock{
		s: parsed.Add(time.Minute * time.Duration(m)).Format(layoutFormat),
	}
}

func (c Clock) Subtract(m int) Clock {
	parsed, _ := time.Parse(layoutFormat, c.s)
	return Clock{
		s: parsed.Add(time.Minute * time.Duration(-m)).Format(layoutFormat),
	}
}
