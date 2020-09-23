package demopgx

import (
	"flag"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/testingadapter"
	"github.com/rickb777/sqlapi/pgxapi"
	"io"
	"log"
	"math/big"
	"os"
	"testing"
)

// Environment:
// GO_DRIVER  - the driver (sqlite3, mysql, postgres, pgx)
// GO_DSN     - the database DSN
// PGQUOTE    - the identifier quoter (ansi, mysql, none)

func TestMain(m *testing.M) {
	flag.Parse()
	db = connect()
	code := m.Run()
	cleanup(db)
	os.Exit(code)
}

//-------------------------------------------------------------------------------------------------

var db pgxapi.SqlDB

func connect() pgxapi.SqlDB {
	lgr := testingadapter.NewLogger(simpleLogger{})
	var lvl pgx.LogLevel = pgx.LogLevelInfo
	if !testing.Verbose() {
		lvl = pgx.LogLevelWarn
	}
	d, err := pgxapi.ConnectEnv(lgr, lvl)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func cleanup(d pgxapi.Execer) {
	if d != nil {
		if c, ok := d.(io.Closer); ok {
			c.Close()
		}
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

type simpleLogger struct{}

func (l simpleLogger) Log(args ...interface{}) {
	log.Println(args...)
}
