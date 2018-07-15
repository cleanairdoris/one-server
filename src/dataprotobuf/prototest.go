package dataprotobuf

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func GetProtoData() (ret []byte) {
	t1 := &Title{
		Id:   1,
		Name: "kexue",
	}
	t2 := &Title{
		Id:   2,
		Name: "yulele", //EgkIARIFa2V4dWUSCAgCEgR5dWxl
	}
	ts := &Titlelist{}
	ts.Rtncode = 0
	ts.Titles = append(ts.Titles, t1)
	ts.Titles = append(ts.Titles, t2)

	data, _ := proto.Marshal(ts)

	var aaa Titlelist
	err := proto.Unmarshal(data, &aaa)
	if err == nil {
		fmt.Println(aaa.Rtncode, aaa.Titles[0].Name)
	}

	return data
}
