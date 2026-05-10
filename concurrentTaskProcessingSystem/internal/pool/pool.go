package pool

import "sync"

var BufferPool = sync.Pool{
	New: func() interface{} {
		buffer := make([]byte, 1024)
		return buffer
	},
}
