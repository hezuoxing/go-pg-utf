package gopgmock

import (
	"fmt"

	"github.com/go-pg/pg/v10/orm"
)

var ErrNoRows = fmt.Errorf("pg: no rows in result set")
var ErrMultiRows = fmt.Errorf("pg: multiple rows in result set")

type OrmResult struct {
	rowsAffected int
	rowsReturned int
	model        interface{}
}

func (o *OrmResult) Model() orm.Model {
	if o.model == nil {
		return nil
	}

	model, err := orm.NewModel(o.model)
	if err != nil {
		return nil
	}

	return model
}

func (o *OrmResult) RowsAffected() int {
	return o.rowsAffected
}

func (o *OrmResult) RowsReturned() int {
	return o.rowsReturned
}

// NewResult implements orm.Result in go-pg package
func NewResult(rowAffected, rowReturned int, model interface{}) *OrmResult {
	return &OrmResult{
		rowsAffected: rowAffected,
		rowsReturned: rowReturned,
		model:        model,
	}
}

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
