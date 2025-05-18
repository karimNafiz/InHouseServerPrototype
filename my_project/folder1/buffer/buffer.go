package buffer

type Buffer struct {
	bufferCapacity int
	lastByteIndex  int
	buffer         []byte
}

func NewBuffer(capacity int) *Buffer {
	return &Buffer{
		bufferCapacity: capacity,
		buffer:         make([]byte, capacity),
		lastByteIndex:  0,
	}
}

func (b *Buffer) Add(data []byte) int {
	// how much room remains?
	free := b.bufferCapacity - b.lastByteIndex
	if free <= 0 {
		return 0
	}

	// only take as much as will fit
	toCopy := len(data)
	if toCopy > free {
		toCopy = free
	}

	// copy into the existing slice
	copy(b.buffer[b.lastByteIndex:], data[:toCopy])
	b.lastByteIndex += toCopy
	return toCopy
}

func (b *Buffer) Bytes() []byte {
	return b.buffer[:b.lastByteIndex]
}

func (b *Buffer) Reset() {
	b.lastByteIndex = 0
}
