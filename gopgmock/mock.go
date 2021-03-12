package gopgmock

import (
	"strings"
	"sync"
)

type SQLMock struct {
	lock          sync.RWMutex
	currentQuery  string
	currentParams []interface{}
	queries       map[string]buildQuery
}

func (sqlMock *SQLMock) ExpectExec(query string) *SQLMock {
	sqlMock.lock.Lock()
	defer sqlMock.lock.Unlock()

	sqlMock.currentQuery = strings.TrimSpace(query)
	return sqlMock
}

func (sqlMock *SQLMock) ExpectQuery(query string) *SQLMock {
	sqlMock.lock.Lock()
	defer sqlMock.lock.Unlock()

	sqlMock.currentQuery = strings.TrimSpace(query)
	return sqlMock
}

func (sqlMock *SQLMock) WithArgs(params ...interface{}) *SQLMock {
	sqlMock.lock.Lock()
	defer sqlMock.lock.Unlock()

	sqlMock.currentParams = make([]interface{}, 0)
	sqlMock.currentParams = append(sqlMock.currentParams, params...)
	return sqlMock
}

func (sqlMock *SQLMock) Returns(result *OrmResult, err error) {
	sqlMock.lock.Lock()
	defer sqlMock.lock.Unlock()

	q := buildQuery{
		query:  sqlMock.currentQuery,
		params: sqlMock.currentParams,
		result: result,
		err:    err,
	}

	sqlMock.queries[sqlMock.currentQuery] = q
	sqlMock.currentQuery = ""
	sqlMock.currentParams = nil
}

func (sqlMock *SQLMock) FlushAll() {
	sqlMock.lock.Lock()
	defer sqlMock.lock.Unlock()

	sqlMock.currentQuery = ""
	sqlMock.currentParams = nil
	sqlMock.queries = make(map[string]buildQuery)
}
