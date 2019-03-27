package internal

import (
	"math/rand"
	"server/data/entry"
)

func CalGuanKaEarn(gk *entry.GuanKa, res bool) *entry.Earn {
	earn := new(entry.Earn)
	if res {
		earn.HeroExp = gk.Earn.HeroExp
		earn.Gold = gk.Earn.Gold
		earn.ItemIds = make([]int32, 0)
		index := rand.Intn(len(gk.Earn.ItemIds))
		earn.ItemIds = append(earn.ItemIds, gk.Earn.ItemIds[index])
	}
	earn.PlayerExp = gk.Expend.Power * 20
	return earn
}
