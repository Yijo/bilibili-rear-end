package configer

import (
	"testing"
	"fmt"
)

// Test init config.
func TestInitConfig(t *testing.T) {

	InitConfig()

	fmt.Println("Init config success.")
}

// Test get a mysql configuration.
func TestMySQLConfig(t *testing.T) {

	InitConfig()

	mysqlConfig := GetMySQLConfig("bi_member")

	fmt.Println("data source name is: ", mysqlConfig.databaseName)
}

// Test get a redis configuration.
func TestRedisConfig(t *testing.T) {
	InitConfig()

	redisConfig := GetRedisConfig()

	fmt.Println("name: ", redisConfig.name)
	fmt.Println("address: ", redisConfig.address)
	fmt.Println("password: ", redisConfig.password)
}

// Test init time zone.
func TestInitZone(t *testing.T) {
	if err := initZone(); err != nil {
		fmt.Println("Init time zone failure, error is : ", err)
		return
	}

	fmt.Println("Init time zone success.")
}