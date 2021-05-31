package counter

import (
	"fmt"
)

type Counter struct {
	place []uint
	flag  bool
	till  string
	unit  uint
}

func (c *Counter) balance() {
	matches := 0
	for i := range c.till {
		if c.place[i] == 10 {
			c.place[i+1]++
			c.place[i] -= 10
		}
		if i > 0 && i < len(c.till) {
			if fmt.Sprint(c.place[i]) == string(c.till[i]) {
				matches++
			}
		}
	}
	if matches == len(c.till)-1 {
		c.flag = true
	}
}

func (c *Counter) Count() error {
	if len(c.till) == 1 {
		if fmt.Sprint(c.place[0]) == string(c.till[0]) {
			return fmt.Errorf("")
		}
	}
	if c.flag && fmt.Sprint(c.place[0]) == string(c.till[0]) {
		return fmt.Errorf("")
	}
	c.place[0]++
	c.unit++
	if c.unit == 10 {
		c.balance()
		c.unit = 0
	}
	return nil
}

func Initialize(s string) *Counter {
	var reversed string
	var temp []uint
	for _, v := range s {
		reversed = string(v) + reversed
		temp = append(temp, 0)
	}
	return &Counter{
		place: temp,
		till:  reversed,
	}
}
