package orm

import (
	"context"
	"io"

	pgpg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Query interface {
	NewQuery(db *pgpg.DB, model ...interface{}) Query
	NewQueryContext(ctx context.Context, db *pgpg.DB, model ...interface{}) Query
	AllWithDeleted() Query
	AppendQuery(fmter orm.QueryFormatter, b []byte) ([]byte, error)
	Apply(fn func(Query) (Query, error)) Query
	Clone() Query
	Column(columns ...string) Query
	ColumnExpr(expr string, params ...interface{}) Query
	Context(c context.Context) Query
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error)
	Count() (int, error)
	CountEstimate(threshold int) (int, error)
	CreateComposite(opt *orm.CreateCompositeOptions) error
	CreateTable(opt *orm.CreateTableOptions) error
	DB(db *pgpg.DB) Query
	Delete(values ...interface{}) (orm.Result, error)
	Deleted() Query
	Distinct() Query
	DistinctOn(expr string, params ...interface{}) Query
	DropComposite(opt *orm.DropCompositeOptions) error
	DropTable(opt *orm.DropTableOptions) error
	Except(other Query) Query
	ExceptAll(other Query) Query
	ExcludeColumn(columns ...string) Query
	Exec(query interface{}, params ...interface{}) (orm.Result, error)
	ExecOne(query interface{}, params ...interface{}) (orm.Result, error)
	Exists() (bool, error)
	First() error
	For(s string, params ...interface{}) Query
	ForEach(fn interface{}) error
	ForceDelete(values ...interface{}) (orm.Result, error)
	Group(columns ...string) Query
	GroupExpr(group string, params ...interface{}) Query
	Having(having string, params ...interface{}) Query
	Insert(values ...interface{}) (orm.Result, error)
	Intersect(other Query) Query
	IntersectAll(other Query) Query
	Join(join string, params ...interface{}) Query
	JoinOn(condition string, params ...interface{}) Query
	JoinOnOr(condition string, params ...interface{}) Query
	Last() error
	Limit(n int) Query
	Model(model ...interface{}) Query
	New() Query
	Offset(n int) Query
	OnConflict(s string, params ...interface{}) Query
	Order(orders ...string) Query
	OrderExpr(order string, params ...interface{}) Query
	Query(model, query interface{}, params ...interface{}) (orm.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (orm.Result, error)
	Relation(name string, apply ...func(Query) (Query, error)) Query
	Returning(s string, params ...interface{}) Query
	Select(values ...interface{}) error
	SelectAndCount(values ...interface{}) (count int, firstErr error)
	SelectAndCountEstimate(threshold int, values ...interface{}) (count int, firstErr error)
	SelectOrInsert(values ...interface{}) (inserted bool, _ error)
	Set(set string, params ...interface{}) Query
	Table(tables ...string) Query
	TableExpr(expr string, params ...interface{}) Query
	TableModel() orm.TableModel
	Union(other Query) Query
	UnionAll(other Query) Query
	Update(scan ...interface{}) (orm.Result, error)
	UpdateNotZero(scan ...interface{}) (orm.Result, error)
	Value(column string, value string, params ...interface{}) Query
	Where(condition string, params ...interface{}) Query
	WhereGroup(fn func(Query) (Query, error)) Query
	WhereIn(where string, slice interface{}) Query
	WhereInMulti(where string, values ...interface{}) Query
	WhereNotGroup(fn func(Query) (Query, error)) Query
	WhereOr(condition string, params ...interface{}) Query
	WhereOrGroup(fn func(Query) (Query, error)) Query
	WhereOrNotGroup(fn func(Query) (Query, error)) Query
	WherePK() Query
	With(name string, subq Query) Query
	WithDelete(name string, subq Query) Query
	WithInsert(name string, subq Query) Query
	WithUpdate(name string, subq Query) Query
	WrapWith(name string) Query

	// customize function
	GetQuery() *orm.Query
}

type QueryWrap struct {
	query *orm.Query
}

func NewQuery(query *orm.Query) *QueryWrap {
	return &QueryWrap{query}
}

func (q *QueryWrap) GetQuery() *orm.Query {
	return q.query
}

func (q *QueryWrap) NewQuery(db *pgpg.DB, model ...interface{}) Query {
	return NewQuery(db.Model(model...))
}

func (q *QueryWrap) NewQueryContext(ctx context.Context, db *pgpg.DB, model ...interface{}) Query {
	return NewQuery(db.ModelContext(ctx, model...))
}

func (q *QueryWrap) AllWithDeleted() Query {
	q.query.AllWithDeleted()
	return q
}

func (q *QueryWrap) AppendQuery(fmter orm.QueryFormatter, b []byte) ([]byte, error) {
	return q.query.AppendQuery(fmter, b)
}

func (q *QueryWrap) Apply(fn func(Query) (Query, error)) Query {
	qq := q.query.Apply(func(pgquery *orm.Query) (*orm.Query, error) {
		query := NewQuery(pgquery)
		_, err := fn(query)
		return query.GetQuery(), err
	})
	return NewQuery(qq)
}

func (q *QueryWrap) Clone() Query {
	return NewQuery(q.query.Clone())
}

func (q *QueryWrap) Column(columns ...string) Query {
	q.query.Column(columns...)
	return q
}

func (q *QueryWrap) ColumnExpr(expr string, params ...interface{}) Query {
	q.query.ColumnExpr(expr, params...)
	return q
}

func (q *QueryWrap) Context(c context.Context) Query {
	q.query.Context(c)
	return q
}

func (q *QueryWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.query.CopyFrom(r, query, params...)
}

func (q *QueryWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.query.CopyTo(w, query, params...)
}

func (q *QueryWrap) Count() (int, error) {
	return q.query.Count()
}

func (q *QueryWrap) CountEstimate(threshold int) (int, error) {
	return q.query.CountEstimate(threshold)
}

func (q *QueryWrap) CreateComposite(opt *orm.CreateCompositeOptions) error {
	return q.query.CreateComposite(opt)
}

func (q *QueryWrap) CreateTable(opt *orm.CreateTableOptions) error {
	return q.query.CreateTable(opt)
}

func (q *QueryWrap) DB(db *pgpg.DB) Query {
	q.query.DB(db)
	return q
}

func (q *QueryWrap) Delete(values ...interface{}) (orm.Result, error) {
	return q.query.Delete(values...)
}

func (q *QueryWrap) Deleted() Query {
	q.query.Deleted()
	return q
}

func (q *QueryWrap) Distinct() Query {
	q.query.Distinct()
	return q
}

func (q *QueryWrap) DistinctOn(expr string, params ...interface{}) Query {
	q.query.DistinctOn(expr, params...)
	return q
}

func (q *QueryWrap) DropComposite(opt *orm.DropCompositeOptions) error {
	return q.query.DropComposite(opt)
}

func (q *QueryWrap) DropTable(opt *orm.DropTableOptions) error {
	return q.query.DropTable(opt)
}

func (q *QueryWrap) Except(other Query) Query {
	q.query.Except(other.GetQuery())
	return q
}

func (q *QueryWrap) ExceptAll(other Query) Query {
	q.query.ExceptAll(other.GetQuery())
	return q
}

func (q *QueryWrap) ExcludeColumn(columns ...string) Query {
	q.query.ExcludeColumn(columns...)
	return q
}

func (q *QueryWrap) Exec(query interface{}, params ...interface{}) (orm.Result, error) {
	return q.query.Exec(query, params...)
}

func (q *QueryWrap) ExecOne(query interface{}, params ...interface{}) (orm.Result, error) {
	return q.query.ExecOne(query, params...)
}

func (q *QueryWrap) Exists() (bool, error) {
	return q.query.Exists()
}

func (q *QueryWrap) First() error {
	return q.query.First()
}

func (q *QueryWrap) For(s string, params ...interface{}) Query {
	q.query.For(s, params...)
	return q
}

func (q *QueryWrap) ForEach(fn interface{}) error {
	return q.query.ForEach(fn)
}

func (q *QueryWrap) ForceDelete(values ...interface{}) (orm.Result, error) {
	return q.query.ForceDelete(values...)
}

func (q *QueryWrap) Group(columns ...string) Query {
	q.query.Group(columns...)
	return q
}

func (q *QueryWrap) GroupExpr(group string, params ...interface{}) Query {
	q.query.GroupExpr(group, params...)
	return q
}

func (q *QueryWrap) Having(having string, params ...interface{}) Query {
	q.query.Having(having, params...)
	return q
}

func (q *QueryWrap) Insert(values ...interface{}) (orm.Result, error) {
	return q.query.Insert(values...)
}

func (q *QueryWrap) Intersect(other Query) Query {
	q.query.Intersect(other.GetQuery())
	return q
}

func (q *QueryWrap) IntersectAll(other Query) Query {
	q.query.IntersectAll(other.GetQuery())
	return q
}

func (q *QueryWrap) Join(join string, params ...interface{}) Query {
	q.query.Join(join, params...)
	return q
}

func (q *QueryWrap) JoinOn(condition string, params ...interface{}) Query {
	q.query.JoinOn(condition, params...)
	return q
}

func (q *QueryWrap) JoinOnOr(condition string, params ...interface{}) Query {
	q.query.JoinOnOr(condition, params...)
	return q
}

func (q *QueryWrap) Last() error {
	return q.query.Last()
}

func (q *QueryWrap) Limit(n int) Query {
	q.query.Limit(n)
	return q
}

func (q *QueryWrap) Model(model ...interface{}) Query {
	q.query.Model(model...)
	return q
}

func (q *QueryWrap) New() Query {
	return NewQuery(q.query.New())
}

func (q *QueryWrap) Offset(n int) Query {
	q.query.Offset(n)
	return q
}

func (q *QueryWrap) OnConflict(s string, params ...interface{}) Query {
	q.query.OnConflict(s, params...)
	return q
}

func (q *QueryWrap) Order(orders ...string) Query {
	q.query.Order(orders...)
	return q
}

func (q *QueryWrap) OrderExpr(order string, params ...interface{}) Query {
	q.query.OrderExpr(order, params...)
	return q
}

func (q *QueryWrap) Query(model, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.query.Query(model, query, params...)
}

func (q *QueryWrap) QueryOne(model, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.query.QueryOne(model, query, params...)
}

func (q *QueryWrap) Relation(name string, apply ...func(Query) (Query, error)) Query {
	qfn := make([]func(*orm.Query) (*orm.Query, error), len(apply))
	for i, fn := range apply {
		qfn[i] = func(pgquery *orm.Query) (*orm.Query, error) {
			query := NewQuery(pgquery)
			_, err := fn(query)
			return query.GetQuery(), err
		}
	}
	q.query.Relation(name, qfn...)
	return q
}

func (q *QueryWrap) Returning(s string, params ...interface{}) Query {
	q.query.Returning(s, params...)
	return q
}

func (q *QueryWrap) Select(values ...interface{}) error {
	return q.query.Select(values...)
}

func (q *QueryWrap) SelectAndCount(values ...interface{}) (count int, firstErr error) {
	return q.query.SelectAndCount(values...)
}

func (q *QueryWrap) SelectAndCountEstimate(threshold int, values ...interface{}) (count int, firstErr error) {
	return q.query.SelectAndCountEstimate(threshold, values...)
}

func (q *QueryWrap) SelectOrInsert(values ...interface{}) (inserted bool, _ error) {
	return q.query.SelectOrInsert(values...)
}

func (q *QueryWrap) Set(set string, params ...interface{}) Query {
	q.query.Set(set, params...)
	return q
}

func (q *QueryWrap) Table(tables ...string) Query {
	q.query.Table(tables...)
	return q
}

func (q *QueryWrap) TableExpr(expr string, params ...interface{}) Query {
	q.query.TableExpr(expr, params...)
	return q
}

func (q *QueryWrap) TableModel() orm.TableModel {
	return q.query.TableModel()
}

func (q *QueryWrap) Union(other Query) Query {
	q.query.Union(other.GetQuery())
	return q
}

func (q *QueryWrap) UnionAll(other Query) Query {
	q.query.UnionAll(other.GetQuery())
	return q
}

func (q *QueryWrap) Update(scan ...interface{}) (orm.Result, error) {
	return q.query.Update(scan...)
}

func (q *QueryWrap) UpdateNotZero(scan ...interface{}) (orm.Result, error) {
	return q.query.UpdateNotZero(scan...)
}

func (q *QueryWrap) Value(column string, value string, params ...interface{}) Query {
	q.query.Value(column, value, params...)
	return q
}

func (q *QueryWrap) Where(condition string, params ...interface{}) Query {
	q.query.Where(condition, params...)
	return q
}

func (q *QueryWrap) WhereGroup(fn func(Query) (Query, error)) Query {
	q.query.WhereGroup(func(pgquery *orm.Query) (*orm.Query, error) {
		query := NewQuery(pgquery)
		_, err := fn(query)
		return query.GetQuery(), err
	})
	return q
}

func (q *QueryWrap) WhereIn(where string, slice interface{}) Query {
	q.query.WhereIn(where, slice)
	return q
}

func (q *QueryWrap) WhereInMulti(where string, values ...interface{}) Query {
	q.query.WhereInMulti(where, values...)
	return q
}

func (q *QueryWrap) WhereNotGroup(fn func(Query) (Query, error)) Query {
	q.query.WhereNotGroup(func(pgquery *orm.Query) (*orm.Query, error) {
		query := NewQuery(pgquery)
		_, err := fn(query)
		return query.GetQuery(), err
	})
	return q
}

func (q *QueryWrap) WhereOr(condition string, params ...interface{}) Query {
	q.query.WhereOr(condition, params...)
	return q
}

func (q *QueryWrap) WhereOrGroup(fn func(Query) (Query, error)) Query {
	q.query.WhereOrGroup(func(pgquery *orm.Query) (*orm.Query, error) {
		query := NewQuery(pgquery)
		_, err := fn(query)
		return query.GetQuery(), err
	})
	return q
}

func (q *QueryWrap) WhereOrNotGroup(fn func(Query) (Query, error)) Query {
	q.query.WhereOrNotGroup(func(pgquery *orm.Query) (*orm.Query, error) {
		query := NewQuery(pgquery)
		_, err := fn(query)
		return query.GetQuery(), err
	})
	return q
}

func (q *QueryWrap) WherePK() Query {
	q.query.WherePK()
	return q
}

func (q *QueryWrap) With(name string, subq Query) Query {
	q.query.With(name, subq.GetQuery())
	return q
}

func (q *QueryWrap) WithDelete(name string, subq Query) Query {
	q.query.WithDelete(name, subq.GetQuery())
	return q
}

func (q *QueryWrap) WithInsert(name string, subq Query) Query {
	q.query.WithInsert(name, subq.GetQuery())
	return q
}

func (q *QueryWrap) WithUpdate(name string, subq Query) Query {
	q.query.WithUpdate(name, subq.GetQuery())
	return q
}

func (q *QueryWrap) WrapWith(name string) Query {
	q.query.WrapWith(name)
	return q
}
