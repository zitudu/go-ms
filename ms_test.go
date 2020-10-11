package ms

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

type testParseCases struct {
	input    string
	expected float64
}

func TestParseShort(t *testing.T) {
	testParse(t, []testParseCases{
		{"100", 100},
		{"1" + strings.Repeat("0", 99), 1e99},
		{"1m", 60000},
		{"1h", 3600000},
		{"2d", 172800000},
		{"3w", 1814400000},
		{"4y", 126230400000},
		{"1s", 1000},
		{"100ms", 100},
		{"1.5h", 5400000},
		{"1   s", 1000},
		{"1.5H", 5400000},
		{".5ms", 0.5},
		{"-100ms", -100},
		{"-1.5h", -5400000},
		{"-10.5h", -37800000},
		{"-.5h", -1800000},
	})
}

func TestParseLong(t *testing.T) {
	testParse(t, []testParseCases{
		{"53 milliseconds", 53},
		{"17 msecs", 17},
		{"1 sec", 1000},
		{"1 min", 60000},
		{"1 hr", 3600000},
		{"2 days", 172800000},
		{"1.5 hours", 5400000},
		{"-100 milliseconds", -100},
		{"-1.5 hours", -5400000},
		{"-.5 hr", -1800000},
	})
}

func TestParseInvalidInput(t *testing.T) {
	var tests = []string{
		"",
		"☃",
		"10-.5",
	}
	for _, test := range tests {
		_, err := Parse(test)
		if err != ErrInvalidFormat {
			t.Errorf("Parse(%s) expects ErrInvalidFormat, got: %v", test, err)
		}
	}
	if _, err := Parse(strings.Repeat("1", 101)); err != ErrTooLong {
		t.Error("Parse(101*\"1\") expects ErrTooLong")
	}
}

func testParse(t *testing.T, cases []testParseCases) {
	for _, c := range cases {
		act, err := Parse(c.input)
		if err != nil {
			t.Errorf("Parse(%s) unexcpted error: %v", c.input, err)
		} else if act != c.expected {
			t.Errorf("Parse(%s) = %f, expected %f", c.input, act, c.expected)
		} else if MustParse(c.input) != c.expected {
			t.Errorf("MustParse(%s) = %f, expected %f", c.input, act, c.expected)
		}
	}
}

type testFormatCases struct {
	input    float64
	expected string
}

func TestFormatLong(t *testing.T) {
	cases := []testFormatCases{
		{1.00, "1 ms"},
		{1.01, "1.01 ms"},
		{500, "500 ms"},
		{-500, "-500 ms"},
		{1000, "1 second"},
		{1200, "1 second"},
		{10000, "10 seconds"},
		{-1000, "-1 second"},
		{-1200, "-1 second"},
		{-10000, "-10 seconds"},
		{60 * 1000, "1 minute"},
		{60 * 1200, "1 minute"},
		{60 * 10000, "10 minutes"},
		{-60 * 1000, "-1 minute"},
		{-60 * 1200, "-1 minute"},
		{-60 * 10000, "-10 minutes"},
		{60 * 60 * 1000, "1 hour"},
		{60 * 60 * 1200, "1 hour"},
		{60 * 60 * 10000, "10 hours"},
		{-60 * 60 * 1000, "-1 hour"},
		{-60 * 60 * 1200, "-1 hour"},
		{-60 * 60 * 10000, "-10 hours"},
		{24 * 60 * 60 * 1000, "1 day"},
		{24 * 60 * 60 * 1200, "1 day"},
		{24 * 60 * 60 * 10000, "10 days"},
		{-24 * 60 * 60 * 1000, "-1 day"},
		{-24 * 60 * 60 * 1200, "-1 day"},
		{-24 * 60 * 60 * 10000, "-10 days"},
		{234234234, "3 days"},
		{-234234234, "-3 days"},
	}
	for _, c := range cases {
		act, err := FormatLong(c.input)
		if err != nil {
			t.Errorf("FormatLong(%f) unexcpted error: %v", c.input, err)
		} else if act != c.expected {
			t.Errorf("FormatLong(%f) = %s, expected %s", c.input, act, c.expected)
		}
	}
}

