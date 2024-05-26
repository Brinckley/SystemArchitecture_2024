package circuit_breaker

const (
	StateClosed   = "closed"
	StateOpen     = "open"
	StateHalfOpen = "half-open"
)

type CircuitBreaker interface {
	UpdateState() error
	ClearCounter() error
	CheckValidCounter() bool
	IsCircuitOpen() bool
	IsCircuitClosed() bool
	IsCircuitHalfOpen() bool
}
