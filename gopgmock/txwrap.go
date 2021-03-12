package gopgmock

import (
	"context"
	"io"

	"go-pg-utf/pg"

	pgpg "github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

type MockTxWrap struct {
	*baseDB
	ctx context.Context
	tx  pg.Tx
}

func NewTxWrap(db *MockDBWrap, tx pg.Tx) *MockTxWrap {
	mock := &MockTxWrap{
		ctx: db.Context(),
		tx:  tx,
	}
	mock.baseDB = newBaseDB(mock, db.sqlMock)
	return mock
}

func (t *MockTxWrap) Begin() (pg.Tx, error) {
	_, lastErr := t.ExecContext(t.db.Context(), "BEGIN")
	if lastErr != nil {
		return nil, lastErr
	}
	return t.tx, nil
}

func (t *MockTxWrap) Close() error {
	return nil
}

func (t *MockTxWrap) CloseContext(_ context.Context) error {
	return nil
}

func (t *MockTxWrap) Commit() error {
	_, err := t.Exec("COMMIT")
	return err
}

func (t *MockTxWrap) CommitContext(ctx context.Context) error {
	_, err := t.ExecContext(ctx, "COMMIT")
	return err
}

func (t *MockTxWrap) Context() context.Context {
	return t.ctx
}

func (t *MockTxWrap) CopyFrom(_ io.Reader, _ interface{}, _ ...interface{}) (res pgpg.Result, err error) {
	return nil, nil
}

func (t *MockTxWrap) CopyTo(_ io.Writer, _ interface{}, _ ...interface{}) (res pgpg.Result, err error) {
	return nil, nil
}

func (t *MockTxWrap) Exec(query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.doQuery(context.Background(), nil, query, params...)
}

func (t *MockTxWrap) ExecContext(c context.Context, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.doQuery(c, nil, query, params...)
}

func (t *MockTxWrap) ExecOne(query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.execOne(context.Background(), query, params...)
}

func (t *MockTxWrap) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.execOne(c, query, params...)
}

func (t *MockTxWrap) execOne(c context.Context, query interface{}, params ...interface{}) (pgpg.Result, error) {
	res, err := t.ExecContext(c, query, params...)
	if err != nil {
		return nil, err
	}

	if err := AssertOneRow(res.RowsAffected()); err != nil {
		return nil, err
	}
	return res, nil
}

func (t *MockTxWrap) Formatter() pgorm.QueryFormatter {
	return t.db.Formatter()
}

func (t *MockTxWrap) Model(model ...interface{}) *pgorm.Query {
	return pgorm.NewQuery(t, model...)
}

func (t *MockTxWrap) ModelContext(c context.Context, model ...interface{}) *pgorm.Query {
	return pgorm.NewQueryContext(c, t, model...)
}

func (t *MockTxWrap) Prepare(_ string) (*pgpg.Stmt, error) {
	return nil, nil
}

func (t *MockTxWrap) Query(model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.doQuery(context.Background(), model, query, params...)
}

func (t *MockTxWrap) QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.doQuery(c, model, query, params...)
}

func (t *MockTxWrap) QueryOne(model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.queryOne(context.Background(), model, query, params...)
}

func (t *MockTxWrap) QueryOneContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pgpg.Result, error) {
	return t.queryOne(c, model, query, params...)
}

func (t *MockTxWrap) queryOne(c context.Context, model, query interface{}, params ...interface{}) (pgpg.Result, error) {
	res, err := t.QueryContext(c, model, query, params...)
	if err != nil {
		return nil, err
	}

	if err := AssertOneRow(res.RowsAffected()); err != nil {
		return nil, err
	}
	return res, nil
}

func (t *MockTxWrap) Rollback() error {
	_, err := t.Exec("ROLLBACK")
	return err
}

func (t *MockTxWrap) RollbackContext(ctx context.Context) error {
	_, err := t.ExecContext(ctx, "ROLLBACK")
	return err
}

func (t *MockTxWrap) RunInTransaction(ctx context.Context, fn func(tx pg.Tx) error) error {
	defer func() {
		if err := recover(); err != nil {
			_ = t.RollbackContext(ctx)
			panic(err)
		}
	}()
	tx, err := t.Begin()
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		_ = t.RollbackContext(ctx)
		return err
	}
	return t.CommitContext(ctx)
}

func (t *MockTxWrap) Stmt(_ pg.Stmt) pg.Stmt {
	return nil
}
