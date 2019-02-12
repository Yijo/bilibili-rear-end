package redis

import (
	"testing"
	"fmt"
)

// Get redis config test.
func TestGetRedisConfig(t *testing.T) {
	redisConfig := GetRedisConfig()

	fmt.Printf("redis www is: %s, redis address is: %s, redis password is: %s", redisConfig.Name, redisConfig.Address, redisConfig.Password)
}