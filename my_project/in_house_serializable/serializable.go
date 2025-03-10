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

type KeyValPair[T Serializable] struct {
	Key   Str
	Value T
}

func (key_val_pair KeyValPair[T]) Serialize(byteArray *[]byte) {
	key_val_pair.Key.Serialize(byteArray)
	key_val_pair.Value.Serialize(byteArray)
}
