package identifydateformat

import (
	"fmt"
	"time"
)

var (
	//DateFormats set which date formats are going to be tested against the dates provided by user
	DateFormats = []string{
		// DMY
		"02-01-2006",
		"02.01.2006",
		"02/01/2006",
		"02-01-06",
		"02.01.06",
		"02/01/06",

		// YMD
		"2006-01-02",
		"2006/01/02",
		"06-01-02",
		"06/01/02",

		// MDY
		"01-02-2006",
		"01.02.2006",
		"01/02/2006",
		"01-02-06",
		"01.02.06",
		"01/02/06",
	}
)

// YEAR: 2006
// MONTH: 01
// DAY: 02

// DefinePossibleFormats replaces the original date formats to a custom one
func DefinePossibleFormats(formats []string) {
	DateFormats = formats
}

// AddPossibleFormats includes any new formats that were not present yet in
// DateFormats
func AddPossibleFormats(formats ...string) {
	for i := 0; i < len(formats); i++ {
		formatFound := false

		for j := 0; j < len(DateFormats); j++ {
			if formats[i] == DateFormats[j] {
				formatFound = true
				break
			}
		}
		if !formatFound {
			DateFormats = append(DateFormats, formats[i])
		}
	}
}

// IdentifyDateFormat runs all provided dates against DateFormats (a set of predefined date
// formats). It returns all the date formats that can be parsed for the given
// set of dates
func IdentifyDateFormat(dates []string) ([]string, error) {
	dateFormats := DateFormats
	for _, date := range dates {
		for i := 0; i < len(dateFormats); i++ {
			_, err := time.Parse(dateFormats[i], date)
			if err != nil {
				dateFormats = append(dateFormats[:i], dateFormats[i+1:]...)
				i--
			}
		}
	}
	if len(dateFormats) == 0 {
		return dateFormats, fmt.Errorf("no date formats could be parsed with the given dates")
	}
	return dateFormats, nil
}
