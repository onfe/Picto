package server

//CircularQueue is a circular queue
type CircularQueue struct {
	queue     []interface{}
	HeadIndex int `json:"HeadIndex"`
	TailIndex int `json:"TailIndex"`
	Capacity  int `json:"Capacity"`
	Len       int `json:"Len"`
}

func newCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		queue:     make([]interface{}, capacity),
		HeadIndex: 1,
		TailIndex: 1,
		Capacity:  capacity,
	}
}

func (c *CircularQueue) push(x interface{}) {
	c.HeadIndex = (c.HeadIndex + 1) % c.Capacity
	c.Len++
	c.queue[c.HeadIndex] = x
}

func (c *CircularQueue) getAll() []interface{} {
	return c.queue
}
