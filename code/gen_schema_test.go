package code

import (
	"bytes"
	"github.com/rickb777/sqlgen2/parse/exit"
	"strings"
	"testing"
)

//TODO
func xTestWritePrimaryDeclarations(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WritePrimaryDeclarations(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// NumXExampleColumns is the total number of columns in XExample.
const NumXExampleColumns = 17

// NumXExampleDataColumns is the number of columns in XExample not including the auto-increment key.
const NumXExampleDataColumns = 16

// XExampleColumnNames is the list of columns in XExample.
const XExampleColumnNames = "id,cat,username,mobile,qual,diff,age,bmi,active,labels,fave,avatar,foo1,foo2,foo3,bar1,updated"

// XExampleDataColumnNames is the list of data columns in XExample.
const XExampleDataColumnNames = "cat,username,mobile,qual,diff,age,bmi,active,labels,fave,avatar,foo1,foo2,foo3,bar1,updated"
`, "¬", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//TODO
func xTestWriteSchemaDeclarations(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSchemaDeclarations(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

const sqlXExampleTableCreateColumnsSqlite = "\n"+
" ¬id¬       integer not null primary key autoincrement,\n"+
" ¬cat¬      int not null,\n"+
" ¬username¬ text not null default 'anon',\n"+
" ¬mobile¬   text default null,\n"+
" ¬qual¬     text default null,\n"+
" ¬diff¬     int default null,\n"+
" ¬age¬      int unsigned default null,\n"+
" ¬bmi¬      float default null,\n"+
" ¬active¬   boolean not null,\n"+
" ¬labels¬   text,\n"+
" ¬fave¬     text,\n"+
" ¬avatar¬   blob not null,\n"+
" ¬foo1¬     text not null,\n"+
" ¬foo2¬     text default null,\n"+
" ¬foo3¬     int default null,\n"+
" ¬bar1¬     text not null,\n"+
" ¬updated¬  text"

const sqlXExampleTableCreateColumnsMysql = "\n"+
" ¬id¬       bigint not null primary key auto_increment,\n"+
" ¬cat¬      int not null,\n"+
" ¬username¬ varchar(2048) not null default 'anon',\n"+
" ¬mobile¬   varchar(255) default null,\n"+
" ¬qual¬     varchar(255) default null,\n"+
" ¬diff¬     int default null,\n"+
" ¬age¬      int unsigned default null,\n"+
" ¬bmi¬      float default null,\n"+
" ¬active¬   tinyint(1) not null,\n"+
" ¬labels¬   json,\n"+
" ¬fave¬     json,\n"+
" ¬avatar¬   mediumblob not null,\n"+
" ¬foo1¬     varchar(255) not null,\n"+
" ¬foo2¬     varchar(255) default null,\n"+
" ¬foo3¬     int default null,\n"+
" ¬bar1¬     varchar(255) not null,\n"+
" ¬updated¬  varchar(100)"

const sqlXExampleTableCreateColumnsPostgres = ¬
 "id"       bigserial not null primary key,
 "cat"      integer not null,
 "username" varchar(2048) not null default 'anon',
 "mobile"   varchar(255) default null,
 "qual"     varchar(255) default null,
 "diff"     integer default null,
 "age"      bigint default null,
 "bmi"      real default null,
 "active"   boolean not null,
 "labels"   json,
 "fave"     json,
 "avatar"   bytea not null,
 "foo1"     varchar(255) not null,
 "foo2"     varchar(255) default null,
 "foo3"     integer default null,
 "bar1"     varchar(255) not null,
 "updated"  varchar(100)¬

const sqlXExampleTableCreateColumnsPgx = ¬
 "id"       bigserial not null primary key,
 "cat"      integer not null,
 "username" varchar(2048) not null default 'anon',
 "mobile"   varchar(255) default null,
 "qual"     varchar(255) default null,
 "diff"     integer default null,
 "age"      bigint default null,
 "bmi"      real default null,
 "active"   boolean not null,
 "labels"   json,
 "fave"     json,
 "avatar"   bytea not null,
 "foo1"     varchar(255) not null,
 "foo2"     varchar(255) default null,
 "foo3"     integer default null,
 "bar1"     varchar(255) not null,
 "updated"  varchar(100)¬

const sqlConstrainXExampleTable = ¬
¬

//--------------------------------------------------------------------------------

const sqlXCatIdxIndexColumns = "cat"

const sqlXNameIdxIndexColumns = "username"
`, "¬", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
