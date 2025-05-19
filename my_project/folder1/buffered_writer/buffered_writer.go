package buffered_writer

import (
	buffer "my_project/folder1/buffer"
	"net"
)

type BufferedWriter struct {
	// should have a pointer to buffer
	buffer *buffer.Buffer

	// need to have the connection object that we are encapsulating
	// the net.Conn is an interface
	// so we dont need a pointer to it
	ioConn net.Conn
}

// Write buffers data and flushes to the connection when necessary.
// It returns the total number of bytes accepted (either buffered or sent)
// and any error encountered during flushing or direct writes.
func (bWriter *BufferedWriter) Write(bytes []byte) (int, error) {
	// Track total bytes successfully handled
	totalWritten := 0

	// Incoming payload length and buffer capacity
	lenBytes := len(bytes)
	bufCap := bWriter.buffer.GetCapacity()

	// Compute current free space in buffer
	available := bWriter.getAvailableSpaceInBuffer()

	// ----- Case 1: Entire payload fits in existing buffer -----
	if lenBytes <= available {
		// Add all data to buffer
		n := bWriter.buffer.Add(bytes)
		totalWritten += n

		// If buffer is now full, flush its contents
		if bWriter.buffer.IsFull() {
			// if there is an error with the flush
			// we will return totalWritten which could be zero or the bytes added
			if err := bWriter.flush(); err != nil {
				return totalWritten, err
			}
		}
		return totalWritten, nil
	}

	// ----- Case 2: Payload too large for current buffer space -----
	// Flush any existing buffered data first
	if !bWriter.buffer.IsEmpty() {
		if err := bWriter.flush(); err != nil {
			return totalWritten, err
		}
	}

	// Recompute available space after flush
	available = bWriter.getAvailableSpaceInBuffer()

	// ----- Case 2a: After flush, payload now fits entirely -----
	if lenBytes <= available {
		n := bWriter.buffer.Add(bytes)
		totalWritten += n

		// If buffer is now full, flush its contents
		if bWriter.buffer.IsFull() {
			// if there is an error with the flush
			// we will return totalWritten which could be zero or the bytes added
			if err := bWriter.flush(); err != nil {
				return totalWritten, err
			}
		}

		return totalWritten, nil
	}

	// ----- Case 2b: Payload still larger than buffer capacity -----
	// Break payload into buffer-sized chunks, buffer each, then flush
	start := 0
	for start < lenBytes {
		// Determine chunk boundaries
		end := start + bufCap
		if end > lenBytes {
			end = lenBytes
		}

		// Buffer this chunk (Add returns actual bytes added)
		n := bWriter.buffer.Add(bytes[start:end])
		totalWritten += n

		// Flush the buffered chunk to the connection
		if err := bWriter.flush(); err != nil {
			return totalWritten, err
		}

		// Advance past the chunk we just wrote
		start += n
	}

	return totalWritten, nil
}
func (bWriter *BufferedWriter) flush() error {
	// if buffer is empty we will return a error
}

func (bWriter *BufferedWriter) getAvailableSpaceInBuffer() int {
	return bWriter.buffer.GetCapacity() - bWriter.buffer.GetTotalBytes()
}
