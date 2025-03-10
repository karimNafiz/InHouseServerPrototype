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
	var key_str serializable.Str = serializable.String(key) // we know key is a string so we directly create the wrapper type Str
	ok, val_serializable := serializable.CreateSerializableType(value)
	if !ok {
		return false
	}
	hash := GetHashStr(key_str, hashMap.capacity)
	hashMap.hashMap[hash].AppendHead(serializable.KeyValPair{
		Key:   key_str,
		Value: val_serializable,
	})
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
		return true, keyValPair.Value.GetValue()
	}

	return false, keyValPair
}

func get_element_linked_list(linkedList in_house_linked_list.LinkedList[serializable.KeyValPair], key string) (bool, serializable.KeyValPair) {
	current := linkedList.Head.Next

	for current != linkedList.Tail {
		if current.Value.Key.GetValue() == key { // must provide getters and setters remember to implement those changes
			return true, current.Value
		}
		current = current.Next
	}

	var zero_value serializable.KeyValPair
	return false, zero_value
}

func GetHashStr(key serializable.Str, cap int) int {
	return GetHash(key.Value, cap)
}

func GetHash(key string, cap int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % cap
}

func Getway() {
	testHashMap := CreateHashMap()
	if testHashMap.Add("key1", 42) {
		fmt.Println("for the key key1 42 is added as value ")
	}
	// if testHashMap.Add("key1", "hello world 2") {
	// 	fmt.Println(" for the key1 hello world 2 is added as value ")
	// }

	//testHashMap.Add("key2", 5.67)
	//testHashMap.Add("key3", false)
	testHashMap.Add("key2", "hello world")

	flag, value := testHashMap.Get("key1")
	if flag {
		fmt.Println(" value of key1 ", value)

	}
	flag, value = testHashMap.Get("key2")
	if flag {
		fmt.Println(" value of key2 ", value)
	}

	flag, value = testHashMap.Get("fail")
	if flag {
		fmt.Println(" value of key1 ", value)
	} else {
		fmt.Println("fail case:")
	}

}
