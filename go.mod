module github.com/rickb777/sqlgen2

require (
	github.com/acsellers/inflections v0.0.0-20141027155830-cb98bfe9e3ee
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.16.1
	github.com/kortschak/utter v1.5.0
	github.com/lib/pq v1.10.5
	github.com/markbates/inflect v1.0.4
	github.com/mattn/go-sqlite3 v1.14.12
	github.com/onsi/gomega v1.19.0
	github.com/pkg/errors v0.9.1
	github.com/rickb777/collection v1.0.1
	github.com/rickb777/date v1.17.0
	github.com/rickb777/filemod v0.9.1
	github.com/rickb777/sqlapi v0.59.0
	github.com/rickb777/where v0.18.0
	github.com/spf13/cast v1.5.0
	golang.org/x/tools v0.1.10
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/gobuffalo/envy v1.10.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/rickb777/plural v1.4.1 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20220511200225-c6db032c6c88 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

//replace github.com/rickb777/sqlapi => ../sqlapi

go 1.17
