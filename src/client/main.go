package main

import (
	"encoding/binary"
	"net"
	"server/msg"

	"github.com/golang/protobuf/proto"

	"github.com/name5566/leaf/log"

	"github.com/name5566/leaf/network/protobuf"
)

func main() {
	log.Debug("conneting ...")
	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		panic(err)
	}

	var Processor = protobuf.NewProcessor()
	Processor.Register(100, &msg.LoginRequest{})

	data, err := Processor.Marshal(&msg.LoginRequest{
		Account:  "name",
		Password: "pwd",
	})
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// len + id + data
	m := make([]byte, 4+len(data[1]))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(2+len(data[1])))

	copy(m[2:], data[0])
	copy(m[4:], data[1])

	// 发送消息
	conn.Write(m)

	buf := make([]byte, 32)
	// 接收消息
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("read error:", err)
	}

	recv := &msg.LoginResponse{}
	err = proto.Unmarshal(buf[4:n], recv)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Debug(recv.GetUid())
}
