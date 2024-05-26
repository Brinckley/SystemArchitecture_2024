package circuit_breaker

import (
	"Gateway/internal/storage/cache"
	"log"
)

type CircuitBreakerMessage struct {
	serviceName  string
	state        string
	failureCount int
	maxFailures  int
	cache        cache.CircuitBreakerCache
}

func NewMessageCircuitBreaker(maxFailures int, cache cache.CircuitBreakerCache) *CircuitBreakerPost {
	return &CircuitBreakerPost{
		serviceName:  "message_service",
		state:        StateClosed,
		failureCount: 0,
		maxFailures:  maxFailures,
		cache:        cache,
	}
}

func (c *CircuitBreakerMessage) UpdateState() error {
	counter, err := c.cache.IncrementCounter(c.serviceName)
	if err != nil {
		return err
	}
	c.failureCount = counter + 1
	log.Printf("In the message circuit breaker new counter setting %d", c.failureCount)
	if c.failureCount >= c.maxFailures {
		c.state = StateOpen
	}
	return nil
}

func (c *CircuitBreakerMessage) ClearCounter() error {
	c.failureCount = 0
	c.state = StateClosed
	return c.cache.ClearCounter(c.serviceName)
}

func (c *CircuitBreakerMessage) CheckValidCounter() bool {
	return c.failureCount < c.maxFailures
}

func (c *CircuitBreakerMessage) IsCircuitOpen() bool {
	return c.state == StateOpen
}

func (c *CircuitBreakerMessage) IsCircuitClosed() bool {
	return c.state == StateClosed
}

func (c *CircuitBreakerMessage) IsCircuitHalfOpen() bool {
	return c.state == StateHalfOpen
}
