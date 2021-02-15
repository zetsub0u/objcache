package manager

import "sync"

// LazyCache is so named because i just use the std library, this is what i would use to reuse objects efficiently
// in high concurrency situations where i need to create and delete tons of objects quickly, this would avoid GC pressure
// This implementation makes no sense to be exposed via a REST api tho as it never gives objects just pointers.
type LazyCache struct {
	pool sync.Pool
}

func NewLazyCache() *LazyCache {
	return &LazyCache{
		pool: sync.Pool{
			New: func() interface{} {
				return new(int)
			},
		},
	}
}

func (c *LazyCache) GetObject() *int {
	return c.pool.Get().(*int)
}

func (c *LazyCache) FreeObject(obj *int) {
	c.pool.Put(obj)
}
