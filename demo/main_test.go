package demo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/log/testingadapter"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/driver"
	"github.com/rickb777/where/quote"
	"io"
	"log"
	"math/big"
	"os"
	"testing"
)

// Environment:
// GO_DRIVER  - the driver (sqlite3, mysql, postgres, pgx)
// GO_DSN     - the database DSN
// GO_VERBOSE - true for query logging
// PGQUOTE    - the identifier quoter (ansi, mysql, none)

func TestMain(m *testing.M) {
	db = newDatabase()
	code := m.Run()
	cleanup(db)
	os.Exit(code)
}

//-------------------------------------------------------------------------------------------------

var verbose = false
var db sqlapi.SqlDB

func connect() (*sql.DB, driver.Dialect) {
	dbDriver, ok := os.LookupEnv("GO_DRIVER")
	if !ok {
		dbDriver = "sqlite3"
	}

	di := driver.PickDialect(dbDriver) //.WithQuoter(dialect.NoQuoter)
	quoter, ok := os.LookupEnv("PGQUOTE")
	if ok {
		q := quote.PickQuoter(quoter)
		if q == nil {
			log.Fatalf("Warning: unrecognised quoter %q.\n", quoter)
		}
		di = di.WithQuoter(q)
	}

	dsn, ok := os.LookupEnv("GO_DSN")
	if !ok {
		dsn = "file::memory:?mode=memory&cache=shared"
	}

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatalf("Error: Unable to connect to %s (%v); test is only partially complete.\n\n", dbDriver, err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error: Unable to ping %s (%v); test is only partially complete.\n\n", dbDriver, err)
	}

	fmt.Printf("Successfully connected to %s.\n", dbDriver)
	return db, di
}

func newDatabase() sqlapi.SqlDB {
	db, di := connect()
	lgr := testingadapter.NewLogger(simpleLogger{})
	return sqlapi.WrapDB(db, di, sqlapi.NewLogger(lgr))
}

func cleanup(db sqlapi.Execer) {
	if db != nil {
		if c, ok := db.(io.Closer); ok {
			c.Close()
		}
		os.Remove("test.db")
	}
}

func user(i int) *User {
	return &User{
		Name:         fmt.Sprintf("user%02d", i),
		EmailAddress: fmt.Sprintf("foo%d@x.z", i),
		Active:       true,
		Fave:         big.NewInt(int64(i)),
		Numbers: Numbers{
			I8:  int8(i * 5),
			U8:  uint8(i * 6),
			I16: int16(i * 10),
			U16: uint16(i * 11),
			I32: int32(i * 100),
			U32: uint32(i * 101),
			I64: int64(i * 200),
			U64: uint64(i * 201),
			F32: float32(i * 300),
			F64: float64(i * 301),
		},
	}
}

//-------------------------------------------------------------------------------------------------

type simpleLogger struct{}

func (l simpleLogger) Log(args ...interface{}) {
	if testing.Verbose() {
		log.Println(args...)
	}
}
