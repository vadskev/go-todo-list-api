package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	d, err := time.Parse("20060102", date)
	if err != nil {
		return "", err
	}

	switch {
	case strings.HasPrefix(repeat, "d "):
		n, err := strconv.Atoi(repeat[2:])
		if err != nil {
			return "", err
		}

		if n > 365 {
			return "", fmt.Errorf("interval exceed the maximum allowed value")
		}

		d = d.AddDate(0, 0, n)
	case strings.HasPrefix(repeat, "y"):
		d = d.AddDate(1, 0, 0)
	case strings.HasPrefix(repeat, "w "):
		return "", fmt.Errorf("unknown repeat format: %s", repeat)
	case strings.HasPrefix(repeat, "m "):
		return "", fmt.Errorf("unknown repeat format: %s", repeat)
	default:
		return "", fmt.Errorf("unknown repeat format: %s", repeat)
	}

	if d.Before(now) {
		return NextDate(now, d.Format("20060102"), repeat)
	}

	return d.Format("20060102"), nil
}
