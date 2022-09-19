package health

import "context"

type Checker interface {
	Check(ctx context.Context) CheckResult
}

type CheckResult struct {
	Service string
	Status  bool
}
