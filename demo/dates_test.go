package demo

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/gomega"
	"github.com/rickb777/date"
	"github.com/rickb777/sqlapi/require"
	"testing"
)

func TestDatesCrud_using_database(t *testing.T) {
	g := NewGomegaWithT(t)

	examples := []date.Date{
		{}, // zero date
		date.New(2000, 3, 31),
		date.New(2020, 12, 31),
	}

	d := newDatabase(t)
	defer cleanup(d.DB())

	dt := NewDatesTable("dates", d)

	_, err := dt.DropTable(true)
	g.Expect(err).NotTo(HaveOccurred())

	_, err = dt.CreateTable(false)
	g.Expect(err).NotTo(HaveOccurred())

	for _, e := range examples {
		d0 := NewDates(e)
		err = dt.Insert(require.One, d0)
		g.Expect(err).NotTo(HaveOccurred())

		dx, err := dt.GetDatesById(require.One, d0.Id)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(dx).To(Equal(d0))
	}
}
