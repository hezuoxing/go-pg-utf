package pg

import (
	"context"
	"io"

	"go-pg-utf/pg/orm"

	"github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

type Tx interface {
	Begin() (Tx, error)
	Close() error
	CloseContext(ctx context.Context) error
	Commit() error
	CommitContext(ctx context.Context) error
	Context() context.Context
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error)
	Exec(query interface{}, params ...interface{}) (pg.Result, error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	ExecOne(query interface{}, params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	Formatter() pgorm.QueryFormatter
	Model(model ...interface{}) orm.Query
	ModelContext(c context.Context, model ...interface{}) orm.Query
	Prepare(q string) (*pg.Stmt, error)
	Query(model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	Rollback() error
	RollbackContext(ctx context.Context) error
	RunInTransaction(ctx context.Context, fn func(Tx) error) error
	Stmt(stmt Stmt) Stmt
}

type TxWrap struct {
	tx *pg.Tx
}

func NewTx(tx *pg.Tx) *TxWrap {
	return &TxWrap{tx}
}

func (t *TxWrap) Begin() (Tx, error) {
	tx, err := t.tx.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (t *TxWrap) Close() error {
	return t.tx.Close()
}

func (t *TxWrap) CloseContext(ctx context.Context) error {
	return t.tx.CloseContext(ctx)
}

func (t *TxWrap) Commit() error {
	return t.tx.Commit()
}

func (t *TxWrap) CommitContext(ctx context.Context) error {
	return t.tx.CommitContext(ctx)
}

func (t *TxWrap) Context() context.Context {
	return t.tx.Context()
}

func (t *TxWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return t.tx.CopyFrom(r, query, params...)
}

func (t *TxWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return t.tx.CopyTo(w, query, params...)
}

func (t *TxWrap) Exec(query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.Exec(query, params...)
}

func (t *TxWrap) ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.ExecContext(c, query, params...)
}

func (t *TxWrap) ExecOne(query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.ExecOne(query, params...)
}

func (t *TxWrap) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.ExecOneContext(c, query, params...)
}

func (t *TxWrap) Formatter() pgorm.QueryFormatter {
	return t.tx.Formatter()
}

func (t *TxWrap) Model(model ...interface{}) orm.Query {
	return orm.NewQuery(t.tx.Model(model...))
}

func (t *TxWrap) ModelContext(c context.Context, model ...interface{}) orm.Query {
	return orm.NewQuery(t.tx.ModelContext(c, model...))
}

func (t *TxWrap) Prepare(q string) (*pg.Stmt, error) {
	return t.tx.Prepare(q)
}

func (t *TxWrap) Query(model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.Query(model, query, params...)
}

func (t *TxWrap) QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.QueryContext(c, model, query, params...)
}

func (t *TxWrap) QueryOne(model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.QueryOne(model, query, params...)
}

func (t *TxWrap) QueryOneContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.QueryOneContext(c, model, query, params...)
}

func (t *TxWrap) Rollback() error {
	return t.tx.Rollback()
}

func (t *TxWrap) RollbackContext(ctx context.Context) error {
	return t.tx.RollbackContext(ctx)
}

func (t *TxWrap) RunInTransaction(ctx context.Context, fn func(Tx) error) error {
	return t.tx.RunInTransaction(ctx, func(tx *pg.Tx) error { return fn(NewTx(tx)) })
}

func (t *TxWrap) Stmt(stmt Stmt) Stmt {
	return NewStmt(t.tx.Stmt(stmt.(*StmtWrap).stmt))
}
