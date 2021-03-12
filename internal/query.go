package internal

import (
	"fmt"

	pgorm "github.com/go-pg/pg/v10/orm"
)

func AppendQuery(fmter pgorm.QueryFormatter, dst []byte, query interface{}, params ...interface{}) ([]byte, error) {
	switch query := query.(type) {
	case pgorm.QueryAppender:
		if v, ok := fmter.(*pgorm.Formatter); ok {
			fmter = v.WithModel(query)
		}
		return query.AppendQuery(fmter, dst)
	case string:
		if len(params) > 0 {
			model, ok := params[len(params)-1].(pgorm.TableModel)
			if ok {
				if v, ok := fmter.(*pgorm.Formatter); ok {
					fmter = v.WithTableModel(model)
					params = params[:len(params)-1]
				}
			}
		}
		return fmter.FormatQuery(dst, query, params...), nil
	default:
		return nil, fmt.Errorf("pg: can't append %T", query)
	}
}
