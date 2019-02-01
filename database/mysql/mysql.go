// Package mysql
package mysql

import (
	"errors"
	"database/sql"
	"sync"
	"fmt"
	"api-gbss/configer"
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

// Init databases.
func InitDB() {
	once.Do(func() {
		for _, databaseName := range DBList {
			db, err := OpenDB(configer.MySqlConfig().DataSourceName)
			// Open failure, termination procedure
			if err != nil {
				panic(err)
			}

			DBWorkers[databaseName] = &DBWorker{db}
			fmt.Printf("%v db create success \n", databaseName)
		}
	})
}

// Open mysql DB.
func OpenDB(dataSourceName string) (*sql.DB, error) {
	// Open doesn't open a connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	// Validate DSN data
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db, err
}

// 创建db连接.
func createDB(databaseName string) (*sql.DB, error) {
	// sql.Open返回的sql.DB对象是协程并发安全的
	db, err := sql.Open("mysql", databaseName)
	// 打开错误
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open %v fail: %v", databaseName, err))
	}

	// 验证连接, 连接错误
	if err := db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("ping %v fail: %v", databaseName, err))
	}

	return db, nil
}

// 获取指定DB.
func GetDB(key string) *DBWorker {
	return DBWorkers[key]
}

func MemberDB() *DBWorker {
	return GetDB(member)
}

// 查询一条数据.
func (db *DBWorker) FetchRow(sqlstr string, args ...interface{}) (map[string] string, error) {
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

// 查询一组数据.
func (db *DBWorker) FetchRows(sqlstr string, args ...interface{}) (*[]map[string] string, error) {
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

// 插入.
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

// 修改.