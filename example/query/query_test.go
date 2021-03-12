package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.xinghuolive.com/birds-backend/chameleon/store/postgres/gopgmock"
)

func TestQueryWrap_Clone(t *testing.T) {
	db, mock := gopgmock.StartMockDB()
	query := db.Model((*Schema)(nil))
	newQuery := query.Clone().ColumnExpr("id")
	query.Column("*")
	assert.NotEqual(t, query, newQuery)

	mock.ExpectQuery("select * from oto.book as book")
	mock.Returns(gopgmock.NewResult(1, 1, Schema{Name: "111"}), nil)
	mock.ExpectQuery("select id from oto.book as book")
	mock.Returns(gopgmock.NewResult(1, 1, Schema{Id: 111}), nil)
	var q1, q2 Schema
	err1 := query.Select(&q1)
	err2 := newQuery.Select(&q2)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, "111", q1.Name)
	assert.Equal(t, int64(111), q2.Id)
	str1, _ := query.AppendQuery(db.Formatter(), []byte{})
	str2, _ := newQuery.AppendQuery(db.Formatter(), []byte{})
	assert.NotEqual(t, string(str1), string(str2))
}

func TestQueryWrap_Column(t *testing.T) {
	db, mock := gopgmock.StartMockDB()
	query := db.Model((*Schema)(nil))
	query.Column("id")

	mock.ExpectQuery("select id from oto.book as book")
	mock.Returns(gopgmock.NewResult(1, 1, Schema{Id: 111}), nil)
	var q1 Schema
	err1 := query.Select(&q1)
	assert.Nil(t, err1)
	assert.Equal(t, int64(111), q1.Id)
}

func TestQueryWrap_ColumnExpr(t *testing.T) {
	db, mock := gopgmock.StartMockDB()
	query := db.Model((*Schema)(nil))
	query.ColumnExpr("id")

	mock.ExpectQuery("select id from oto.book as book")
	mock.Returns(gopgmock.NewResult(1, 1, Schema{Id: 111}), nil)
	var q1 Schema
	err1 := query.Select(&q1)
	assert.Nil(t, err1)
	assert.Equal(t, int64(111), q1.Id)
}

func TestQueryWrap_Select(t *testing.T) {
	db, mock := gopgmock.StartMockDB()
	var schema = Schema{Id: 111, Name: "111"}
	query := db.Model(&schema)
	query.ColumnExpr("id").ColumnExpr("name").WherePK()

	mock.ExpectQuery("select id, name from oto.book as book where book.id = 111")
	mock.Returns(gopgmock.NewResult(1, 1, Schema{Id: 111, Name: "111"}), nil)
	var q1 Schema
	err1 := query.Select(&q1)
	assert.Nil(t, err1)
	assert.Equal(t, int64(111), q1.Id)
	assert.Equal(t, "111", q1.Name)
}

func TestQueryWrap_Update(t *testing.T) {
	db, mock := gopgmock.StartMockDB()
	var schema = Schema{Id: 111, Name: "222"}
	query := db.Model(&schema)
	query.Column("name").WherePK()

	mock.ExpectQuery("update oto.book as book set name='222' where book.id = 111")
	mock.Returns(gopgmock.NewResult(1, 1, Schema{Id: 111, Name: "222"}), nil)
	_, err1 := query.Update()
	assert.Nil(t, err1)
	assert.Equal(t, int64(111), schema.Id)
	assert.Equal(t, "222", schema.Name)
}
