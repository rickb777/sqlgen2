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

const NumXExampleColumns = 17

const NumXExampleDataColumns = 16

const XExamplePk = "Id"

const XExampleDataColumnNames = "cat, username, mobile, qual, diff, age, bmi, active, labels, fave, avatar, foo1, foo2, bar1, bar2, updated"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl XExampleTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.createTableSql(ifNotExists))
}

func (tbl XExampleTable) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.dialect {
	case schema.Sqlite: stmt = sqlCreateXExampleTableSqlite
    case schema.Postgres: stmt = sqlCreateXExampleTablePostgres
    case schema.Mysql: stmt = sqlCreateXExampleTableMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.name)
	return query
}

func (tbl XExampleTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl XExampleTable) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.dropTableSql(ifExists))
}

func (tbl XExampleTable) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", extra, tbl.name)
	return query
}

const sqlCreateXExampleTableSqlite = |
CREATE TABLE %s%s (
 id       integer primary key autoincrement,
 cat      int,
 username text,
 mobile   text default null,
 qual     text default null,
 diff     int default null,
 age      int unsigned default null,
 bmi      float default null,
 active   boolean,
 labels   text,
 fave     text,
 avatar   blob,
 foo1     text,
 foo2     text default null,
 bar1     text,
 bar2     text default null,
 updated  text
)
|

const sqlCreateXExampleTablePostgres = |
CREATE TABLE %s%s (
 id       bigserial primary key,
 cat      int,
 username varchar(2048),
 mobile   varchar(255) default null,
 qual     varchar(255) default null,
 diff     int default null,
 age      int unsigned default null,
 bmi      float default null,
 active   boolean,
 labels   json,
 fave     json,
 avatar   byteaa,
 foo1     varchar(255),
 foo2     varchar(255) default null,
 bar1     varchar(255),
 bar2     varchar(255) default null,
 updated  varchar(100)
)
|

const sqlCreateXExampleTableMysql = |
CREATE TABLE %s%s (
 id       bigint primary key auto_increment,
 cat      int,
 username varchar(2048),
 mobile   varchar(255) default null,
 qual     varchar(255) default null,
 diff     int default null,
 age      int unsigned default null,
 bmi      float default null,
 active   tinyint(1),
 labels   json,
 fave     json,
 avatar   mediumblob,
 foo1     varchar(255),
 foo2     varchar(255) default null,
 bar1     varchar(255),
 bar2     varchar(255) default null,
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

	err = tbl.CreateCatIdxIndex(ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateNameIdxIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateCatIdxIndex creates the catIdx index.
func (tbl XExampleTable) CreateCatIdxIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropCatIdxIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createXCatIdxIndexSql(ine))
	return err
}

func (tbl XExampleTable) createXCatIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE INDEX %s%scatIdx ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlXCatIdxIndexColumns)
}

// DropCatIdxIndex drops the catIdx index.
func (tbl XExampleTable) DropCatIdxIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropXCatIdxIndexSql(ifExists))
	return err
}

func (tbl XExampleTable) dropXCatIdxIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%scatIdx%s", ie, indexPrefix, onTbl)
}

// CreateNameIdxIndex creates the nameIdx index.
func (tbl XExampleTable) CreateNameIdxIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropNameIdxIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createXNameIdxIndexSql(ine))
	return err
}

func (tbl XExampleTable) createXNameIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%snameIdx ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlXNameIdxIndexColumns)
}

// DropNameIdxIndex drops the nameIdx index.
func (tbl XExampleTable) DropNameIdxIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropXNameIdxIndexSql(ifExists))
	return err
}

func (tbl XExampleTable) dropXNameIdxIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%snameIdx%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the Example table.
func (tbl XExampleTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropCatIdxIndex(ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropNameIdxIndex(ifExist)
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

const sqlXCatIdxIndexColumns = "cat"

const sqlXNameIdxIndexColumns = "username"

//--------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl XExampleTable) Truncate(force bool) (err error) {
	for _, query := range tbl.dialect.TruncateDDL(tbl.Name().String(), force) {
		_, err = tbl.Exec(nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}
`, "|", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
