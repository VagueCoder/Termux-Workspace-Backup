package main

import (
	"fmt"
	"time"
)

func inWindow(from, to string) bool {
        var start, end time.Time
        start, _ = time.Parse("01-02-2006", from)
        end, _ = time.Parse("01-02-2006", to)
        now := time.Now()
        if now.After(start) && now.Before(end) {
                  return true
        }
        return false
}

func main() {
	fmt.Println(inWindow("01-20-2021", "12-31-2021"))
}

