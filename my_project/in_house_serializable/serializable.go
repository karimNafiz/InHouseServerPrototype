package serializable

import (
	"encoding/binary"
	"fmt"
	"math"
	"my_project/in_house_linked_list"
)

const Integer32ByteLen = 4
const Integer32valueType = 0
const Float32valueType = 1
const Float32ByteLen = 4
const stringWrapperingvalueType = 2

type Serializable interface {
	Serialize(*[]byte)
	Getvalue() any
}

type metaData struct {
	metaData [2]uint8
}

func create_meta_data(value_type uint8, byte_len uint8) metaData {
	metaDataArray := [2]uint8{value_type, byte_len}
	return metaData{metaData: metaDataArray}
}

type list struct {
	values   *in_house_linked_list.LinkedList[Serializable]
	metaData metaData
}

func List(value []any) list {
	var returnVal list
	returnVal.values = in_house_linked_list.CreateLinkedList[Serializable]()
	for _, value := range value {
		ok, serializableType := CreateSerializableType(value)
		if !ok {
			fmt.Println("the package does not support the data type of  the value: ", value)
			fmt.Println("it will be skipped ")
			continue
		}
		returnVal.values.AppendHead(serializableType)
	}
	return returnVal
}

// TODO need to finish this
func (wrapper list) Serialize(byteArr *[]byte) {
	wrapper.values.ForEach(func(val Serializable) {
		val.Serialize(byteArr)
	})

}

func (wrapper list) Getvalue() any {
	var returnVal []any
	wrapper.values.ForEach(func(val Serializable) {
		returnVal = append(returnVal, val.Getvalue())
	})
	return returnVal
}

type float32Wrapper struct {
	value    float32
	metaData metaData
}

func Float32(value float32) float32Wrapper {
	return float32Wrapper{
		value:    value,
		metaData: create_meta_data(Float32valueType, Float32ByteLen),
	}
}

func (wrapper float32Wrapper) Serialize(byteArr *[]byte) {
	var tempArr [4]byte
	var uInteger32Representation uint32 = math.Float32bits(wrapper.value)
	binary.BigEndian.PutUint32(tempArr[:], uInteger32Representation)
	*byteArr = append(*byteArr, wrapper.metaData.metaData[:]...)
	*byteArr = append(*byteArr, tempArr[:]...)
}

func (wrapper float32Wrapper) Getvalue() any {
	return wrapper.value
}

type integer32Wrapper struct {
	value    int
	metaData metaData
}

func Integer32(value int) integer32Wrapper {
	return integer32Wrapper{
		value:    value,
		metaData: create_meta_data(Integer32valueType, Integer32ByteLen),
	}
}

func (wrapper integer32Wrapper) Serialize(byteArr *[]byte) {
	var tempArr [4]byte
	binary.BigEndian.PutUint32(tempArr[:], uint32(wrapper.value))

	*byteArr = append(*byteArr, wrapper.metaData.metaData[:]...)
	*byteArr = append(*byteArr, tempArr[:]...)
}

func (wrapper integer32Wrapper) Getvalue() any {
	return wrapper.value
}

type stringWrapper struct {
	value    string
	metaData metaData
}

func String(value string) stringWrapper {
	var stringWrapper_len uint8
	if len(value) > int(math.MaxUint8) {
		fmt.Println("The length of the stringWrappering '" + value + "' exceeds the max value (255). It will be truncated.")
		stringWrapper_len = math.MaxUint8
		value = value[:math.MaxUint8]
	} else {
		stringWrapper_len = uint8(len(value))
	}

	return stringWrapper{
		value:    value,
		metaData: create_meta_data(stringWrapperingvalueType, stringWrapper_len),
	}
}

func (wrapper stringWrapper) Serialize(byteArray *[]byte) {
	*byteArray = append(*byteArray, wrapper.metaData.metaData[:]...)
	*byteArray = append(*byteArray, []byte(wrapper.value)...)
}

func (wrapper stringWrapper) Getvalue() any {
	return wrapper.value
}

type KeyValPair struct {
	key   stringWrapper
	value Serializable
}

func CreateKeyValPair(key string, value any) (bool, KeyValPair) {
	ok, serializableValue := CreateSerializableType(value)
	var returnVal KeyValPair
	if ok {
		returnVal.key = String(key)
		returnVal.value = serializableValue
	}
	return ok, returnVal
}

func (key_val_pair KeyValPair) Serialize(byteArray *[]byte) {
	key_val_pair.key.Serialize(byteArray)
	key_val_pair.value.Serialize(byteArray)
}

func (wrapper KeyValPair) Getvalue() any {
	return wrapper.value.Getvalue()
}

func (wrapper KeyValPair) GetKey() string {
	return wrapper.key.value
}

func CreateSerializableType(value any) (bool, Serializable) {
	switch v := value.(type) {
	case float32:
		return true, Float32(v)
	case string:
		return true, String(v)
	case int:
		return true, Integer32(v)
	case []any:
		return true, List(v)
	default:
		fmt.Println("Unsupported type")
		return false, nil
	}
}
