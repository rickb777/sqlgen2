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

const NumXExampleColumns = 8

const NumXExampleDataColumns = 7

const XExamplePk = "Id"

const XExampleDataColumnNames = "cat, username, active, labels, fave, avatar, updated"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl XExampleTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl XExampleTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case schema.Sqlite: stmt = sqlCreateXExampleTableSqlite
    case schema.Postgres: stmt = sqlCreateXExampleTablePostgres
    case schema.Mysql: stmt = sqlCreateXExampleTableMysql
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
 id       integer primary key autoincrement,
 cat      int,
 username text,
 active   boolean,
 labels   text,
 fave     text,
 avatar   blob,
 updated  text
)
|

const sqlCreateXExampleTablePostgres = |
CREATE TABLE %s%s%s (
 id       bigserial primary key,
 cat      int,
 username varchar(2048),
 active   boolean,
 labels   json,
 fave     json,
 avatar   byteaa,
 updated  varchar(100)
)
|

const sqlCreateXExampleTableMysql = |
CREATE TABLE %s%s%s (
 id       bigint primary key auto_increment,
 cat      int,
 username varchar(2048),
 active   tinyint(1),
 labels   json,
 fave     json,
 avatar   mediumblob,
 updated  varchar(100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
|

//--------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl XExampleTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the Example table.
func (tbl XExampleTable) CreateIndexes(ifNotExist bool) (err error) {
	ine := tbl.ternary(ifNotExist && tbl.Dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect == schema.Mysql {
		tbl.DropIndexes(false)
		ine = ""
	}

	_, err = tbl.Exec(tbl.createXCatIdxIndexSql(ine))
	if err != nil {
		return err
	}

	_, err = tbl.Exec(tbl.createXNameIdxIndexSql(ine))
	if err != nil {
		return err
	}

	return nil
}

func (tbl XExampleTable) createXCatIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE INDEX %s%scatIdx ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlXCatIdxIndexColumns)
}

func (tbl XExampleTable) dropXCatIdxIndexSql(ifExists, onTbl string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%scatIdx%s", ifExists, indexPrefix, onTbl)
}

func (tbl XExampleTable) createXNameIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%snameIdx ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlXNameIdxIndexColumns)
}

func (tbl XExampleTable) dropXNameIdxIndexSql(ifExists, onTbl string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%snameIdx%s", ifExists, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the Example table.
func (tbl XExampleTable) DropIndexes(ifExist bool) (err error) {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExist && tbl.Dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.Prefix, tbl.Name), "")

	_, err = tbl.Exec(tbl.dropXCatIdxIndexSql(ie, onTbl))
	if err != nil {
		return err
	}

	_, err = tbl.Exec(tbl.dropXNameIdxIndexSql(ie, onTbl))
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

const sqlXCatIdxIndexColumns = "cat"

const sqlXNameIdxIndexColumns = "username"
`, "|", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
