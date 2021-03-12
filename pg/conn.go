package pg

import (
	"context"
	"io"
	"time"

	"go-pg-utf/pg/orm"

	"github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

// Base on go-pg v10.7.5
type Conn interface {
	AddQueryHook(hook pg.QueryHook)
	Begin() (Tx, error)
	Close() error
	Context() context.Context
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error)
	Exec(query interface{}, params ...interface{}) (res pg.Result, err error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	ExecOne(query interface{}, params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	Formatter() pgorm.QueryFormatter
	Model(model ...interface{}) orm.Query
	ModelContext(c context.Context, model ...interface{}) orm.Query
	Param(param string) interface{}
	Ping(ctx context.Context) error
	PoolStats() *pg.PoolStats
	Prepare(q string) (Stmt, error)
	Query(model, query interface{}, params ...interface{}) (res pg.Result, err error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	RunInTransaction(ctx context.Context, fn func(Tx) error) error
	WithContext(ctx context.Context) Conn
	WithParam(param string, value interface{}) Conn
	WithTimeout(d time.Duration) Conn
}

type ConnWrap struct {
	conn *pg.Conn
}

func NewConn(conn *pg.Conn) *ConnWrap {
	return &ConnWrap{conn}
}

func (c *ConnWrap) AddQueryHook(hook pg.QueryHook) {
	c.conn.AddQueryHook(hook)
}

func (c *ConnWrap) Begin() (Tx, error) {
	tx, err := c.conn.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (c *ConnWrap) Close() error {
	return c.conn.Close()
}

func (c *ConnWrap) Context() context.Context {
	return c.conn.Context()
}

func (c *ConnWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.CopyFrom(r, query, params...)
}

func (c *ConnWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.CopyTo(w, query, params...)
}

func (c *ConnWrap) Exec(query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.Exec(query, params...)
}

func (c *ConnWrap) ExecContext(ctx context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.ExecContext(ctx, query, params...)
}

func (c *ConnWrap) ExecOne(query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.ExecOne(query, params...)
}

func (c *ConnWrap) ExecOneContext(ctx context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.ExecOneContext(ctx, query, params...)
}

func (c *ConnWrap) Formatter() pgorm.QueryFormatter {
	return c.conn.Formatter()
}

func (c *ConnWrap) Model(model ...interface{}) orm.Query {
	return orm.NewQuery(c.conn.Model(model...))
}

func (c *ConnWrap) ModelContext(ctx context.Context, model ...interface{}) orm.Query {
	return orm.NewQuery(c.conn.ModelContext(ctx, model...))
}

func (c *ConnWrap) Param(param string) interface{} {
	return c.conn.Param(param)
}

func (c *ConnWrap) Ping(ctx context.Context) error {
	return c.conn.Ping(ctx)
}

func (c *ConnWrap) PoolStats() *pg.PoolStats {
	return c.conn.PoolStats()
}

func (c *ConnWrap) Prepare(q string) (Stmt, error) {
	stmt, err := c.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	return NewStmt(stmt), nil
}

func (c *ConnWrap) Query(model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.Query(model, query, params...)
}

func (c *ConnWrap) QueryContext(ctx context.Context, model, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.QueryContext(ctx, model, query, params...)
}

func (c *ConnWrap) QueryOne(model, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.QueryOne(model, query, params...)
}

func (c *ConnWrap) QueryOneContext(ctx context.Context, model, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.QueryOneContext(ctx, model, query, params...)
}

func (c *ConnWrap) RunInTransaction(ctx context.Context, fn func(Tx) error) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return err
	}
	return NewTx(tx).RunInTransaction(ctx, fn)
}

func (c *ConnWrap) WithContext(ctx context.Context) Conn {
	conn := c.conn.WithContext(ctx)
	return NewConn(conn)
}

func (c *ConnWrap) WithParam(param string, value interface{}) Conn {
	conn := c.conn.WithParam(param, value)
	return NewConn(conn)
}

func (c *ConnWrap) WithTimeout(d time.Duration) Conn {
	conn := c.conn.WithTimeout(d)
	return NewConn(conn)
}
