package internal

import (
	"server/data"
	"server/data/entry"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleAllItem(args []interface{}) {
	log.Debug("game handleAllItem")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())

	response := new(msg.ItemResponse)
	response.Code = msg.ResponseCode_SUCCESS

	items := make([]*msg.Item, 0)
	for _, v := range data.Module.AllItems() {
		item := ConverItemToMsgItem(v)
		items = append(items, item)
	}
	response.Items = items

	a.WriteMsg(response)
}

func ConverItemToMsgItem(v *entry.Item) *msg.Item {
	item := new(msg.Item)
	item.Id = v.Id
	item.Name = v.Name
	item.Price = v.Price
	// item.Effect = v.Effect
	item.Desc = v.Desc
	// item.Mixs = make([]*msg.Mix, 0)
	// if v.Mixs != nil {
	// 	for _, mix := range v.Mixs {
	// 		item.Mixs = append(item.Mixs, ConverMixToMsgMix(mix))
	// 	}
	// }
	item.ItemId = v.ItemId
	item.HeroId = v.HeroId
	return item
}

func ConverMixToMsgMix(v *entry.Mix) *msg.Mix {
	mix := new(msg.Mix)
	mix.ItemId = v.ItemId
	mix.Num = v.Num
	return mix
}
