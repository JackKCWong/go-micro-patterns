package cnp

import "time"

func Retry[T any](op func() (T, error), retryCount int) (T, error) {
	var err error
	var result T
	for i := 0; i < retryCount; i++ {
		result, err = op()
		if err == nil {
			return result, nil
		}
	}
	return result, err
}

func RetryWithInterval[T any](op func() (T, error), retryCount int, interval time.Duration) (T, error) {
	var err error
	var result T
	for i := 0; i < retryCount; i++ {
		result, err = op()
		if err == nil {
			return result, nil
		}
		time.Sleep(interval)
	}
	return result, err
}

func RetryWithBackoff[T any](op func() (T, error), retryCount int, initialBackoff time.Duration) (T, error) {
	var err error
	var result T
	var expBackoff time.Duration = initialBackoff
	for i := 0; i < retryCount; i++ {
		result, err = op()
		if err == nil {
			return result, nil
		}

		expBackoff = expBackoff * 2
		time.Sleep(expBackoff)
	}

	return result, err
}
