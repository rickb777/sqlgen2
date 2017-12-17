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
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSchema(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl XExampleTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl XExampleTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case sqlgen2.Sqlite: stmt = sqlCreateXExampleTableSqlite
    case sqlgen2.Postgres: stmt = sqlCreateXExampleTablePostgres
    case sqlgen2.Mysql: stmt = sqlCreateXExampleTableMysql
    }
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl XExampleTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

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

// CreateIndexes executes queries that create the indexes needed by the Example table.
func (tbl XExampleTable) CreateIndexes(ifNotExist bool) (err error) {
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	_, err = tbl.Exec(tbl.createXNameIdxIndexSql(extra))
	if err != nil {
		return err
	}

	return nil
}


func (tbl XExampleTable) createXNameIdxIndexSql(ifNotExist string) string {
	indexPrefix := tbl.Prefix
	if strings.HasSuffix(indexPrefix, ".") {
		indexPrefix = tbl.Prefix[0:len(indexPrefix)-1]
	}
	return fmt.Sprintf(sqlCreateXNameIdxIndex, ifNotExist, indexPrefix, tbl.Prefix, tbl.Name)
}


// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl XExampleTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}
	return tbl.CreateIndexes(ifNotExist)
}

//--------------------------------------------------------------------------------

const sqlCreateXNameIdxIndex = |
CREATE INDEX %s%snameIdx ON %s%s (username)
|

//--------------------------------------------------------------------------------

const NumXExampleColumns = 8

const NumXExampleDataColumns = 7

const XExamplePk = "Id"

const XExampleDataColumnNames = "cat, username, active, labels, fave, avatar, updated"
`, "|", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
