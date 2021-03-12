package gopgmock

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go-pg-utf/internal"

	pgorm "github.com/go-pg/pg/v10/orm"
)

type baseDB struct {
	db      pgorm.DB
	sqlMock *SQLMock
}

func newBaseDB(db pgorm.DB, mock *SQLMock) *baseDB {
	return &baseDB{
		db:      db,
		sqlMock: mock,
	}
}

func (t *baseDB) doQuery(_ context.Context, dst interface{}, query interface{}, params ...interface{}) (pgorm.Result, error) {
	wantedQuery, err := internal.AppendQuery(t.db.Formatter(), nil, query, params...)
	if err != nil {
		return nil, err
	}
	wantedQueryStr := strings.ReplaceAll(string(wantedQuery), " ", "")
	wantedQueryStr = strings.ReplaceAll(wantedQueryStr, "\"", "")
	wantedQueryStr = strings.ToLower(wantedQueryStr)
	for k, v := range t.sqlMock.queries {
		onTheList := t.db.Formatter().FormatQuery(nil, k, v.params...)
		onTheListQueryStr := strings.ReplaceAll(string(onTheList), " ", "")
		onTheListQueryStr = strings.ReplaceAll(onTheListQueryStr, "\"", "")
		onTheListQueryStr = strings.ToLower(onTheListQueryStr)
		if onTheListQueryStr == wantedQueryStr {
			log.Printf("tx query %s=%s", onTheListQueryStr, wantedQueryStr)
			if dst == nil {
				return v.result, v.err
			}

			if v.result.model == nil {
				return v.result, v.err
			}

			err = internal.MockData(dst, v.result.model)
			if err != nil {
				return v.result, err
			}
			return v.result, v.err
		}
	}

	return nil, fmt.Errorf("no mock expectation result")
}
