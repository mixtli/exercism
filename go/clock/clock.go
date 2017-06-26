package clock

import "fmt"

const testVersion = 4

type Clock struct {
	minute, hour int
}

func New(h, m int) Clock {
	mins := minutes(h, m)
	return Clock{minute: mins % 60, hour: mins / 60}
}

func minutes(hour, minute int) int {
	mins := (hour*60 + minute) % 1440
	if mins < 0 {
		mins += 1440
	}
	return mins
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
