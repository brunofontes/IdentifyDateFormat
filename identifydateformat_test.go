package identifydateformat

import (
	"fmt"
	"testing"
)

// YEAR: 2006
// MONTH: 01
// DAY: 02

func TestOpenCsv(t *testing.T) {
	fmt.Printf("\n%#v\n", identifydateformat([]string{"02/02/02", "23/12/01", "10/12/32", "01/10/01"}))
}
