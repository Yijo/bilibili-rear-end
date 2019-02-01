package mysql

import (
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"fmt"
)

//// Mysql driver.
//type DBWorker struct {
//	*sql.DB
//}

var (
	dbList = []string{member}
	driver *DBWorker
)

type mysqlConfig struct {
	DataBaseName string
}

//func InitMysqlConfiger() mysqlConfig {
//
//}

//// Open mysql DB.
//func OpenDB() {
//	// Open doesn't open a connection
//	driver, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bi_member?charset=utf8")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// Validate DSN data
//	err = driver.Ping()
//	if err != nil {
//		panic(err.Error())
//	}
//}


//// Close mysql DB.
//func (driver *DBWorker) CloseDB() {
//	err := driver.Close()
//	if err != nil {
//		panic(err)
//	}
//}
//
////
//func (driver *DBWorker) FetchRow(sql string, args ...interface{}) (error){
//	// Pretreatment
//	stmt, err := driver.Prepare(sql)
//	if err != nil {
//		return err
//	}
//	defer stmt.Close()
//
//	row := stmt.QueryRow(args)
//
//	fmt.Println(row)
//
//	return nil
//}
