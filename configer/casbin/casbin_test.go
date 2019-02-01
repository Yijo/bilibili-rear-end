package casbin

import (
	"testing"
	"fmt"
)

func TestCasbin(t *testing.T) {

	authEnforcer := InitCasbin()

	// Check the permission
	fmt.Println(authEnforcer.Enforce("admin", "adminData", "write"))
}
