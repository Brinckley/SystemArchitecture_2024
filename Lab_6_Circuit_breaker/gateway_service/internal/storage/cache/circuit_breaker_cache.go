package cache

type CircuitBreakerCache interface {
	IncrementCounter(serviceName string) (int, error)
	ClearCounter(serviceName string) error
	GetCounter(serviceName string) (int, error)
}
