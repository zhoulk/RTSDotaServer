package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/log"
)

func init() {
	log.Debug("game init")
	handler(&msg.LoginRequest{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// log.Debug("game handleHello")

	// // 收到的 Hello 消息
	// m := args[0].(*msg.SearchRequest)
	// // 消息的发送者
	// a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	// log.Debug("hello %v %v %v", m.GetPageNumber(), m.GetResultPerPage(), m.GetQuery())

	// // 给发送者回应一个 Hello 消息
	// a.WriteMsg(&msg.SearchRequest{
	// 	Query:         "client",
	// 	PageNumber:    100,
	// 	ResultPerPage: 10,
	// })
}
