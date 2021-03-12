package gopgmock

import (
	"context"
	"testing"
	"time"

	pg2 "go-pg-utf/pg"

	"github.com/stretchr/testify/assert"
)

type Schema struct {
	Id        int64     `pg:"id,pk"`
	CreatedAt time.Time `pg:"created_at,default:(now() at time zone 'utc')"`
	UpdatedAt time.Time `pg:"updated_at,default:(now() at time zone 'utc')"`
	tableName struct{}  `pg:"oto.book,alias:book,discard_unknown_columns"`
	Name      string    `pg:"name"`
}

func TestMockDBWrap_Begin(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("begin")
	mock.Returns(NewResult(0, 0, nil), nil)
	tx, err := db.Begin()
	assert.Nil(t, err)
	assert.NotNil(t, tx)
}

func TestMockDBWrap_Exec(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id=1")
	mock.Returns(NewResult(1, 1, &Schema{Id: 1}), nil)
	ret, err := db.Exec("select id from oto.book as book where id=?", 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, ret.RowsAffected())
	assert.Equal(t, 1, ret.RowsReturned())
}

func TestMockDBWrap_ExecContext(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id=1")
	mock.Returns(NewResult(1, 1, &Schema{Id: 1}), nil)
	ret, err := db.ExecContext(context.Background(), "select id from oto.book as book where id=?", 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, ret.RowsAffected())
	assert.Equal(t, 1, ret.RowsReturned())
}

func TestMockDBWrap_ExecOne(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id=1")
	mock.Returns(NewResult(1, 1, &Schema{Id: 1}), nil)
	ret, err := db.ExecOne("select id from oto.book as book where id=?", 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, ret.RowsAffected())
	assert.Equal(t, 1, ret.RowsReturned())
}

func TestMockDBWrap_ExecOneContext(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id=1")
	mock.Returns(NewResult(1, 1, &Schema{Id: 1}), nil)
	ret, err := db.ExecOneContext(context.Background(), "select id from oto.book as book where id=?", 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, ret.RowsAffected())
	assert.Equal(t, 1, ret.RowsReturned())
}

type ID struct {
	ID int64
}

func TestMockDBWrap_Query(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id = 1")
	mock.Returns(NewResult(1, 1, ID{ID: 1}), nil)
	var id int64
	_, err := db.Query(&id, "select id from oto.book as book where id = ?", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
}

func TestMockDBWrap_QueryContext(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id = 1")
	mock.Returns(NewResult(1, 1, ID{ID: 1}), nil)
	var id int64
	_, err := db.QueryContext(context.Background(), &id, "select id from oto.book as book where id = ?", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
}

func TestMockDBWrap_QueryOne(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id = 1")
	mock.Returns(NewResult(1, 1, ID{ID: 1}), nil)
	var id int64
	_, err := db.QueryOne(&id, "select id from oto.book as book where id = ?", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
}

func TestMockDBWrap_QueryOneContext(t *testing.T) {
	db, mock := StartMockDBWrap()
	mock.ExpectQuery("select id from oto.book as book where id = 1")
	mock.Returns(NewResult(1, 1, ID{ID: 1}), nil)
	var id int64
	_, err := db.QueryOneContext(context.Background(), &id, "select id from oto.book as book where id = ?", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
}

func TestMockDBWrap_RunInTransaction(t *testing.T) {
	db, sqlMock := StartMockDB()
	sqlMock.ExpectQuery("begin")
	sqlMock.Returns(NewResult(0, 0, nil), nil)
	sqlMock.ExpectQuery("update oto.book as book set name='sdt' where book.id=1")
	sqlMock.Returns(NewResult(1, 1, Schema{Id: 1}), nil)
	sqlMock.ExpectQuery("commit")
	sqlMock.Returns(NewResult(0, 0, nil), nil)
	var s Schema
	err := db.RunInTransaction(context.TODO(), func(tx pg2.Tx) error {
		s.Id = 1
		s.Name = "sdt"
		_, err := tx.Model(&s).Column("name").WherePK().Update()
		return err
	})
	assert.Nil(t, err)
}
