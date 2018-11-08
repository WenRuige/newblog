package main

import (
	"fmt"

	"github.com/newblog/jsonlike/schema"
	"github.com/vmihailenco/msgpack"
)

type Item struct {
	Foo string
}

func testMsgPack() {
	b, err := msgpack.Marshal(&Item{Foo: "food"})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
}

func testMsgPack2JsonEncode() {
	ps := schema.Person{
		Name:    "gewenrui",
		Address: "Harbin",
		Age:     1,
		Hidden:  "1",
	}
	res, err := ps.MarshalMsg(nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}

func testMsgPack2JsonDecode() {
	p := schema.Person{}

	ps := schema.Person{
		Name:    "gewenrui",
		Address: "Harbin",
		Age:     1,
		Hidden:  "1",
	}
	res, err := ps.MarshalMsg(nil)
	if err != nil {
		fmt.Println(err)
	}

	p.UnmarshalMsg([]byte(res))
	fmt.Println(res)
	fmt.Println(p.Name)
	fmt.Println(p.Msgsize())

}

func main() {
	//testMsgPack()
	testMsgPack2JsonDecode()
}
