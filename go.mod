module github.com/rickb777/sqlgen2

require (
	github.com/acsellers/inflections v0.0.0-20141027155830-cb98bfe9e3ee
	github.com/go-sql-driver/mysql v1.7.1
	github.com/jackc/pgx/v5 v5.4.3
	github.com/kortschak/utter v1.5.0
	github.com/lib/pq v1.10.9
	github.com/markbates/inflect v1.0.4
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/onsi/gomega v1.27.10
	github.com/pkg/errors v0.9.1
	github.com/rickb777/collection v1.0.1
	github.com/rickb777/date v1.20.2
	github.com/rickb777/filemod v0.9.1
	github.com/rickb777/sqlapi v0.61.0
	github.com/rickb777/where v0.18.0
	github.com/spf13/cast v1.5.1
	golang.org/x/tools v0.12.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/gobuffalo/envy v1.10.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/rickb777/plural v1.4.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace github.com/rickb777/sqlapi => ../sqlapi

go 1.19
