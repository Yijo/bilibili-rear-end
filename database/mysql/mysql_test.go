package mysql

import (
	"testing"
	"bilibili-rear-end/configer"
	"fmt"
)

// Init mysql test.
func TestInitDB(t *testing.T) {
	configer.InitConfig()
	InitDB()
}

// Open mysql DB test.
func TestOpenDB(t *testing.T) {
	configer.InitConfig()
	OpenDB(configer.GetMySQLConfig(member).DatabaseName)
}



// Get the specified database connection test.
func TestGetDB(t *testing.T) {
	configer.InitConfig()
	InitDB()

	fmt.Println(getDB(member))
}

// Get member connection test.
func TestMemberDB(t *testing.T) {
	configer.InitConfig()
	InitDB()

	fmt.Println(MemberDB())
}


// Query a piece of data.
func TestFetchRow(t *testing.T) {
	configer.InitConfig()
	InitDB()

}

// Query a set of data test.
func TestFetchRows(t *testing.T) {
	configer.InitConfig()
	InitDB()
}

func TestInsert(t *testing.T) {
	configer.InitConfig()
	InitDB()
}