package identifydateformat

import (
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

func definePossibleFormats(formats []string) {
	DateFormats = formats
}

func identifydateformat(dates []string) []string {
	for _, date := range dates {
		for i := 0; i < len(DateFormats); i++ {
			_, err := time.Parse(DateFormats[i], date)
			if err != nil {
				DateFormats = append(DateFormats[:i], DateFormats[i+1:]...)
				i--
			}
		}
	}
	return DateFormats
}
