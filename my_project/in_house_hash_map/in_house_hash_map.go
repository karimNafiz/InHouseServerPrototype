package in_house_hash_map

import (
	"fmt"
	"hash/fnv"
	"my_project/in_house_linked_list"
	serializable "my_project/in_house_serializable"
)

const HashMapCapacity int = 16

type HashMap struct {
	capacity int
	hashMap  []in_house_linked_list.LinkedList[serializable.KeyValPair]
}

func CreateHashMap() *HashMap {
	hashMap := &HashMap{
		capacity: HashMapCapacity,
		hashMap:  make([]in_house_linked_list.LinkedList[serializable.KeyValPair], HashMapCapacity),
	}

	for i := range hashMap.hashMap {
		hashMap.hashMap[i] = *in_house_linked_list.CreateLinkedList[serializable.KeyValPair]()
	}

	return hashMap
}

func (hashMap *HashMap) Add(key string, value any) bool {
	hash := GetHash(key, hashMap.capacity)
	ok, keyValPair := serializable.CreateKeyValPair(key, value)
	if !ok {
		return ok
	}
	hashMap.hashMap[hash].AppendHead(keyValPair)
	return true
}

func (hashMap *HashMap) Get(key string) (bool, any) {
	hash := GetHash(key, hashMap.capacity)

	if hashMap.hashMap[hash].IsEmpty() {
		var defaultValue serializable.Serializable
		return false, defaultValue
	}

	found, keyValPair := get_element_linked_list(hashMap.hashMap[hash], key)

	if found {
		return true, keyValPair.Getvalue()
	}

	return false, keyValPair
}

func get_element_linked_list(linkedList in_house_linked_list.LinkedList[serializable.KeyValPair], key string) (bool, serializable.KeyValPair) {
	current := linkedList.Head.Next

	for current != linkedList.Tail {
		if current.Value.GetKey() == key { // must provide getters and setters remember to implement those changes
			return true, current.Value
		}
		current = current.Next
	}

	var zero_value serializable.KeyValPair
	return false, zero_value
}

func GetHash(key string, cap int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % cap
}

func Getway() {
	testHashMap := CreateHashMap()

	testHashMap.Add("key1", 42)
	testHashMap.Add("key2", "hello world")
	testHashMap.Add("key3", 150)
	testHashMap.Add("key4", "value for key 4")
	testHashMap.Add("key5", 34.12)
	testHashMap.Add("key6", []any{"crack", 14, 34.12})

	byteStream := SerializeHashMap(testHashMap)
	fmt.Println("hopefully this works")
	fmt.Println(byteStream)

}
func (hashMap *HashMap) Serialize(byteArray *[]byte) {
	// Loop through all linked lists (buckets in the hashmap)
	for i := 0; i < hashMap.capacity; i++ {
		current := hashMap.hashMap[i].Head.Next // Start from the first real node

		// Iterate through the linked list and serialize each KeyValPair
		for current != hashMap.hashMap[i].Tail {
			current.Value.Serialize(byteArray) // Serialize KeyValPair
			current = current.Next
		}
	}
}

// design coice im choosing to send a hashmap pointer
// even toh the hashMap struct is not very big
// as most of the linkedList is stored in the heap
func SerializeHashMap(hashMap *HashMap) []byte {
	// initial size 16
	// later do some math to find out the best length
	byteStream := make([]byte, 0)
	fmt.Println("before serialization ")
	fmt.Println(byteStream)
	hashMap.Serialize(&byteStream)

	return byteStream

}
