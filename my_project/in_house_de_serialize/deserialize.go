package in_house_de_serialize

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	in_house_errors "my_project/in_house_errors"
	hash_map "my_project/in_house_hash_map"
	serializable "my_project/in_house_serializable"
)

// DeSerialize reads a byteStream of back-to-back (key, value) entries and
// builds a HashMap out of them.  Stops when it sees in_house_errors.EOF.
func DeSerialize(byteStream *[]byte) (*hash_map.HashMap, error) {
	hm := hash_map.CreateHashMap()

	for {
		// read the next item (should be a string key)
		anyKey, err := deSerialize(byteStream)
		if err != nil {
			if errors.Is(err, in_house_errors.EOF) {
				break
			}
			return nil, err
		}

		key, ok := anyKey.(string)
		if !ok {
			return nil, fmt.Errorf("deserialize: expected string key, got %T", anyKey)
		}

		// read the matching value
		anyVal, err := deSerialize(byteStream)
		if err != nil {
			if errors.Is(err, in_house_errors.EOF) {
				// truncated: no value for last key
				break
			}
			return nil, err
		}

		hm.Add(key, anyVal)
	}

	return hm, nil
}

// deSerialize reads exactly one item (key _or_ value) from the stream.
func deSerialize(byteStream *[]byte) (any, error) {
	if len(*byteStream) < serializable.METADATANOBYTES {
		return nil, in_house_errors.EOF
	}

	valType, valLen := readMetaData((*byteStream)[:serializable.METADATANOBYTES])
	*byteStream = (*byteStream)[serializable.METADATANOBYTES:]

	if len(*byteStream) < int(valLen) {
		return nil, in_house_errors.EOF
	}
	chunk := (*byteStream)[:valLen]
	*byteStream = (*byteStream)[valLen:]

	return readData(valType, chunk), nil
}

func readMetaData(b []byte) (uint8, uint8) {
	if len(b) < 2 {
		panic("readMetaData: need at least 2 bytes")
	}
	return b[0], b[1]
}

func readData(valType uint8, b []byte) any {
	switch valType {
	case serializable.INTEGER32VALUETYPE:
		if len(b) < serializable.INTEGER32BYTELEN {
			panic("readData: not enough bytes for int32")
		}
		bits := binary.BigEndian.Uint32(b[:serializable.INTEGER32BYTELEN])
		return int32(bits)

	case serializable.FLOAT32VALUETYPE:
		if len(b) < serializable.FLOAT32BYTELEN {
			panic("readData: not enough bytes for float32")
		}
		bits := binary.BigEndian.Uint32(b[:serializable.FLOAT32BYTELEN])
		return math.Float32frombits(bits)

	case serializable.STRINGVALUETYPE:
		return string(b)

	case serializable.LISTVALUETYPE:
		// handle lists elsewhere
		return nil

	default:
		panic(fmt.Sprintf("readData: unknown valueType %d", valType))
	}
}
