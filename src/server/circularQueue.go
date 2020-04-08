package server

//CircularQueue is a circular queue
type circularQueue struct {
	queue     []interface{}
	HeadIndex int `json:"HeadIndex"`
	TailIndex int `json:"TailIndex"`
	Capacity  int `json:"Capacity"`
	Len       int `json:"Len"`
}

func newCircularQueue(capacity int) *circularQueue {
	return &circularQueue{
		queue:     make([]interface{}, capacity),
		HeadIndex: 1,
		TailIndex: 1,
		Capacity:  capacity,
	}
}

func (c *circularQueue) push(x interface{}) {
	c.HeadIndex = (c.HeadIndex + 1) % c.Capacity
	c.Len++
	c.queue[c.HeadIndex] = x
}

func (c *circularQueue) pop() interface{} {
	v := c.queue[c.TailIndex]
	c.Len--
	c.TailIndex = (c.TailIndex + 1) % c.Capacity
	return v
}

func (c *circularQueue) getAll() []interface{} {
	return c.queue
}
