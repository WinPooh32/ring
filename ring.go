package ring

// Ring is a fixed-size circular buffer.
type Ring[T any] struct {
	buf []T
	cap int
	len int
	ptr int
}

// Make makes a new instance of the ring buffer with given capacity.
func Make[T any](cap int) Ring[T] {
	return Ring[T]{
		buf: make([]T, cap),
		cap: cap,
		len: 0,
		ptr: 0,
	}
}

// Len returns count of items.
func (r *Ring[T]) Len() int {
	return r.len
}

// Cap returns maximum capacity of the ring buffer.
func (r *Ring[T]) Cap() int {
	return r.cap
}

// Reset resets buffer to zero length.
func (r *Ring[T]) Reset() {
	r.ptr = 0
	r.len = 0
}

// Full returns true if the buffer is full.
func (r *Ring[T]) Full() bool {
	return r.len >= r.cap
}

// Push adds new item to the end of buffer
// and pops first item if buffer is full.
func (r *Ring[T]) Push(value T) (elem T, pop bool) {
	if r.Full() {
		pop = true
		elem = r.popFront()
	}

	r.pushBack(value)

	return elem, pop
}

// Back returns the last item of the buffer.
func (r *Ring[T]) Back() (value T) {
	return r.buf[(r.ptr+r.len-1)%r.cap]
}

// Front returns the first item of the buffer.
func (r *Ring[T]) Front() (value T) {
	return r.buf[r.ptr]
}

// Range applies f to every element of the buffer.
func (r *Ring[T]) Range(f func(i int, v T)) {
	left, right := r.TwoParts()
	length := len(left)

	for i, v := range left {
		f(i, v)
	}

	for i, v := range right {
		f(length+i, v)
	}
}

// Copy makes copy of the ring buffer to dst slice.
// Returns count of copied items.
func (r *Ring[T]) CopyTo(dst []T) (n int) {
	if len(dst) < r.cap {
		panic("dst length must be equal or greater ring size")
	}
	end := r.ptr + r.len
	if end <= r.cap {
		copy(dst, r.buf[r.ptr:end])
		return r.len
	}
	p := end % r.cap
	copy(dst, r.buf[r.ptr:r.cap])
	copy(dst[r.cap-r.ptr:], r.buf[0:p])
	return r.cap - r.ptr + p
}

// TwoParts returns ordered slices of underlying buffer.
func (r *Ring[T]) TwoParts() (left, right []T) {
	end := r.ptr + r.len
	if end <= r.cap {
		return r.buf[r.ptr:end:end], nil
	}
	p := end % r.cap

	left = r.buf[r.ptr:r.cap:r.cap]
	right = r.buf[0:p:p]

	return left, right
}

// pushBack adds new item to the end of the buffer.
func (r *Ring[T]) pushBack(value T) {
	end := (r.ptr + r.len) % r.cap
	r.buf[end] = value
	r.len++
}

// popFront deletes first item and returns it.
func (r *Ring[T]) popFront() (value T) {
	value = r.buf[r.ptr]
	r.ptr = (r.ptr + 1) % r.cap
	r.len--
	return value
}
