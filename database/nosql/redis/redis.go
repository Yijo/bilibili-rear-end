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
	// Single.
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

// Close releases the resources used by the pool.
func (r *Redis) closePool() {
	r.pool.Close()
}

// Do sends a command to the server and returns the received reply.
func (r *Redis) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Do(commandName, args)
}

// Send writes the command to the client's output buffer.
func (r *Redis) Send(commandName string, args ...interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Send(commandName, args)
}

// Flush flushes the output buffer to the Redis server.
func (r *Redis) Flush() error {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Flush()
}

// Receive receives a single reply from the Redis server.
func (r *Redis) Receive() (reply interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Receive()
}

// Close closes the connection.
func (r *Redis) Close() error {
	conn := r.pool.Get()
	return conn.Close()
}

// Err returns a non-nil value when the connection is not usable.
func (r *Redis) Err() error {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Err()
}