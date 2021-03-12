package gopgmock

import (
	"context"
	"io"

	"go-pg-utf/pg"
	"go-pg-utf/pg/orm"

	pgpg "github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

type MockTx struct {
	tx *MockTxWrap
}

func NewTx(db *MockDBWrap) pg.Tx {
	tx := MockTx{}
	tx.tx = NewTxWrap(db, &tx)
	return &tx
}

func (t *MockTx) Begin() (pg.Tx, error) {
	tx, err := t.tx.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *MockTx) Close() error {
	return t.tx.Close()
}

func (t *MockTx) CloseContext(ctx context.Context) error {
	return t.tx.CloseContext(ctx)
}

func (t *MockTx) Commit() error {
	return t.tx.Commit()
}

func (t *MockTx) CommitContext(ctx context.Context) error {
	return t.tx.CommitContext(ctx)
}

func (t *MockTx) Context() context.Context {
	return t.tx.Context()
}

func (t *MockTx) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return t.tx.CopyFrom(r, query, params...)
}

func (t *MockTx) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pgpg.Result, err error) {
	return t.tx.CopyTo(w, query, params...)
}

func (t *MockTx) Exec(query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.Exec(query, params...)
}

func (t *MockTx) ExecContext(c context.Context, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.ExecContext(c, query, params...)
}

func (t *MockTx) ExecOne(query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.ExecOne(query, params...)
}

func (t *MockTx) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.ExecOneContext(c, query, params...)
}

func (t *MockTx) Formatter() pgorm.QueryFormatter {
	return t.tx.Formatter()
}

func (t *MockTx) Model(model ...interface{}) orm.Query {
	return orm.NewQuery(t.tx.Model(model...))
}

func (t *MockTx) ModelContext(c context.Context, model ...interface{}) orm.Query {
	return orm.NewQuery(t.tx.ModelContext(c, model...))
}

func (t *MockTx) Prepare(q string) (*pgpg.Stmt, error) {
	return t.tx.Prepare(q)
}

func (t *MockTx) Query(model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.Query(model, query, params...)
}

func (t *MockTx) QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.QueryContext(c, model, query, params...)
}

func (t *MockTx) QueryOne(model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.QueryOne(model, query, params...)
}

func (t *MockTx) QueryOneContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.tx.QueryOneContext(c, model, query, params...)
}

func (t *MockTx) Rollback() error {
	return t.tx.Rollback()
}

func (t *MockTx) RollbackContext(ctx context.Context) error {
	return t.tx.RollbackContext(ctx)
}

func (t *MockTx) RunInTransaction(ctx context.Context, fn func(tx pg.Tx) error) error {
	return t.tx.RunInTransaction(ctx, func(tx pg.Tx) error { return fn(tx) })
}

func (t *MockTx) Stmt(stmt pg.Stmt) pg.Stmt {
	return t.tx.Stmt(stmt)
}
