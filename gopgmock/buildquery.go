package gopgmock

type buildQuery struct {
	query  string
	params []interface{}
	result *OrmResult
	err    error
}
