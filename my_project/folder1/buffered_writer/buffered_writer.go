package buffered_writer

import (
	buffer "my_project/folder1/buffer"
	utility "my_project/utility"
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

func (bWriter *BufferedWriter) Write(bytes []byte) (int, error) {

	// lenBytes represent the total no of bytes we need to add to the buffer
	lenBytes := len(bytes)

	// available represents the total space in no of bytes available in the buffer
	available := bWriter.getAvailableSpaceInBuffer()

	// start index (for the slice bytes) for the data to be added to the buffer
	start := 0
	// end index (for the slice bytes) for the data to be added to the buffer
	end := utility.MinInt(start+bWriter.buffer.GetCapacity(), lenBytes)

	// if there is no space for the incoming byte slice
	// we need to flush whats previously stored in the buffer
	// as we need to send bytes over the IO connection as a whole
	// we can't have a scenario where we send x percent of a message in one Write function and send y percent of that message in another write function; (x + y) = 100
	if lenBytes > available && !bWriter.buffer.IsEmpty() {
		_, err := bWriter.flush()
	} else { // if we do have space we add it to the buffer
		// if the buffer becomes full we need to flush too
		n := bWriter.buffer.Add(bytes[start:end])
		if bWriter.buffer.IsFull() {
			_, err := bWriter.flush()
		}
		return n, nil
	}

	//

	for start < end {
		n := bWriter.buffer.Add(bytes[start:end])
		_, err := bWriter.flush()
		start += n
		end = utility.MinInt(start+bWriter.buffer.GetCapacity(), lenBytes)
		available = bWriter.getAvailableSpaceInBuffer()
	}

}

func (bWriter *BufferedWriter) flush() (int, error) {

}

func (bWriter *BufferedWriter) getAvailableSpaceInBuffer() int {
	return bWriter.buffer.GetCapacity() - bWriter.buffer.GetTotalBytes()
}
