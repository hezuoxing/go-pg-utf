package gopgmock

import (
	"context"
	"io"
	"time"

	"go-pg-utf/pg"
	"go-pg-utf/pg/orm"

	pgpg "github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

type MockDB struct {
	db MockDBWrap
}

// NewGoPGDBTest returns method that already implements orm.DB and mock instance to mocking arguments and results.
func StartMockDB() (conn pg.DB, mock *SQLMock) {
	wrapDB, mock := StartMockDBWrap()
	var db = MockDB{db: wrapDB}
	return &db, mock
}

func (d *MockDB) Connect(_ *pgpg.Options) pg.DB {
	return nil
}

func (d *MockDB) AddQueryHook(_ pgpg.QueryHook) {

}

func (d *MockDB) Begin() (pg.Tx, error) {
	return d.db.Begin()
}

func (d *MockDB) BeginContext(_ context.Context) (pg.Tx, error) {
	return d.db.Begin()
}

func (d *MockDB) Close() error {
	return d.db.Close()
}

func (d *MockDB) Conn() pg.Conn {
	return d.db.Conn()
}

func (d *MockDB) Context() context.Context {
	return d.db.context
}

func (d *MockDB) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.CopyFrom(r, query, params...)
}

func (d *MockDB) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.CopyTo(w, query, params...)
}

func (d *MockDB) Exec(query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.Exec(query, params...)
}

func (d *MockDB) ExecContext(ctx context.Context, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.ExecContext(ctx, query, params...)
}

func (d *MockDB) ExecOne(query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.ExecOne(query, params...)
}

func (d *MockDB) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.ExecOneContext(c, query, params...)
}

func (d *MockDB) Formatter() pgorm.QueryFormatter {
	return d.db.Formatter()
}

func (d *MockDB) Listen(ctx context.Context, channels ...string) pg.Listener {
	return d.db.Listen(ctx, channels...)
}

func (d *MockDB) Model(model ...interface{}) orm.Query {
	query := d.db.Model(model...)
	return orm.NewQuery(query)
}

func (d *MockDB) ModelContext(c context.Context, model ...interface{}) orm.Query {
	query := d.db.ModelContext(c, model...)
	return orm.NewQuery(query)
}

func (d *MockDB) Options() *pgpg.Options {
	return d.db.Options()
}

func (d *MockDB) Param(param string) interface{} {
	return d.db.context.Value(param)
}

func (d *MockDB) Ping(_ context.Context) error {
	return nil
}

func (d *MockDB) PoolStats() *pgpg.PoolStats {
	return d.db.PoolStats()
}

func (d *MockDB) Prepare(q string) (pg.Stmt, error) {
	return d.db.Prepare(q)
}

func (d *MockDB) Query(model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.Query(model, query, params...)
}

func (d *MockDB) QueryContext(c context.Context, model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.QueryContext(c, model, query, params...)
}

func (d *MockDB) QueryOne(model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.QueryOne(model, query, params...)
}

func (d *MockDB) QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return d.db.QueryOneContext(c, model, query, params...)
}

func (d *MockDB) RunInTransaction(ctx context.Context, fn func(pg.Tx) error) error {
	return d.db.RunInTransaction(ctx, fn)
}

func (d *MockDB) String() string {
	return d.db.String()
}

func (d *MockDB) WithContext(ctx context.Context) pg.DB {
	db := MockDB{db: d.db.WithContext(ctx)}
	return &db
}

func (d *MockDB) WithParam(param string, value interface{}) pg.DB {
	db := MockDB{db: d.db.WithParam(param, value)}
	return &db
}

func (d *MockDB) WithTimeout(dur time.Duration) pg.DB {
	db := MockDB{db: d.db.WithTimeout(dur)}
	return &db
}
