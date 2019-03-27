package internal

import (
	"server/data"
	"server/data/entry"
	"server/msg"
	"server/tool"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleAllChapter(args []interface{}) {
	log.Debug("game handleAllChapter")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	player := a.UserData().(*entry.Player)

	response := new(msg.ChapterResponse)
	response.Code = msg.ResponseCode_SUCCESS

	chapters := make([]*msg.Chapter, 0)
	for _, v := range data.Module.AllChapters(player) {
		item := ConverChapterToMsgChapter(v)
		chapters = append(chapters, item)
	}
	response.Chapters = chapters

	a.WriteMsg(response)
}

func handleAllGuanKa(args []interface{}) {
	log.Debug("game handleAllGuanKa")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	player := a.UserData().(*entry.Player)

	response := new(msg.GuanKaResponse)
	response.Code = msg.ResponseCode_SUCCESS

	guanKas := make([]*msg.GuanKa, 0)
	for _, v := range data.Module.AllGuanKas(player) {
		guanKa := ConverGuanKaToMsgGuanKa(v)
		guanKas = append(guanKas, guanKa)
	}
	response.GuanKas = guanKas

	a.WriteMsg(response)
}

func ConverChapterToMsgChapter(v *entry.Chapter) *msg.Chapter {
	chapter := new(msg.Chapter)
	chapter.Id = v.Id
	chapter.Name = v.Name
	chapter.IsOpen = v.IsOpen
	return chapter
}

func ConverGuanKaToMsgGuanKa(v *entry.GuanKa) *msg.GuanKa {
	guanKa := new(msg.GuanKa)
	guanKa.Id = v.Id
	guanKa.Name = v.Name
	guanKa.ChapterId = v.ChapterId
	guanKa.IsOpen = v.IsOpen
	guanKa.Earn = ConverEarnToMsgEarn(v.Earn)
	guanKa.Expend = ConverExpendToMsgExpend(v.Expend)
	return guanKa
}

func ConverEarnToMsgEarn(v *entry.Earn) *msg.Earn {
	eran := new(msg.Earn)
	for _, id := range v.ItemIds {
		eran.ItemIds = append(eran.ItemIds, tool.String(id))
	}
	eran.HeroExp = v.HeroExp
	eran.PlayerExp = v.PlayerExp
	eran.Gold = v.Gold
	return eran
}

func ConverExpendToMsgExpend(v *entry.Expend) *msg.Expend {
	expend := new(msg.Expend)
	expend.Power = v.Power
	return expend
}
