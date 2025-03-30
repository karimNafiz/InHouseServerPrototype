package in_house_de_serialize

import (
	"encoding/binary"
	"fmt"
	"math"
	serializable "my_project/in_house_serializable"
)

// main entry
func deSerialize(byteStream *[]byte) any {
	//retHashMap := in_house_hash_map.CreateHashMap()
	valType, valLen := readMetaData((*byteStream)[:serializable.METADATANOBYTES])
	*byteStream = (*byteStream)[serializable.METADATANOBYTES:]
	value := readData(valType, (*byteStream)[:valLen])
	*byteStream = (*byteStream)[valLen:]
	return value
}

func readMetaData(byteStreamChunk []byte) (uint8, uint8) {
	if len(byteStreamChunk) < 2 {
		panic("byteStreamChunk must have at least 2 bytes")
	}
	return byteStreamChunk[0], byteStreamChunk[1]
}

func readData(valType uint8, byteStreamChunk []byte) any {
	switch valType {
	case serializable.INTEGER32VALUETYPE:
		if len(byteStreamChunk) < serializable.INTEGER32BYTELEN {
			panic("not enough bytes for int32")
		}
		val := binary.BigEndian.Uint32(byteStreamChunk[:serializable.INTEGER32BYTELEN])
		return int32(val)

	case serializable.FLOAT32VALUETYPE:
		if len(byteStreamChunk) < serializable.FLOAT32BYTELEN {
			panic("not enough bytes for float32")
		}
		val := binary.BigEndian.Uint32(byteStreamChunk[:serializable.FLOAT32BYTELEN])
		return math.Float32frombits(val)

	case serializable.STRINGVALUETYPE:
		return string(byteStreamChunk)

	case serializable.LISTVALUETYPE:
		// handled elsewhere by you
		return nil

	default:
		panic(fmt.Sprintf("unknown valueType: %d", valType))
	}
}
