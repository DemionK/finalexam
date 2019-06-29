package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//db, err := sql.Open("postgres", "postgres://ynnpcvmn:ZCABSOPcWGU3c61KLafTj1nZQtrRmWeg@john.db.elephantsql.com:5432/ynnpcvmn")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTB() error {
	db, err := initDB()
	if err != nil {
		return err
	}
	defer db.Close()
	createTb := `
		CREATE TABLE IF NOT EXISTS customers(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			status TEXT
		);`
	_, err = db.Exec(createTb)
	if err != nil {
		return err
	}
	return nil
}

func InsertRow(name string, email string, status string) (*sql.Row, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	insertTb := `INSERT INTO customers (NAME, EMAIL, STATUS) VALUES ($1, $2, $3) RETURNING id`
	stmt, err := db.Prepare(insertTb)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(name, email, status)
	return row, nil
}

func SelectByID(rowID string) (*sql.Row, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	selectByIDTb := `select id, name, email, status from customers where id=$1;`
	stmt, err := db.Prepare(selectByIDTb)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(rowID)
	return row, nil
}

func SelectAll() (*sql.Rows, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	selectAllTb := `select id, name, email, status from customers;`
	stmt, err := db.Prepare(selectAllTb)
	if err != nil {
		return nil, err
	}
	row, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	return row, nil
}

func UpdateRow(rowID string, name string, email string, status string) (*sql.Row, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	updateByIDTb := `UPDATE customers SET name=$2, email=$3, status=$4 where id=$1 RETURNING ID;`
	stmt, err := db.Prepare(updateByIDTb)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(rowID, name, email, status)
	return row, nil
}

func DeleteRow(rowID string) (*sql.Row, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	deleteByIDTb := `delete from customers where id=$1 RETURNING id;`
	stmt, err := db.Prepare(deleteByIDTb)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(rowID)
	return row, nil
}
