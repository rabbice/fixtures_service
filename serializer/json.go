package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers: false,
		EmitUnpopulated: true,
		Indent: " ",
		UseProtoNames: true,
	}
	return marshaler.Format(message), nil
}
