package pg

import (
	"context"
	"io"
	"time"

	"go-pg-utf/pg/orm"

	"github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

type DB interface {
	Connect(opt *pg.Options) DB
	AddQueryHook(hook pg.QueryHook)
	Begin() (Tx, error)
	BeginContext(ctx context.Context) (Tx, error)
	Close() error
	Conn() Conn
	Context() context.Context
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error)
	Exec(query interface{}, params ...interface{}) (res pg.Result, err error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	ExecOne(query interface{}, params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	Formatter() pgorm.QueryFormatter
	Listen(ctx context.Context, channels ...string) Listener
	Model(model ...interface{}) orm.Query
	ModelContext(c context.Context, model ...interface{}) orm.Query
	Options() *pg.Options
	Param(param string) interface{}
	Ping(ctx context.Context) error
	PoolStats() *pg.PoolStats
	Prepare(q string) (Stmt, error)
	Query(model, query interface{}, params ...interface{}) (res pg.Result, err error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	RunInTransaction(ctx context.Context, fn func(Tx) error) error
	String() string
	WithContext(ctx context.Context) DB
	WithParam(param string, value interface{}) DB
	WithTimeout(dur time.Duration) DB
}

type DBWrap struct {
	db *pg.DB
}

func NewDB(db *pg.DB) *DBWrap {
	return &DBWrap{db}
}

func (d *DBWrap) Connect(opt *pg.Options) DB {
	db := pg.Connect(opt)
	return NewDB(db)
}

func (d *DBWrap) AddQueryHook(hook pg.QueryHook) {
	d.db.AddQueryHook(hook)
}

func (d *DBWrap) Begin() (Tx, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (d *DBWrap) BeginContext(ctx context.Context) (Tx, error) {
	tx, err := d.db.BeginContext(ctx)
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (d *DBWrap) Close() error {
	return d.db.Close()
}

func (d *DBWrap) Conn() Conn {
	return NewConn(d.db.Conn())
}

func (d *DBWrap) Context() context.Context {
	return d.db.Context()
}

func (d *DBWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.CopyFrom(r, query, params...)
}

func (d *DBWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.CopyTo(w, query, params...)
}

func (d *DBWrap) Exec(query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.Exec(query, params...)
}

func (d *DBWrap) ExecContext(ctx context.Context, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.ExecContext(ctx, query, params...)
}

func (d *DBWrap) ExecOne(query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.ExecOne(query, params...)
}

func (d *DBWrap) ExecOneContext(ctx context.Context, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.ExecOneContext(ctx, query, params...)
}

func (d *DBWrap) Formatter() pgorm.QueryFormatter {
	return d.db.Formatter()
}

func (d *DBWrap) Listen(ctx context.Context, channels ...string) Listener {
	listener := d.db.Listen(ctx, channels...)
	return NewListener(listener)
}

func (d *DBWrap) Model(model ...interface{}) orm.Query {
	query := d.db.Model(model...)
	return orm.NewQuery(query)
}

func (d *DBWrap) ModelContext(c context.Context, model ...interface{}) orm.Query {
	query := d.db.ModelContext(c, model...)
	return orm.NewQuery(query)
}

func (d *DBWrap) Options() *pg.Options {
	return d.db.Options()
}

func (d *DBWrap) Param(param string) interface{} {
	return d.db.Param(param)
}

func (d *DBWrap) Ping(ctx context.Context) error {
	return d.db.Ping(ctx)
}

func (d *DBWrap) PoolStats() *pg.PoolStats {
	return d.db.PoolStats()
}

func (d *DBWrap) Prepare(q string) (Stmt, error) {
	stmt, err := d.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	return NewStmt(stmt), nil
}

func (d *DBWrap) Query(model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.Query(model, query, params...)
}

func (d *DBWrap) QueryContext(c context.Context, model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.QueryContext(c, model, query, params...)
}

func (d *DBWrap) QueryOne(model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.QueryOne(model, query, params...)
}

func (d *DBWrap) QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.QueryOneContext(c, model, query, params...)
}

func (d *DBWrap) RunInTransaction(ctx context.Context, fn func(Tx) error) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	return NewTx(tx).RunInTransaction(ctx, fn)
}

func (d *DBWrap) String() string {
	return d.db.String()
}

func (d *DBWrap) WithContext(ctx context.Context) DB {
	db := d.db.WithContext(ctx)
	return NewDB(db)
}

func (d *DBWrap) WithParam(param string, value interface{}) DB {
	db := d.db.WithParam(param, value)
	return NewDB(db)
}

func (d *DBWrap) WithTimeout(dur time.Duration) DB {
	db := d.db.WithTimeout(dur)
	return NewDB(db)
}
