package casbin

import (
	"github.com/casbin/casbin"
	"log"
)

func InitCasbin() *casbin.Enforcer {
	// New a safe Casbin enforcer with a models file and policy file
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Load the policy from DB.
	authEnforcer.LoadPolicy()

	authEnforcer.Enforce("alice", "data1", "read")

	// Modify the policy
	// authEnforcer.AddPolicy()
	// authEnforcer.RemovePolicy()

	// Save the policy back to DB.
	authEnforcer.SavePolicy()

	return authEnforcer
}
