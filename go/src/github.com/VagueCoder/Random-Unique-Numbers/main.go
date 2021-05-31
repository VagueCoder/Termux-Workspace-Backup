package main

import (
	"os"
	"fmt"
	"strconv"
	"math/rand"
	"strings"
	"time"

	"github.com/VagueCoder/Random-Unique-Numbers/counter"
)

func triggerFunc(in <-chan bool, out chan<- bool, c *counter.Counter) {
	if start, _ := <-in; start {
		for {
			err := c.Count()
			if err != nil {
				break
			}
			out <- true
		}
		close(out)
	}
}

func randomUniqueNumbers(in <-chan bool, out chan<- string, maxSize int) {
	exists := make(map[string]bool, 0)
	var result string
	for {
		if _, status := <-in; !status {
			break
		}

		result = func() string {
			var str string
			for {
				bytes := make([]byte, maxSize)
				for i := 0; i < maxSize; i++ {
					if i == 0 {
						bytes[i] = byte('1' + rand.Intn('9'-'0'))
					} else {
						bytes[i] = byte('0' + rand.Intn('9'-'0'+1))
					}
				}

				str = string(bytes)
				if _, found := exists[str]; !found {
					exists[str] = true
					break
				}
			}
			return str
		}()
		out <- result
	}
	close(out)
}

func main() {
	maxdigits, _ := strconv.Atoi(os.Args[1])
	countbreaker := counter.Initialize("9" + strings.Repeat("0", maxdigits-1))

	trigger0 := make(chan bool, 0)
	trigger := make(chan bool, 1000)
	receiver := make(chan string, 1000)
	rand.Seed(time.Now().UnixNano())

	go triggerFunc(trigger0, trigger, countbreaker)
	go randomUniqueNumbers(trigger, receiver, maxdigits)

	trigger0 <- true
	for {
		val, status := <-receiver
		if !status {
			break
		}
		fmt.Printf("%s\n", val)
	}
	close(trigger0)
}
