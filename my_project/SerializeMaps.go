package main

import (
	"fmt"
	de_serializer "my_project/in_house_de_serialize"
	"my_project/in_house_hash_map"
)

func main() {
	// msg := NewCustomMessage()

	// // Add key-value pairs
	// msg.Add("hellosssss", 42)
	// msg.Add("world", 100)

	// // Print the stored bytes
	// fmt.Println("Key Lengths:", msg.keyLengths)     // Should be [5, 5]
	// fmt.Println("Key Bytes:", stringWrappering(msg.keyBytes)) // Should be "helloworld"
	// fmt.Println("value Bytes:", msg.valueBytes)     // Should contain encoded 42 and 100 in big-endian format
	testHashMap := in_house_hash_map.CreateHashMap()
	testHashMap.Add("first_message_stress_test", "HelloWorld")
	testHashMap.Add("key2", "Python")

	byteStream := in_house_hash_map.SerializeHashMap(testHashMap)
	newHashMap, err := de_serializer.DeSerialize(&byteStream)
	if err == nil {
		fmt.Println(newHashMap.Get("first_message_stress_test"))
		fmt.Println(newHashMap.Get("key2"))
	}
}
