package cockring

const bufferInitialSize = 8

type Buffer struct {
	buffer   []interface{}
	head     int
	tail     int
	nonEmpty bool
}

func (r Buffer) Len() int {
	if !r.nonEmpty {
		return 0
	}

	if r.head < r.tail {
		return r.tail - r.head
	} else if r.head == r.tail {
		return cap(r.buffer)
	} else {
		return cap(r.buffer) + r.tail - r.head
	}
}

func (r *Buffer) AddFirst(element interface{}) {
	if cap(r.buffer) == 0 {
		r.buffer = make([]interface{}, bufferInitialSize)
		r.buffer[0] = element
		r.tail = 1
	} else {
		if r.Len() == cap(r.buffer) {
			newBuffer := make([]interface{}, 2*cap(r.buffer))
			if r.head < r.tail {
				copy(newBuffer[:r.Len()], r.buffer[r.head:r.tail])
			} else {
				copy(newBuffer[:cap(r.buffer)-r.head], r.buffer[r.head:])
				copy(newBuffer[cap(r.buffer)-r.head:r.Len()], r.buffer[:r.tail])
			}
			r.head = 0
			r.tail = cap(r.buffer)
			r.buffer = newBuffer
		}
		r.head = (cap(r.buffer) + r.head - 1) % cap(r.buffer)
		r.buffer[r.head] = element
	}
	r.nonEmpty = true
}

func (r *Buffer) AddLast(element interface{}) {
	if cap(r.buffer) == 0 {
		r.buffer = make([]interface{}, bufferInitialSize)
		r.buffer[0] = element
		r.tail = 1
	} else {
		if r.Len() == cap(r.buffer) {
			newBuffer := make([]interface{}, 2*cap(r.buffer))
			if r.head < r.tail {
				copy(newBuffer[:r.Len()], r.buffer[r.head:r.tail])
			} else {
				copy(newBuffer[:cap(r.buffer)-r.head], r.buffer[r.head:])
				copy(newBuffer[cap(r.buffer)-r.head:r.Len()], r.buffer[:r.tail])
			}
			r.head = 0
			r.tail = cap(r.buffer)
			r.buffer = newBuffer
		}
		r.buffer[r.tail] = element
		r.tail = (r.tail + 1) % cap(r.buffer)
	}
	r.nonEmpty = true
}

func (r Buffer) Get(pos int) interface{} {
	if !r.nonEmpty || pos < 0 || pos >= r.Len() {
		panic("unexpected behavior: index out of bounds")
	}
	return r.buffer[(pos+r.head)%cap(r.buffer)]
}

func (r Buffer) GetFirst() interface{} {
	if !r.nonEmpty {
		panic("unexpected behavior: getting first from empty deque")
	}
	return r.buffer[(cap(r.buffer)+r.tail-1)%cap(r.buffer)]
}

func (r Buffer) GetLast() interface{} {
	if !r.nonEmpty {
		panic("unexpected behavior: getting last from empty deque")
	}

	return r.buffer[(cap(r.buffer)+r.tail-1)%cap(r.buffer)]
}

func (r *Buffer) RemoveFirst() {
	if r.Len() == 0 {
		panic("removing first from empty ring buffer")
	}
	r.buffer[r.head] = nil
	r.head = (r.head + 1) % cap(r.buffer)
	if r.head == r.tail {
		r.nonEmpty = false
	}
}

func (r *Buffer) RemoveLast() {
	if r.Len() == 0 {
		panic("removing last from empty ring buffer")
	}

	lastPos := (cap(r.buffer) + r.tail - 1) % cap(r.buffer)
	r.buffer[lastPos] = nil
	r.tail = lastPos
	if r.tail == r.head {
		r.nonEmpty = false
	}
}

func (r *Buffer) Reset() {
	r.head = 0
	r.tail = 0
	r.nonEmpty = false
}
