package demo

// This example demonstrates
//   * embedded structs
//   * the use of pointers for optional items

//go:generate sqlgen -json -type demo.Hook -o hook_sql.go -list HookList -all -v .

type Hook struct {
	Id  uint64 `sql:"pk: true, auto: true"`
	Sha string
	Bounds
	Category   Category
	Created    bool
	Deleted    bool
	Forced     bool
	HeadCommit *Commit `sql:"name: head"`
}

type HookList []*Hook

type Bounds struct {
	After  string `sql:"size: 20"`
	Before string `sql:"size: 20"`
}

type Commit struct {
	ID        string `sql:"name: commit_id"`
	Message   string
	Timestamp string
	Author    *Author
	Committer *Author
}

type Author struct {
	Name     string `sql:"prefixed: true"`
	Email    Email  `sql:"prefixed: true"`
	Username string `sql:"prefixed: true"`
}

type Email string
