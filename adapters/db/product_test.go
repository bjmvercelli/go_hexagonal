package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/bjmvercelli/go_hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insertSQL := `INSERT INTO products (id, name, price, status) VALUES ("1", "Product 1", 10.0, "disabled");`

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	ProductDb := db.NewProductDb(Db)
	product, err := ProductDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
