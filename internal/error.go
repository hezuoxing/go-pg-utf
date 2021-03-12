package internal

import "fmt"

var ErrNoRows = fmt.Errorf("pg: no rows in result set")
var ErrMultiRows = fmt.Errorf("pg: multiple rows in result set")

func AssertOneRow(l int) error {
	switch {
	case l == 0:
		return ErrNoRows
	case l > 1:
		return ErrMultiRows
	default:
		return nil
	}
}
