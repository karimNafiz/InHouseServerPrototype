package main

import "my_project/in_house_hash_map"

// type CustomMessage struct {
// 	payload    map[stringWrappering]int
// 	keyCount   int
// 	keyLengths []byte
// 	keyBytes   []byte
// 	valueBytes []byte
// }

// // ConstringWrappeructor for CustomMessage
// func NewCustomMessage() *CustomMessage {
// 	customMessage := new(CustomMessage)
// 	customMessage.payload = make(map[stringWrappering]int)
// 	customMessage.keyCount = 0
// 	customMessage.keyLengths = make([]byte, 0)
// 	customMessage.keyBytes = make([]byte, 0)
// 	customMessage.valueBytes = make([]byte, 0)

// 	return customMessage
// }

// // Add method for CustomMessage
// func (msg *CustomMessage) Add(key stringWrappering, value int) bool {
// 	// Ensure the key length is within 0 <= len <= 255 (2^8)
// 	if len(key) > math.MaxUint8 {
// 		return false
// 	}

// 	// Store key length as uint8
// 	var keyLen uint8 = uint8(len(key))
// 	msg.keyLengths = append(msg.keyLengths, keyLen)

// 	// Append key bytes to keyBytes slice
// 	AppendstringWrapperBytes(&msg.keyBytes, key)

// 	// Determine system int size (4 bytes for 32-bit, 8 bytes for 64-bit)
// 	const systemIntSize uint8 = uint8(unsafe.Sizeof(int(0)))

// 	// Create a byte array to store the integer
// 	var valueArr [8]byte // Maximum possible size needed

// 	if systemIntSize == 4 {
// 		binary.BigEndian.PutUInteger32(valueArr[:4], uInteger32(value)) // Convert int to 4-byte slice
// 		AppendIntBytes(&msg.valueBytes, valueArr[:4])                   // Append to valueBytes
// 	} else {
// 		binary.BigEndian.PutUint64(valueArr[:8], uint64(value)) // Convert int to 8-byte slice
// 		AppendIntBytes(&msg.valueBytes, valueArr[:8])           // Append to valueBytes
// 	}

// 	return true
// }

// // Function to append stringWrappering bytes to a byte slice (using pointer)
// func AppendstringWrapperBytes(byteSlice *[]byte, key stringWrappering) {
// 	*byteSlice = append(*byteSlice, key...) // More efficient than looping
// }

// // Function to append an integer byte slice to another byte slice
// func AppendIntBytes(byteSlice *[]byte, value []byte) {
// 	AppendIntBytesMain(byteSlice, value, uint8(len(value)))
// }

// // Core function to append integer bytes
// func AppendIntBytesMain(byteSlice *[]byte, value []byte, intSize uint8) {
// 	for i := uint8(0); i < intSize; i++ {
// 		*byteSlice = append(*byteSlice, value[i])
// 	}
// }

// Main function to test the implementation
func main() {
	// msg := NewCustomMessage()

	// // Add key-value pairs
	// msg.Add("hellosssss", 42)
	// msg.Add("world", 100)

	// // Print the stored bytes
	// fmt.Println("Key Lengths:", msg.keyLengths)     // Should be [5, 5]
	// fmt.Println("Key Bytes:", stringWrappering(msg.keyBytes)) // Should be "helloworld"
	// fmt.Println("value Bytes:", msg.valueBytes)     // Should contain encoded 42 and 100 in big-endian format

	in_house_hash_map.Getway()
}
