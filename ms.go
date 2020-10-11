// Package ms easily convert various time formats to milliseconds.
package ms

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	s = 1000
	m = s * 60
	h = m * 60
	d = h * 24
	w = d * 7
	y = d * 365.25
)

var (
	// ErrTooLong input lenght is over than 100 chars
	ErrTooLong = errors.New("input is over than 100 charactors")
	// ErrInvalidFormat invalid format
	ErrInvalidFormat = errors.New("invalid ms format")
	// ErrNaN ms is NaN
	ErrNaN = errors.New("ms is NaN")
	// ErrInfinity ms is infinity
	ErrInfinity = errors.New("ms is NaN")
)

var regex *regexp.Regexp

func init() {
	regex, _ = regexp.Compile(`(?i)^(-?(?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|weeks?|w|years?|yrs?|y)?$`)
}

// Parse parse the given str and return milliseconds.
func Parse(str string) (float64, error) {
	if len(str) > 100 {
		return 0, ErrTooLong
	}
	match := regex.FindStringSubmatch(str)
	if match == nil {
		return 0, ErrInvalidFormat
	}
	n, _ := strconv.ParseFloat(match[1], 10)
	switch strings.ToLower(match[2]) {
	case "years", "year", "yrs", "yr", "y":
		return n * y, nil
	case "weeks", "week", "w":
		return n * w, nil
	case "days", "day", "d":
		return n * d, nil
	case "hours", "hour", "hrs", "hr", "h":
		return n * h, nil
	case "minutes", "minute", "mins", "min", "m":
		return n * m, nil
	case "seconds", "second", "secs", "sec", "s":
		return n * s, nil
	case "milliseconds", "millisecond", "msecs", "msec", "ms", "":
		return n, nil
	default:
		panic("unreachable")
	}
}

// MustParse is like Parse but panics if error encoutered.
func MustParse(str string) float64 {
	ms, err := Parse(str)
	if err != nil {
		panic(err)
	}
	return ms
}

// FormatShort short format for ms.
func FormatShort(ms float64) (string, error) {
	if math.IsNaN(ms) {
		return "", ErrNaN
	}
	if math.IsInf(ms, 0) {
		return "", ErrInfinity
	}
	msAbs := math.Abs(ms)
	if msAbs >= d {
		return fmt.Sprintf("%.0fd", ms/d), nil
	}
	if msAbs >= h {
		return fmt.Sprintf("%.0fh", ms/h), nil
	}
	if msAbs >= m {
		return fmt.Sprintf("%.0fm", ms/m), nil
	}
	if msAbs >= s {
		return fmt.Sprintf("%.0fs", ms/s), nil
	}
	return fmt.Sprintf("%sms", strconv.FormatFloat(ms, 'f', -1, 64)), nil
}

// FormatLong long format for ms.
func FormatLong(ms float64) (string, error) {
	if math.IsNaN(ms) {
		return "", ErrNaN
	}
	if math.IsInf(ms, 0) {
		return "", ErrInfinity
	}
	msAbs := math.Abs(ms)
	if msAbs >= d {
		return plural(ms, msAbs, d, "day"), nil
	}
	if msAbs >= h {
		return plural(ms, msAbs, h, "hour"), nil
	}
	if msAbs >= m {
		return plural(ms, msAbs, m, "minute"), nil
	}
	if msAbs >= s {
		return plural(ms, msAbs, s, "second"), nil
	}
	return fmt.Sprintf("%s ms", strconv.FormatFloat(ms, 'f', -1, 64)), nil
}

func plural(ms, msAbs, n float64, name string) string {
	isPlural := msAbs >= n*1.5
	if isPlural {
		return fmt.Sprintf("%.0f %ss", ms/n, name)
	}
	return fmt.Sprintf("%.0f %s", ms/n, name)
}
