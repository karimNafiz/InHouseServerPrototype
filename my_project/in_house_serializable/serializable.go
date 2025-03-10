package serializable

import (
	"encoding/binary"
	"fmt"
	"math"
)

const Int32ByteLen = 4
const Int32ValueType = 0
const StringValueType = 2

type Serializable interface {
	Serialize(*[]byte)
	GetValue() any
}

type MetaData struct {
	MetaData [2]uint8
}

type Integer32 struct {
	Value    int
	MetaData MetaData
}

func create_meta_data(value_type uint8, byte_len uint8) MetaData {
	metaDataArray := [2]uint8{value_type, byte_len}
	return MetaData{MetaData: metaDataArray}
}

func Int32(value int) Integer32 {
	return Integer32{
		Value:    value,
		MetaData: create_meta_data(Int32ValueType, Int32ByteLen),
	}
}

func (wrapper Integer32) Serialize(byteArr *[]byte) {
	var tempArr [4]byte
	binary.BigEndian.PutUint32(tempArr[:], uint32(wrapper.Value))

	*byteArr = append(*byteArr, wrapper.MetaData.MetaData[:]...)
	*byteArr = append(*byteArr, tempArr[:]...)
}

func (wrapper Integer32) GetValue() any {
	return wrapper.Value
}

type Str struct {
	Value    string
	MetaData MetaData
}

func String(value string) Str {
	var str_len uint8
	if len(value) > int(math.MaxUint8) {
		fmt.Println("The length of the string '" + value + "' exceeds the max value (255). It will be truncated.")
		str_len = math.MaxUint8
		value = value[:math.MaxUint8]
	} else {
		str_len = uint8(len(value))
	}

	return Str{
		Value:    value,
		MetaData: create_meta_data(StringValueType, str_len),
	}
}

func (wrapper Str) Serialize(byteArray *[]byte) {
	*byteArray = append(*byteArray, wrapper.MetaData.MetaData[:]...)
	*byteArray = append(*byteArray, []byte(wrapper.Value)...)
}

func (wrapper Str) GetValue() any {
	return wrapper.Value
}

type KeyValPair struct {
	Key   Str
	Value Serializable
}

func (key_val_pair KeyValPair) Serialize(byteArray *[]byte) {
	key_val_pair.Key.Serialize(byteArray)
	key_val_pair.Value.Serialize(byteArray)
}

func (wrapper KeyValPair) GetValue() any {
	return wrapper.Value.GetValue()
}

func CreateSerializableType(value any) (bool, Serializable) {
	switch v := value.(type) {
	case string:
		return true, String(v)
	case int:
		return true, Int32(v)
	default:
		fmt.Println("Unsupported type")
		return false, nil // ✅ Explicitly return nil
	}
}
