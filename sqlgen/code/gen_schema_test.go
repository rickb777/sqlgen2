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
 id       bigint primary key,
 cat      int,
 username text,
 active   boolean,
 labels   text,
 fave     blob,
 avatar   blob,
 updated  blob
)
|

const sqlCreateXExampleTablePostgres = |
CREATE TABLE %s%s%s (
 id       bigserial primary key,
 cat      integer,
 username varchar(2048),
 active   boolean,
 labels   json,
 fave     byteaa,
 avatar   byteaa,
 updated  byteaa
)
|

const sqlCreateXExampleTableMysql = |
CREATE TABLE %s%s%s (
 id       bigint primary key auto_increment,
 cat      int,
 username varchar(2048),
 active   tinyint(1),
 labels   json,
 fave     mediumblob,
 avatar   mediumblob,
 updated  mediumblob
) ENGINE=InnoDB DEFAULT CHARSET=utf8
|

//--------------------------------------------------------------------------------

const sqlCreateXNameIdxIndex = |
CREATE INDEX %s%snameIdx ON %s%s (username)
|

//--------------------------------------------------------------------------------

const NumXExampleColumns = 8

const NumXExampleDataColumns = 7

const XExamplePk = "Id"

const XExampleDataColumnNames = "cat, username, active, labels, fave, avatar, updated"

//--------------------------------------------------------------------------------
`, "|", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
