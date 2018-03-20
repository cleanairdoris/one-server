package main

import (
	"log"

	proto "github.com/golang/protobuf/proto"
)

/*
TestPro get protobuf data
*/
func TestPro() (ret []byte) {
	// 为 AllPerson 填充数据
	p1 := Person{
		Id:   *proto.Int32(1),
		Name: *proto.String("xieyanke"),
	}

	p2 := Person{
		Id:   2,
		Name: "gopher",
	}

	allP := AllPerson{
		Per: []*Person{&p1, &p2},
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&allP)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}

	// 对已经序列化的数据进行反序列化
	var target AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("UnMashal data error:", err)
	}

	println(target.Per[0].Name) // 打印第一个 person Name 的值进行反序列化验证
	return data
}
