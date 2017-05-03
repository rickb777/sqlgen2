package demo

//go:generate sqlgen -type demo.Hook -o hook_sql.go -db mysql -v -z hook.go

type Hook struct {
	Id  int64 `sql:"pk: true, auto: true"`
	Sha string
	Dates
	//Category   Category
	Created    bool
	Deleted    bool
	Forced     bool
	HeadCommit *Commit `sql:"name: head"`
}

type Dates struct {
	After  string
	Before string
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

type Category int

const (
	Alpha Category = iota
	Beta
	Gamma
	Delta
)
