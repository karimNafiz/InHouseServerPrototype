package serializable

import (
	"encoding/binary"
)

const Int32ByteLen = 4
const Int32ValueType = 0

type Serializable interface {
	Serialize(*[]byte)
}

type MetaData struct {
	MetaData *[2]uint8
}

type Integer32 struct {
	Value    int
	MetaData MetaData
}

func create_meta_data(value_type uint8, byte_len uint8) MetaData {
	metaDataArray := &[2]uint8{value_type, byte_len}
	return MetaData{
		MetaData: metaDataArray,
	}
}

func CreateInt32(value int) Integer32 {
	var byte_len uint8 = Int32ByteLen     //
	var value_type uint8 = Int32ValueType // this is hard coded change this later

	return Integer32{
		Value:    value,
		MetaData: create_meta_data(value_type, byte_len),
	}
}

// this function is not complete
func (wrapper Integer32) Serialize(byteArr *[]byte) {
	var tempArr [4]byte
	binary.BigEndian.PutUint32(tempArr[:], uint32(wrapper.Value))

	// Append metadata (value type and byte length)
	*byteArr = append(*byteArr, wrapper.MetaData.MetaData[:]...)

	// Append the integer value in BigEndian format
	*byteArr = append(*byteArr, tempArr[:]...)
}
