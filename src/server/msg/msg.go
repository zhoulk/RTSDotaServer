package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

// Processor .
var Processor = protobuf.NewProcessor()

func init() {
	Processor.SetByteOrder(true)
	Processor.Register(100, &LoginRequest{})
	Processor.Register(101, &LoginResponse{})
	Processor.Register(102, &RegisteRequest{})
	Processor.Register(103, &RegisteResponse{})
	Processor.Register(104, &HeroRequest{})
	Processor.Register(105, &HeroResponse{})
	Processor.Register(106, &HeroRandomRequest{})
	Processor.Register(107, &HeroRandomResponse{})
	Processor.Register(108, &HeroOwnRequest{})
	Processor.Register(109, &HeroOwnResponse{})
}
