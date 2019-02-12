package redis

import (
	"strconv"
	"bilibili-rear-end/configer"
	"sync"
	"github.com/garyburd/redigo/redis"
	"time"
)

type Redis struct {
	pool *redis.Pool
}

var (
	redisOnce sync.Once

	redisInstance *Redis
)

type RedisConfig struct {
	DB int
	Address string
	Password string
}

// Get redis config.
func GetRedisConfig() RedisConfig {
	name, err := strconv.Atoi(configer.GetRedisConfig().DB)
	if err != nil {
		panic(err.Error())
	}

	return RedisConfig{
		DB: name,
		Address: configer.GetRedisConfig().Address,
		Password: configer.GetRedisConfig().Password,
	}
}

// Init Redis.
func Init() *Redis {
	redisOnce.Do(func() {
		pool := &redis.Pool{
			MaxIdle: 3,
			IdleTimeout: 60 * time.Second,
			Dial: func() (redis.Conn, error) {
				redisConfig := GetRedisConfig()

				conn, err := redis.Dial("tcp", redisConfig.Address)
				if err != nil {
					return nil, err
				}

				if redisConfig.Password != "" {
					if _, err := conn.Do("AUTH", redisConfig.Password); err != nil {
						conn.Close()
						return nil, err
					}
				}

				if _, err := conn.Do("SELECT", redisConfig.DB); err != nil {
					conn.Close()
					return nil, err
				}

				return conn, nil
			},

			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}

		redisInstance = &Redis{
			pool: pool,
		}

		redisInstance.closePool()
	})

	return redisInstance
}

// Get Redis instance.
func GetInstance() *Redis {
	if redisInstance == nil {
		panic("")
	}
	return  redisInstance
}


func (r *Redis) closePool() {

}