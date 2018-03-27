package protobuffencoder

import(
	"errors"
	// "github.com/golang/protobuf/proto"
)

// commented out for the time being to stop errors
// needs to be able to encode from a type
// dont understand this well enough to fully implement
func EncodeSomeProtoBuf(obj interface{}) ([]byte, error){
	// if v, ok := obj.(); ok {
	// 	return proto.Marshal(v)
	// }
	return nil, errors.New("Proto: Unknown message type")
}