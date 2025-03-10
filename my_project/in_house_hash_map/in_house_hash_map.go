package in_house_hash_map

import (
	"fmt"
	"hash/fnv"
	"my_project/in_house_linked_list"
)

// ✅ Capacity constant
const HashMapCapacity int = 16

// ✅ KeyValPair struct (Key must be exported)
type KeyValPair[T any] struct {
	Key   string // ✅ Now accessible outside
	Value T
}

// ✅ HashMap struct
type HashMap[T any] struct {
	capacity int
	hashMap  []in_house_linked_list.LinkedList[KeyValPair[T]]
}

// ✅ CreateHashMap function (properly initializes linked lists)
func CreateHashMap[T any]() *HashMap[T] {
	hashMap := &HashMap[T]{
		capacity: HashMapCapacity,
		hashMap:  make([]in_house_linked_list.LinkedList[KeyValPair[T]], HashMapCapacity),
	}

	// ✅ Initialize each linked list in the slice
	for i := range hashMap.hashMap {
		hashMap.hashMap[i] = *in_house_linked_list.CreateLinkedList[KeyValPair[T]]()
	}

	return hashMap
}

// ✅ Add function (correctly inserts into the hash map)
func (hashMap *HashMap[T]) Add(key string, value T) {
	hash := GetHash(key, hashMap.capacity)
	hashMap.hashMap[hash].AppendHead(KeyValPair[T]{
		Key:   key,
		Value: value,
	})
}

// ✅ Get function (fetches value by key)
func (hashMap *HashMap[T]) Get(key string) (bool, T) {
	hash := GetHash(key, hashMap.capacity)

	if hashMap.hashMap[hash].IsEmpty() {
		var defaultValue T
		return false, defaultValue
	}

	// ✅ Pass linked list by value (not pointer)
	keyValPair, found := GetHelper[T](hashMap.hashMap[hash], key)

	if found {
		return true, keyValPair.Value // ✅ Return the found value
	}

	var defaultValue T
	return false, defaultValue // ❌ Key not found, return zero-value
}

// ✅ GetHelper function (iterates through linked list to find key)
func GetHelper[T any](linkedList in_house_linked_list.LinkedList[KeyValPair[T]], key string) (KeyValPair[T], bool) {
	current := linkedList.Head.Next // ✅ Start from first actual node

	for current != linkedList.Tail { // ✅ Traverse the list
		if current.Value.Key == key { // ✅ Correctly access Key (capitalized)
			return current.Value, true // ✅ Found the key
		}
		current = current.Next
	}

	var zeroValue KeyValPair[T] // ✅ Return empty value
	return zeroValue, false     // ❌ Key not found
}

// ✅ GetHash function (generates a consistent hash for the given key)
func GetHash(key string, cap int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % cap
}

func Getway() {
	testHashMap := CreateHashMap[interface{}]()
	testHashMap.Add("key1", 42)
	testHashMap.Add("key1", "hello world 2")
	testHashMap.Add("key2", 5.67)
	testHashMap.Add("key3", false)
	testHashMap.Add("key4", "hello world")

	flag, value := testHashMap.Get("key1")
	if flag {
		fmt.Println(" value of key1 ", value)
	}
	flag, value = testHashMap.Get("key2")
	if flag {
		fmt.Println(" value of key1 ", value)
	}

	flag, value = testHashMap.Get("fail")
	if flag {
		fmt.Println(" value of key1 ", value)
	} else {
		fmt.Println("fail case:")
	}

}
