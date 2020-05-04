package server

type eventCache struct {
	queue     []interface{}
	HeadIndex int `json:"HeadIndex"`
	TailIndex int `json:"TailIndex"`
	Capacity  int `json:"Capacity"`
	Len       int `json:"Len"`
}

func newEventCache(capacity int) *eventCache {
	return &eventCache{
		queue:     make([]interface{}, capacity),
		HeadIndex: 1,
		TailIndex: 1,
		Capacity:  capacity,
	}
}

func (c *eventCache) push(x interface{}) {
	c.HeadIndex = (c.HeadIndex + 1) % c.Capacity
	c.Len++
	c.queue[c.HeadIndex] = x
}

func (c *eventCache) getAll() []interface{} {
	return c.queue
}
