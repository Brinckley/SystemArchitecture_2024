package circuit_breaker

import "Gateway/internal/storage/cache"

type CircuitBreakerPost struct {
	serviceName  string
	state        string
	failureCount int
	maxFailures  int
	cache        cache.CircuitBreakerCache
}

func NewPostCircuitBreaker(maxFailures int, cache cache.CircuitBreakerCache) *CircuitBreakerPost {
	return &CircuitBreakerPost{
		serviceName:  "post_service",
		state:        StateClosed,
		failureCount: 0,
		maxFailures:  maxFailures,
		cache:        cache,
	}
}

func (c *CircuitBreakerPost) UpdateState() error {
	counter, err := c.cache.IncrementCounter(c.serviceName)
	if err != nil {
		return err
	}
	c.failureCount = counter
	if c.failureCount >= c.maxFailures {
		c.state = StateOpen
	}
	return nil
}

func (c *CircuitBreakerPost) ClearCounter() error {
	c.failureCount = 0
	c.state = StateClosed
	return c.cache.ClearCounter(c.serviceName)
}

func (c *CircuitBreakerPost) CheckValidCounter() bool {
	return c.failureCount < c.maxFailures
}

func (c *CircuitBreakerPost) IsCircuitOpen() bool {
	return c.state == StateOpen
}

func (c *CircuitBreakerPost) IsCircuitClosed() bool {
	return c.state == StateClosed
}

func (c *CircuitBreakerPost) IsCircuitHalfOpen() bool {
	return c.state == StateHalfOpen
}
