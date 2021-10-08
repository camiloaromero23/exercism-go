package clock

import (
	"fmt"
)

type clock struct {
	h int
	m int
}

func New(h int, m int) clock {
	c := clock{h, m}
	c.adjust()
	return c
}

func (c clock) Add(m int) clock {
	c.m += m
	c.adjust()
	return clock{c.h, c.m}
}

func (c clock) Subtract(m int) clock {
	c.m -= m
	c.adjust()
	return clock{c.h, c.m}
}

func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c *clock) adjust() {
	for c.m >= 60 {
		c.m -= 60
		c.h++
	}
	for c.m < 0 {
		c.m += 60
		c.h--
	}
	c.h = c.h % 24
	if c.h < 0 {
		c.h += 24
	}
}
