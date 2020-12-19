module github.com/rickb777/sqlgen2

require (
	github.com/acsellers/inflections v0.0.0-20141027155830-cb98bfe9e3ee
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.10.0
	github.com/kortschak/utter v1.0.1
	github.com/lib/pq v1.9.0
	github.com/markbates/inflect v1.0.4
	github.com/mattn/go-sqlite3 v1.14.5
	github.com/onsi/gomega v1.10.4
	github.com/pkg/errors v0.9.1
	github.com/rickb777/collection v0.7.0
	github.com/rickb777/date v1.14.3
	github.com/rickb777/filemod v0.9.1
	github.com/rickb777/sqlapi v0.58.0
	github.com/rickb777/where v0.11.0
	github.com/spf13/cast v1.3.1
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/tools v0.0.0-20201208233053-a543418bbed2
	gopkg.in/yaml.v2 v2.4.0
)

//replace github.com/rickb777/sqlapi => ../sqlapi

go 1.15
