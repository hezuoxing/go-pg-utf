package pg

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type Stmt interface {
	Close() error
	Exec(params ...interface{}) (pg.Result, error)
	ExecContext(c context.Context, params ...interface{}) (pg.Result, error)
	ExecOne(params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, params ...interface{}) (pg.Result, error)
	Query(model interface{}, params ...interface{}) (pg.Result, error)
	QueryContext(c context.Context, model interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model interface{}, params ...interface{}) (pg.Result, error)
}

type StmtWrap struct {
	stmt *pg.Stmt
}

func NewStmt(stmt *pg.Stmt) *StmtWrap {
	return &StmtWrap{stmt}
}

func (s *StmtWrap) Close() error {
	return s.stmt.Close()
}

func (s *StmtWrap) Exec(params ...interface{}) (pg.Result, error) {
	return s.stmt.Exec(params...)
}

func (s *StmtWrap) ExecContext(c context.Context, params ...interface{}) (pg.Result, error) {
	return s.stmt.ExecContext(c, params)
}

func (s *StmtWrap) ExecOne(params ...interface{}) (pg.Result, error) {
	return s.stmt.ExecOne(params)
}

func (s *StmtWrap) ExecOneContext(c context.Context, params ...interface{}) (pg.Result, error) {
	return s.stmt.ExecOneContext(c, params)
}

func (s *StmtWrap) Query(model interface{}, params ...interface{}) (pg.Result, error) {
	return s.stmt.Query(model, params)
}

func (s *StmtWrap) QueryContext(c context.Context, model interface{}, params ...interface{}) (pg.Result, error) {
	return s.stmt.QueryContext(c, model, params...)
}

func (s *StmtWrap) QueryOne(model interface{}, params ...interface{}) (pg.Result, error) {
	return s.stmt.QueryOne(model, params...)
}

func (s *StmtWrap) QueryOneContext(c context.Context, model interface{}, params ...interface{}) (pg.Result, error) {
	return s.stmt.QueryOneContext(c, model, params...)
}
