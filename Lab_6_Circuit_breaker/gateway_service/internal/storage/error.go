package storage

type CacheError struct {
	message string
}

func NewCacheError(message string) *CacheError {
	return &CacheError{
		message: message,
	}
}

func (c *CacheError) Error() string {
	return c.message
}
