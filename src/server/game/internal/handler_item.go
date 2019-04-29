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
	allItems := data.Module.AllItems()
	if allItems != nil {
		for _, v := range allItems {
			item := ConverItemToMsgItem(v)
			items = append(items, item)
		}
	}
	response.Items = items

	a.WriteMsg(response)
}

func handleOwnEquip(args []interface{}) {
	log.Debug("game handleOwnEquip")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())
	player := a.UserData().(*entry.Player)

	response := new(msg.EquipResponse)
	response.Code = msg.ResponseCode_SUCCESS

	equips := make([]*msg.Equip, 0)
	ownEquips := data.Module.AllOwnEquips(player)
	if ownEquips != nil {
		for _, v := range ownEquips {
			equip := ConverItemToMsgEquip(v)
			equips = append(equips, equip)
		}
	}
	response.Equips = equips

	a.WriteMsg(response)
}

func ConverItemToMsgItem(v *entry.Item) *msg.Item {
	item := new(msg.Item)
	item.Id = v.Id
	item.Name = v.Name
	item.Price = v.Price
	item.Desc = v.Desc
	return item
}

func ConverItemToMsgEquip(v *entry.Item) *msg.Equip {
	equip := new(msg.Equip)
	equip.Id = v.Id
	equip.Name = v.Name
	equip.Price = v.Price
	equip.Effect = v.Equip.Effect
	equip.Desc = v.Desc
	equip.Mixs = make([]*msg.Mix, 0)
	if v.Equip.Mixs != nil {
		for _, mix := range v.Equip.Mixs {
			equip.Mixs = append(equip.Mixs, ConverMixToMsgMix(mix))
		}
	}
	equip.ItemId = v.ItemId
	equip.HeroId = v.Equip.HeroId
	return equip
}

func ConverMixToMsgMix(v *entry.Mix) *msg.Mix {
	mix := new(msg.Mix)
	mix.ItemId = v.ItemId
	mix.Num = v.Num
	return mix
}
