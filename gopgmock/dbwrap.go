package gopgmock

import (
	"context"
	"io"
	"time"

	"go-pg-utf/pg"

	pgpg "github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

type MockDBWrap struct {
	*baseDB
	context context.Context
	fmter   pgorm.QueryFormatter
}

// NewGoPGDBTest returns method that already implements orm.DB and mock instance to mocking arguments and results.
func StartMockDBWrap() (conn MockDBWrap, mock *SQLMock) {
	sqlMock := &SQLMock{
		queries: make(map[string]buildQuery),
	}

	goPG := MockDBWrap{
		fmter: pgorm.NewFormatter(),
	}
	goPG.baseDB = newBaseDB(&goPG, sqlMock)
	return goPG, sqlMock
}

func (d *MockDBWrap) copyDB(c context.Context) MockDBWrap {
	sqlMock := &SQLMock{
		queries: make(map[string]buildQuery),
	}

	goPG := MockDBWrap{
		fmter:   pgorm.NewFormatter(),
		context: c,
	}
	goPG.baseDB = newBaseDB(&goPG, sqlMock)
	return goPG
}

func (d *MockDBWrap) Connect(_ *pgpg.Options) pg.DB {
	return nil
}

func (d *MockDBWrap) AddQueryHook(_ pgpg.QueryHook) {

}

func (d *MockDBWrap) Begin() (pg.Tx, error) {
	tx := NewTx(d)
	return tx.Begin()
}

func (d *MockDBWrap) Close() error {
	return nil
}

func (d *MockDBWrap) Conn() pg.Conn {
	return nil
}

func (d *MockDBWrap) Context() context.Context {
	return d.context
}

func (d *MockDBWrap) CopyFrom(_ io.Reader, _ interface{}, _ ...interface{}) (res pgpg.Result, err error) {
	return nil, nil
}

func (d *MockDBWrap) CopyTo(_ io.Writer, _ interface{}, _ ...interface{}) (res pgpg.Result, err error) {
	return nil, nil
}

func (d *MockDBWrap) Exec(query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.doQuery(context.Background(), nil, query, params...)
}

func (d *MockDBWrap) ExecContext(ctx context.Context, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.doQuery(ctx, nil, query, params...)
}

func (d *MockDBWrap) ExecOne(query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.execOne(context.Background(), query, params...)
}

func (d *MockDBWrap) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.execOne(c, query, params...)
}

func (d *MockDBWrap) execOne(c context.Context, query interface{}, params ...interface{}) (pgpg.Result, error) {
	res, err := d.ExecContext(c, query, params...)
	if err != nil {
		return nil, err
	}

	if err := AssertOneRow(res.RowsAffected()); err != nil {
		return nil, err
	}
	return res, nil
}

func (d *MockDBWrap) Formatter() pgorm.QueryFormatter {
	return d.fmter
}

func (d *MockDBWrap) Listen(_ context.Context, _ ...string) pg.Listener {
	return nil
}

func (d *MockDBWrap) Model(model ...interface{}) *pgorm.Query {
	return pgorm.NewQuery(d, model...)
}

func (d *MockDBWrap) ModelContext(c context.Context, model ...interface{}) *pgorm.Query {
	return pgorm.NewQueryContext(c, d, model...)
}

func (d *MockDBWrap) Options() *pgpg.Options {
	return nil
}

func (d *MockDBWrap) Param(param string) interface{} {
	return d.context.Value(param)
}

func (d *MockDBWrap) Ping(_ context.Context) error {
	return nil
}

func (d *MockDBWrap) PoolStats() *pgpg.PoolStats {
	return nil
}

func (d *MockDBWrap) Prepare(_ string) (pg.Stmt, error) {
	return nil, nil
}

func (d *MockDBWrap) Query(model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.doQuery(context.Background(), model, query, params...)
}

func (d *MockDBWrap) QueryContext(c context.Context, model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.doQuery(c, model, query, params...)
}

func (d *MockDBWrap) QueryOne(model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.queryOne(context.Background(), model, query, params...)
}

func (d *MockDBWrap) QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.queryOne(c, model, query, params...)
}

func (d *MockDBWrap) RunInTransaction(ctx context.Context, fn func(pg.Tx) error) error {
	tx := NewTx(d)
	return tx.RunInTransaction(ctx, fn)
}

func (d *MockDBWrap) String() string {
	return ""
}

func (d *MockDBWrap) WithContext(ctx context.Context) MockDBWrap {
	db := d.copyDB(ctx)
	return db
}

func (d *MockDBWrap) WithParam(param string, value interface{}) MockDBWrap {
	db := d.copyDB(context.WithValue(d.context, param, value))
	return db
}

func (d *MockDBWrap) WithTimeout(dur time.Duration) MockDBWrap {
	c, cancel := context.WithTimeout(d.context, dur)
	defer cancel()
	db := d.copyDB(c)
	return db
}

func (d *MockDBWrap) queryOne(c context.Context, model, query interface{}, params ...interface{}) (pgpg.Result, error) {
	res, err := d.QueryContext(c, model, query, params...)
	if err != nil {
		return nil, err
	}

	if err := AssertOneRow(res.RowsAffected()); err != nil {
		return nil, err
	}
	return res, nil
}
