package main

import (
	"fmt"
	pb "k8s-protos-issue/proto" // protoc -Iproto -I$GOPATH/src --go_out=proto proto/embedded_deployment.proto

	"github.com/golang/protobuf/proto"
	"k8s.io/api/apps/v1"
)

func main() {
	deployment := &v1.Deployment{}
	marshalled, err := deployment.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Printf("deployment.Marshal(): %v\n", marshalled)

	marshalled2, err := proto.Marshal(deployment)
	fmt.Printf("proto.Marshal(deployment): %v\n", marshalled2)

	embededDeployment := &pb.EmbeddedDeployment{Foobar: "foobar", Deployment: deployment}
	marshalled3, err := proto.Marshal(embededDeployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("proto.Marshal(embededDeployment): %v\n", marshalled3)
}
