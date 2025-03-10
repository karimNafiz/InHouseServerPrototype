package in_house_hash_map

import (
	"fmt"
	"hash/fnv"
	"my_project/in_house_linked_list"
	serializable "my_project/in_house_serializable"
)

const HashMapCapacity int = 16

type HashMap[T serializable.Serializable] struct {
	capacity int
	hashMap  []in_house_linked_list.LinkedList[serializable.KeyValPair[T]]
}

func CreateHashMap[T serializable.Serializable]() *HashMap[T] {
	hashMap := &HashMap[T]{
		capacity: HashMapCapacity,
		hashMap:  make([]in_house_linked_list.LinkedList[serializable.KeyValPair[T]], HashMapCapacity),
	}

	for i := range hashMap.hashMap {
		hashMap.hashMap[i] = *in_house_linked_list.CreateLinkedList[serializable.KeyValPair[T]]()
	}

	return hashMap
}

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

func GetHash(key serializable.Str, cap int) int {
	h := fnv.New32a()
	h.Write([]byte(key.Value))
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
