package demo

// This example demonstrates
//   * embedded structs
//   * the use of pointers for optional items

//go:generate sqlgen -type demo.Hook -o hook_sql.go -list HookList -v .

type Hook struct {
	Id  int64 `sql:"pk: true, auto: true"`
	Sha string
	Dates
	Category   Category
	Created    bool
	Deleted    bool
	Forced     bool
	HeadCommit *Commit `sql:"name: head"`
}

type HookList []*Hook

type Dates struct {
	After  string `sql:"size: 20"`
	Before string `sql:"size: 20"`
}

type Commit struct {
	ID        string
	Message   string
	Timestamp string
	Author    *Author
	Committer *Author
}

type Author struct {
	Name     string
	Email    string
	Username string
}
