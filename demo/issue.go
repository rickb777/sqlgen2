package demo

//go:generate sqlgen -file issue.go -type Issue -o issue_sql.go -db postgres

type Issue struct {
	Id       int64 `sql:"pk: true, auto: true"`
	Number   int
	Title    string   `sql:"size: 512"`
	Body     string   `sql:"size: 2048, name: bigbody"`
	Assignee string   `sql:"index: issue_assignee"`
	State    string   `sql:"size: 50"`
	Labels   []string `sql:"encode: json"`

	locked bool `sql:"-"`
}
