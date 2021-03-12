package query

import "time"

type Schema struct {
	Id        int64     `pg:"id,pk"`
	CreatedAt time.Time `pg:"created_at,default:(now() at time zone 'utc')"`
	UpdatedAt time.Time `pg:"updated_at,default:(now() at time zone 'utc')"`
	tableName struct{}  `pg:"oto.book,alias:book,discard_unknown_columns"`
	Name      string    `pg:"name"`
}
