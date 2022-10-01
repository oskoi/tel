package checkers

import (
	health "github.com/tel-io/tel/v2/monitoring/heallth"
)

type ShutDownChecker struct{}

func ShutDown() *ShutDownChecker {
	return &ShutDownChecker{}
}

func (ShutDownChecker) Check() health.Health {
	check := health.NewHealth()
	check.Set(health.Down)

	return check
}
