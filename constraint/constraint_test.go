package constraint_test

import (
	"database/sql"
	"testing"
	. "github.com/onsi/gomega"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/schema"
	"log"
	"os"
	"strings"
	"fmt"
	"github.com/rickb777/sqlgen2/vanilla"
	"github.com/rickb777/sqlgen2/constraint"
)

var db *sql.DB
var dialect schema.Dialect = schema.Sqlite

func connect() {
	dbDriver, ok := os.LookupEnv("GO_DRIVER")
	if !ok {
		dbDriver = "sqlite3"
	}
	dialect = schema.PickDialect(dbDriver)
	dsn, ok := os.LookupEnv("GO_DSN")
	if !ok {
		dsn = ":memory:"
	}
	conn, err := sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err)
	}
	db = conn
}

func cleanup() {
	if db != nil {
		db.Close()
		db = nil
	}
}

func TestIdsUsedAsForeignKeys(t *testing.T) {
	RegisterTestingT(t)
	connect()
	defer cleanup()

	fkc0 := constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "cascade", "cascade"}

	d := sqlgen2.NewDatabase(db, dialect, nil, nil)
	if testing.Verbose() {
		lgr := log.New(os.Stderr, "", log.LstdFlags)
		d = sqlgen2.NewDatabase(db, dialect, lgr, nil)
	}

	persons := vanilla.NewRecordTable("persons", d).WithPrefix("pfx_").WithConstraint(fkc0)

	setupSql := strings.Replace(createTables, "¬", "`", -1)
	_, err := d.Exec(setupSql)
	Ω(err).Should(BeNil())

	aid1 := insertOne(d, address1)
	aid2 := insertOne(d, address2)
	aid3 := insertOne(d, address3)
	aid4 := insertOne(d, address4)

	insertOne(d, fmt.Sprintf(person1a, aid1))
	insertOne(d, fmt.Sprintf(person1b, aid1))
	insertOne(d, fmt.Sprintf(person2a, aid2))

	fkc := persons.Constraints().FkConstraints()[0]

	m1, err := fkc.RelationshipWith(persons.Name()).IdsUsedAsForeignKeys(persons)

	Ω(err).Should(BeNil())
	Ω(m1).Should(HaveLen(2))
	Ω(m1.Contains(aid1)).Should(BeTrue())
	Ω(m1.Contains(aid2)).Should(BeTrue())

	m2, err := fkc.RelationshipWith(persons.Name()).IdsUnusedAsForeignKeys(persons)

	Ω(err).Should(BeNil())
	Ω(m2).Should(HaveLen(2))
	Ω(m2.Contains(aid3)).Should(BeTrue())
	Ω(m2.Contains(aid4)).Should(BeTrue())
}

func insertOne(d *sqlgen2.Database, query string) int64 {
	fmt.Fprintf(os.Stderr, "%s\n", query)
	r, err := d.Exec(query)
	Ω(err).Should(BeNil())

	id, err := r.LastInsertId()
	Ω(err).Should(BeNil())
	return id
}

//-------------------------------------------------------------------------------------------------

const createTables = `
CREATE TABLE IF NOT EXISTS pfx_addresses (
 ¬id¬        integer primary key autoincrement,
 ¬lines¬     text,
 ¬postcode¬  text
);

CREATE TABLE IF NOT EXISTS pfx_persons (
 ¬uid¬       integer primary key autoincrement,
 ¬name¬      text,
 ¬addressid¬ integer default null
);

DELETE FROM pfx_persons;
DELETE FROM pfx_addresses;
`

const address1 = `INSERT INTO pfx_addresses (lines, postcode) VALUES ('Laurel Cottage', 'FX1 1AA')`
const address2 = `INSERT INTO pfx_addresses (lines, postcode) VALUES ('2 Nutmeg Lane', 'FX1 2BB')`
const address3 = `INSERT INTO pfx_addresses (lines, postcode) VALUES ('Corner Shop', 'FX1 3CC')`
const address4 = `INSERT INTO pfx_addresses (lines, postcode) VALUES ('4 The Oaks', 'FX1 5EE')`

const person1a = `INSERT INTO pfx_persons (name, addressid) VALUES ('John Brown', %d)`
const person1b = `INSERT INTO pfx_persons (name, addressid) VALUES ('Mary Brown', %d)`

const person2a = `INSERT INTO pfx_persons (name, addressid) VALUES ('Anne Bollin', %d)`
