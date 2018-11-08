package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/newblog/jsonlike/protobuf"
)

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
	}

	data, err := proto.Marshal(test)
	fmt.Println(string(data))
	if err != nil {
	}

	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
