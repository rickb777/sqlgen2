package code

import (
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
	"strings"
)

func TestWriteSchema(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	buf := &bytes.Buffer{}

	WriteSchema(buf, view, fixtureTable())

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

const sqlCreateXExampleTableSqlite = |
CREATE TABLE %s%s%s (
 id      integer primary key autoincrement,
 cat     integer,
 name    text,
 active  boolean,
 labels  blob,
 fave    blob,
 updated blob
)
|

const sqlCreateXExampleTablePostgres = |
CREATE TABLE %s%s%s (
 id      bigserial primary key ,
 cat     integer,
 name    varchar(2048),
 active  boolean,
 labels  byteaa,
 fave    byteaa,
 updated byteaa
)
|

const sqlCreateXExampleTableMysql = |
CREATE TABLE %s%s%s (
 id      bigint primary key auto_increment,
 cat     int,
 name    varchar(2048),
 active  tinyint(1),
 labels  mediumblob,
 fave    mediumblob,
 updated mediumblob
) ENGINE=InnoDB DEFAULT CHARSET=utf8
|

//--------------------------------------------------------------------------------

const sqlCreateXNameIdxIndex = |
CREATE INDEX %s%snameIdx ON %s%s (name)
|

//--------------------------------------------------------------------------------

const NumXExampleColumns = 7

const NumXExampleDataColumns = 6

const XExamplePk = "Id"

const XExampleDataColumnNames = "cat, name, active, labels, fave, updated"

//--------------------------------------------------------------------------------
`, "|", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
