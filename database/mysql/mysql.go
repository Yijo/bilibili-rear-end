// Package mysql
package mysql

import (
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"fmt"
	"bilibili-rear-end/configer"
	"time"
)

// Defined database name.
const (
	member = "bi_member"
)

// Defined db error.
var dbNilErr = errors.New("db is nil")

var (
	// Single
	once sync.Once
	// Database list
	DBList = []string{member}
	DBWorkers = make(map[string]*DBWorker, len(DBList))
)

// Driver.
type DBWorker struct {
	*sql.DB
}

// Init mysql.
func InitDB() {
	once.Do(func() {
		for _, databaseName := range DBList {
			db := OpenDB(configer.GetMySQLConfig(databaseName).DatabaseName)

			// set the maximum amount of time a connection may be reused
			db.SetConnMaxLifetime(time.Second * 10)

			DBWorkers[databaseName] = &DBWorker{db}
			fmt.Printf("%v db create success \n", databaseName)
		}
	})
}

// Open mysql DB.
func OpenDB(dataSourceName string) *sql.DB {
	// Open doesn't open a connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("open %v fail: %v", dataSourceName, err.Error()))
	}

	// Validate DSN data
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("ping %v fail: %v", dataSourceName, err.Error()))
	}

	return db
}



// Get the specified database connection.
func getDB(key string) *DBWorker {
	return DBWorkers[key]
}

// Get member connection.
func MemberDB() *DBWorker {
	return getDB(member)
}



// Query a piece of data.
func (db *DBWorker) FetchRow(sqlStr string, args ...interface{}) (map[string] string, error) {
	rows, err := db.getRows(sqlStr, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// return an array of column name
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// create an array of receive values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string, len(scanArgs))

	// assign the value points of the subscript specified by values to the subscript corresponding to scanArgs
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var value string

		for i, column := range values {
			if column == nil {
				value = ""
			} else {
				value = string(column)
			}
			ret[columns[i]] = value
		}

		break	// exit only by taking the first one
	}
	return ret, nil
}

// Query a set of data.
func (db *DBWorker) FetchRows(sqlStr string, args ...interface{}) ([]map[string]string, error) {
	rows, err := db.getRows(sqlStr, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()


	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make([]map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var value string

		temp := make(map[string]string, len(values))
		for i, column := range values {
			if column == nil {
				value = ""
			} else {
				value = string(column)
			}
			temp[columns[i]] = value
		}
		ret = append(ret, temp)
	}
	return ret, nil
}


// Insert data.
func (db *DBWorker) Insert(sqlStr string, args ...interface{}) (int64, error) {
	if db == nil {
		return 0, dbNilErr
	}

	result, err := db.Exec(sqlStr, args)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// Change data.
func (db *DBWorker) ExecD(sqlStr string, args ...interface{}) (int64, error) {
	if db == nil {
		return 0, dbNilErr
	}

	result, err := db.Exec(sqlStr, args)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}




// Returns the query results as a *Rows
func (db *DBWorker) getRows(sqlStr string, args ...interface{}) (*sql.Rows, error) {

	// validate db
	if db == nil {
		return nil, dbNilErr
	}

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()


	rows, err := stmt.Query(args)
	if err != nil {
		return nil, err
	}

	return rows, nil
}