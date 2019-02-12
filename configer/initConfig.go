package configer

import (
	//"bilibili-rear-end/configer/casbin"
	"bilibili-rear-end/configer/viper"
	"time"
	"fmt"
)

const (
	Sql = "mysql"
	Nosql = "redis"
)
// Init config.
func InitConfig() {
	// Init time zone
	if err := initZone(); err != nil {
		panic(fmt.Sprintf("Init time zone failure, error is : %v", err))
	}
	fmt.Println("Init time zone success.")


	// Init file config
	if err := viper.InitViperConfig("app"); err != nil {
		panic(fmt.Sprintf("Init file config failure, error is : %v", err))
	}
	fmt.Println("Init file config success.")

	//casbin.InitCasbin()       // Init Casbin
	//fmt.Println("Init casbin success.")

	// TODO: Init log collector
}

// Init time zone.
func initZone() error {
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	time.Local = local
	return nil
}


/*
Example:
{
	mysql: {
		db1: {
			data_source_name: "db1"
			},
		db2: {
			data_source_name: "db2"
			}
		}
}
*/

type MysqlConfig struct {
	DatabaseName string
}
// Get a mysql configuration.
func GetMySQLConfig(key string) MysqlConfig {
	if mysqlInfos := viper.AllViperSettings()[Sql].(map[string]interface{}); mysqlInfos != nil {
		// Get mysql information
		if mysqlInfo := mysqlInfos[key].(map[string]interface{}); mysqlInfo != nil {
			return MysqlConfig{
				DatabaseName: mysqlInfo["data_source_name"].(string),
			}
		}
	}
	// TODO: Collection logs
	panic("MySQLConfig is nil")
}


//
/*
Example:
{
	redis: {
		name: "redis",
		address: "127.0.0.1:6379",
		password: "",
		},
}
*/

type RedisConfig struct {
	DB string
	Address string
	Password string
}

// Get a redis configuration.
func GetRedisConfig() RedisConfig {
	// Get redis information
	if redisInfo := viper.AllViperSettings()[Nosql].(map[string]interface{}); redisInfo != nil {

		// address and name can't is nil
		if redisInfo["address"] == nil || redisInfo["name"] == nil {
			panic("RedisConfig is nil")
		}

		redisConfig := RedisConfig{
			DB: redisInfo["DB"].(string),
			Address: redisInfo["address"].(string),
		}

		if redisInfo["password"] != nil {
			redisConfig.Password = fmt.Sprintf("%v", redisInfo["password"])
		}

		return redisConfig
	}
	// TODO: Collection logs
	panic("RedisConfig is nil")
}

func GetCurrentEnvironment() string {
	return "production"
	return "test"
	return "release"
}

