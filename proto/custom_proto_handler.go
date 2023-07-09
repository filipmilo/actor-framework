package proto

import (
	"log"
	reflect "reflect"

	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func Encode(msg interface{}) *ProtoEnvelope {
	bytes, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		return nil
	}

	data := &any.Any{
		Value: bytes,
	}

	envelope := &ProtoEnvelope{
		Message: data,
	}
	envelope.Message.TypeUrl = reflect.TypeOf(msg).String()

	return envelope
}

func Decode(protoEnvelope *ProtoEnvelope) interface{} {

	n, _ := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(protoEnvelope.Message.TypeUrl)[1:])
	pm := n.New().Interface()
	err := proto.Unmarshal(protoEnvelope.Message.Value, pm)
	if err != nil {
		log.Fatal("Error while deserializing Ping message:", err)
	}

	return pm
}
