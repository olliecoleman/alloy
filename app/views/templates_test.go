package views

import (
	"fmt"
	"html/template"
	"testing"
	"time"
)

func Test_emptyFn(t *testing.T) {
	var tests = []struct {
		input  interface{}
		output bool
	}{
		{"", true},
		{"hello", false},
		{[]string{}, true},
		{[]string{"foo"}, false},
		{nil, true},
		{time.Now, false},
	}

	for _, test := range tests {
		result := emptyFn(test.input)
		if result != test.output {
			t.Errorf("Expected %v : Got %v", test.output, result)
		}
	}
}

func Test_dateFn(t *testing.T) {
	d := time.Now()
	var tests = []struct {
		format string
		input  interface{}
		output string
	}{
		{"2-1-2006", nil, fmt.Sprintf("%d-%d-%d", d.Day(), d.Month(), d.Year())},
		{"02-01-2006", time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), "17-11-2009"},
	}

	for _, test := range tests {
		result := dateFn(test.format, test.input)
		if result != test.output {
			t.Errorf("Expected %v : Got %v", test.output, result)
		}
	}
}

func Test_readableDateFn(t *testing.T) {
	day := 24 * time.Hour
	month := 30 * day
	year := 12 * month
	var tests = []struct {
		input  time.Time
		output string
	}{
		{time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), "7 years ago"},
		{time.Now().Add(-2 * time.Minute), "a few minutes ago"},
		{time.Now().Add(-50 * time.Minute), "50 minutes ago"},
		{time.Now().Add(-11 * time.Hour), "11 hours ago"},
		{time.Now().Add(-29 * 24 * time.Hour), "29 days ago"},
		{time.Now().Add(-11 * month), "11 months ago"},
		{time.Now().Add(10 * year), ""},
	}

	for _, test := range tests {
		result := readableDateFn(test.input)
		if result != test.output {
			t.Errorf("Expected %v : Got %v", test.output, result)
		}
	}
}

func Test_raw(t *testing.T) {
	var tests = []struct {
		input  string
		output template.HTML
	}{
		{"<p>hello</p>", "<p>hello</p>"},
		{"<script>hello</script>", ""},
		{`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS<a>`, "XSS"},
		{"str", "str"},
	}
	for _, test := range tests {
		result := raw(test.input)
		if result != test.output {
			t.Errorf("Expected %v : Got %v", test.output, result)
		}
	}
}
