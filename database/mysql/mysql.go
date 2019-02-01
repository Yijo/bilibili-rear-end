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
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// 创建接收value的数组
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string, len(scanArgs))

	// 将values指定下标的value指针赋值给scanArgs对应的下标当中
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 遍历row
	for rows.Next() {
		// 对scanArgs进行赋值，响应的也改变了values
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var value string

		for i, column := range  values {
			if column == nil {
				value = ""
			} else {
				value = string(column)
			}
			ret[columns[i]] = value
		}

		break	// 只取第一条就退出
	}




	return ret, nil
}

// Query a set of data.
func (db *DBWorker) FetchRows(sqlStr string, args ...interface{}) (*[]map[string] string, error) {
	if db == nil {
		return nil, dbNilErr
	}

	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// 创建接收value的数组
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make([]map[string]string, len(scanArgs))

	// 将values指定下标的value指针赋值给scanArgs对应的下标当中
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 遍历row
	for rows.Next() {
		// 对scanArgs进行赋值，响应的也改变了values
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var value string

		temp := make(map[string]string, len(scanArgs))
		for i, column := range  values {
			if column == nil {
				value = ""
			} else {
				value = string(column)
			}
			temp[columns[i]] = value
		}
		ret = append(ret, temp)
	}
	return &ret, nil
}



// Insert data.
func Insert(db *sql.DB, sqlstr string, args ...interface{}) (int64, error){
	if db == nil {
		return 0, dbNilErr
	}

	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()


	return 0, nil
}

// Change data.





// Returns the query results as a *Rows
func (db *DBWorker) getRows(sqlStr string, args ...interface{}) (*sql.Rows, error) {
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