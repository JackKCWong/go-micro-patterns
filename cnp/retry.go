package cnp

import "time"

func Retry(op func() (interface{}, error), retryCount int) (interface{}, error) {
	var err error
	var result interface{}
	for i := 0; i < retryCount; i++ {
		result, err = op()
		if err == nil {
			return result, nil
		}
	}
	return result, err
}

func RetryWithInterval(op func() (interface{}, error), retryCount int, interval time.Duration) (interface{}, error) {
	var err error
	var result interface{}
	for i := 0; i < retryCount; i++ {
		result, err = op()
		if err == nil {
			return result, nil
		}
		time.Sleep(interval)
	}
	return result, err
}

func RetryWithBackoff(op func() (interface{}, error), retryCount int, initialBackoff time.Duration) (interface{}, error) {
	var err error
	var result interface{}
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
