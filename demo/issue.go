package demo

//go:generate sqlgen -type demo.Issue -o issue_sql.go -db postgres -v -z issue.go

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
