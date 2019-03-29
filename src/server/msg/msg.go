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
	Processor.Register(110, &SkillRequest{})
	Processor.Register(111, &SkillResponse{})
	Processor.Register(112, &ItemRequest{})
	Processor.Register(113, &ItemResponse{})
	Processor.Register(114, &ChapterRequest{})
	Processor.Register(115, &ChapterResponse{})
	Processor.Register(116, &GuanKaRequest{})
	Processor.Register(117, &GuanKaResponse{})
	Processor.Register(118, &BattleGuanKaRequest{})
	Processor.Register(119, &BattleGuanKaResponse{})
	Processor.Register(120, &HeroSelectRequest{})
	Processor.Register(121, &HeroSelectResponse{})
	Processor.Register(122, &HeroUnSelectRequest{})
	Processor.Register(123, &HeroUnSelectResponse{})
	Processor.Register(124, &HeroSkillsRequest{})
	Processor.Register(125, &HeroSkillsResponse{})
	Processor.Register(126, &HeroItemsRequest{})
	Processor.Register(127, &HeroItemsResponse{})
	Processor.Register(128, &SkillUpgradeRequest{})
	Processor.Register(129, &SkillUpgradeResponse{})
	Processor.Register(130, &ZoneRequest{})
	Processor.Register(131, &ZoneResponse{})
}
