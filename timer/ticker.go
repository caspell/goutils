package timer

import (
	"log"
	"time"
)

func ParseSince(since string) time.Duration {
	d, err := time.ParseDuration(since)
	if err != nil {
		return time.Second * 60
	}

	return d
}

func init() {

}

func Main() {

	d := ParseSince("1s")

	ticker := time.NewTicker(d)

	for {
		select {
		case <-ticker.C:
			log.Println("time ticker!!")
		}
	}

}
