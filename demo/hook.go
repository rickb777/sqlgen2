package demo

//go:generate sqlgen -type demo.Hook -o hook_sql.go -db mysql -v -z hook.go

type Hook struct {
	Id         int64 `sql:"pk: true, auto: true"`
	Sha        string
	After      string
	Before     string
	Created    bool
	Deleted    bool
	Forced     bool
	HeadCommit *Commit `sql:"name: head"`
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
