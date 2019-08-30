package utli

import (
	"strings"
	"time"
)

var dateFormat = []string{
	"Y", "2006",
	"y", "06",
	"m", "01",
	"n", "1",
	"M", "Jan",
	"F", "January",
	"d", "02",
	"j", "2",
	"D", "Mon",
	"l", "Monday",
	"g", "3",
	"G", "15",
	"h", "03",
	"H", "15",
	"a", "pm",
	"A", "PM",
	"i", "04",
	"s", "05",
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",
	"r", time.RFC1123Z,
}

//Date php-date function
func Date(format string, ts ...time.Time) string {
	replacer := strings.NewReplacer(dateFormat...)
	format = replacer.Replace(format)
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}