func TestFormatShort(t *testing.T) {
	cases := []testFormatCases{
		{1.00, "1ms"},
		{1.01, "1.01ms"},
		{500, "500ms"},
		{-500, "-500ms"},
		{1000, "1s"},
		{1200, "1s"},
		{10000, "10s"},
		{-1000, "-1s"},
		{-1200, "-1s"},
		{-10000, "-10s"},
		{60 * 1000, "1m"},
		{60 * 1200, "1m"},
		{60 * 10000, "10m"},
		{-60 * 1000, "-1m"},
		{-60 * 1200, "-1m"},
		{-60 * 10000, "-10m"},
		{60 * 60 * 1000, "1h"},
		{60 * 60 * 1200, "1h"},
		{60 * 60 * 10000, "10h"},
		{-60 * 60 * 1000, "-1h"},
		{-60 * 60 * 1200, "-1h"},
		{-60 * 60 * 10000, "-10h"},
		{24 * 60 * 60 * 1000, "1d"},
		{24 * 60 * 60 * 1200, "1d"},
		{24 * 60 * 60 * 10000, "10d"},
		{-24 * 60 * 60 * 1000, "-1d"},
		{-24 * 60 * 60 * 1200, "-1d"},
		{-24 * 60 * 60 * 10000, "-10d"},
		{234234234, "3d"},
		{-234234234, "-3d"},
	}
	for _, c := range cases {
		act, err := FormatShort(c.input)
		if err != nil {
			t.Errorf("FormatShort(%f) unexcpted error: %v", c.input, err)
		} else if act != c.expected {
			t.Errorf("FormatShort(%f) = %s, expected %s", c.input, act, c.expected)
		}
	}
}

func TestFormatInvalidMs(t *testing.T) {
	var err error
	_, err = FormatLong(math.NaN())
	if err != ErrNaN {
		t.Error("FormatLong(NaN) expects ErrNaN")
	}
	_, err = FormatShort(math.NaN())
	if err != ErrNaN {
		t.Error("FormatShort(NaN) expects ErrNaN")
	}
	for _, sign := range []int{-1, 0, 1} {
		inf := math.Inf(sign)
		_, err = FormatLong(inf)
		if err != ErrInfinity {
			t.Errorf("FormatLong(%f) expects ErrInfinity", inf)
		}
		_, err = FormatShort(inf)
		if err != ErrInfinity {
			t.Errorf("FormatShort(%f) expects ErrInfinity", inf)
		}
	}
}

func ExampleParse() {
	ms, _ := Parse("2 days")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("1 day")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("1 day")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("10h")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("2.5 hrs")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("2h")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("1m")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("5s")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("1y")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("100")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("-3 days")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("-1h")
	fmt.Printf("%.0f\n", ms)
	ms, _ = Parse("-200")
	fmt.Printf("%.0f\n", ms)
	// -200
	// Output:
	// 172800000
	// 86400000
	// 36000000
	// 9000000
	// 7200000
	// 60000
	// 5000
	// 31557600000
	// 100
	// -259200000
	// -3600000
}
func ExampleFormatShort() {
	ms, _ := FormatShort(60000)
	fmt.Println(ms)
	ms, _ = FormatShort(2 * 60000)
	fmt.Println(ms)
	ms, _ = FormatShort(-3 * 60000)
	fmt.Println(ms)
	ms, _ = FormatShort(MustParse("10 hours"))
	fmt.Println(ms)
	// Output:
	// 1m
	// 2m
	// -3m
	// 10h
}

func ExampleFormatLong() {
	ms, _ := FormatLong(60000)
	fmt.Println(ms)
	ms, _ = FormatLong(2 * 60000)
	fmt.Println(ms)
	ms, _ = FormatLong(-3 * 60000)
	fmt.Println(ms)
	ms, _ = FormatLong(MustParse("10 hours"))
	fmt.Println(ms)
	// Output:
	// 1 minute
	// 2 minutes
	// -3 minutes
	// 10 hours
}
