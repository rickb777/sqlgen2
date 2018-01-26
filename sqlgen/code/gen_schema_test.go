package code

import (
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
	"strings"
)

func TestWritePrimaryDeclarations(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WritePrimaryDeclarations(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

const NumXExampleColumns = 17

const NumXExampleDataColumns = 16

const XExampleColumnNames = "id,cat,username,mobile,qual,diff,age,bmi,active,labels,fave,avatar,foo1,foo2,bar1,bar2,updated"

const XExampleDataColumnNames = "cat,username,mobile,qual,diff,age,bmi,active,labels,fave,avatar,foo1,foo2,bar1,bar2,updated"

const XExamplePk = "id"
`, "¬", "`", -1)

	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
func TestWriteSchemaDeclarations(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSchemaDeclarations(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

const sqlXExampleTableCreateColumnsSqlite = "\n"+
" ¬id¬       integer primary key autoincrement,\n"+
" ¬cat¬      int,\n"+
" ¬username¬ text,\n"+
" ¬mobile¬   text default null,\n"+
" ¬qual¬     text default null,\n"+
" ¬diff¬     int default null,\n"+
" ¬age¬      int unsigned default null,\n"+
" ¬bmi¬      float default null,\n"+
" ¬active¬   boolean,\n"+
" ¬labels¬   text,\n"+
" ¬fave¬     text,\n"+
" ¬avatar¬   blob,\n"+
" ¬foo1¬     text,\n"+
" ¬foo2¬     text default null,\n"+
" ¬bar1¬     text,\n"+
" ¬bar2¬     text default null,\n"+
" ¬updated¬  text"

const sqlXExampleTableCreateColumnsMysql = "\n"+
" ¬id¬       bigint primary key auto_increment,\n"+
" ¬cat¬      int,\n"+
" ¬username¬ varchar(2048),\n"+
" ¬mobile¬   varchar(255) default null,\n"+
" ¬qual¬     varchar(255) default null,\n"+
" ¬diff¬     int default null,\n"+
" ¬age¬      int unsigned default null,\n"+
" ¬bmi¬      float default null,\n"+
" ¬active¬   tinyint(1),\n"+
" ¬labels¬   json,\n"+
" ¬fave¬     json,\n"+
" ¬avatar¬   mediumblob,\n"+
" ¬foo1¬     varchar(255),\n"+
" ¬foo2¬     varchar(255) default null,\n"+
" ¬bar1¬     varchar(255),\n"+
" ¬bar2¬     varchar(255) default null,\n"+
" ¬updated¬  varchar(100)"

const sqlXExampleTableCreateColumnsPostgres = ¬
 "id"       bigserial primary key,
 "cat"      int,
 "username" varchar(2048),
 "mobile"   varchar(255) default null,
 "qual"     varchar(255) default null,
 "diff"     int default null,
 "age"      int unsigned default null,
 "bmi"      float default null,
 "active"   boolean,
 "labels"   json,
 "fave"     json,
 "avatar"   byteaa,
 "foo1"     varchar(255),
 "foo2"     varchar(255) default null,
 "bar1"     varchar(255),
 "bar2"     varchar(255) default null,
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
