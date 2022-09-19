package health

import (
	"context"
	"sync"
)

func RunHealthChecks(ctx context.Context, checkers []Checker) []CheckResult {
	checkersCount := len(checkers)
	results := make([]CheckResult, 0, checkersCount)

	wg := sync.WaitGroup{}
	wg.Add(checkersCount)

	for _, checker := range checkers {
		checker := checker

		go func() {
			results = append(results, checker.Check(ctx))

			wg.Done()
		}()
	}

	wg.Wait()

	return results
}
