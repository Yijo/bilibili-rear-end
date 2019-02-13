package redis

import (
	"testing"
	"fmt"
)

// Get redis config test.
func TestGetRedisConfig(t *testing.T) {
	redisConfig := GetRedisConfig()

	fmt.Printf("redis www is: %s, redis address is: %s, redis password is: %s", redisConfig.DB, redisConfig.Address, redisConfig.Password)
}

// Init Redis test.
func TestInit(t *testing.T) {

}
