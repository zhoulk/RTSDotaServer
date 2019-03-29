package internal

import (
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleAllSkill(args []interface{}) {
	log.Debug("game handleAllSkill")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())

	response := new(msg.SkillResponse)
	response.Code = msg.ResponseCode_SUCCESS

	skills := make([]*msg.Skill, 0)
	for _, v := range data.Module.AllSkills() {
		skill := ConverSkillToMsgSkill(v)
		skills = append(skills, skill)
	}
	response.Skills = skills

	a.WriteMsg(response)
}

func handleSkillUpgrade(args []interface{}) {
	log.Debug("game handleSkillUpgrade")

	m := args[0].(*msg.SkillUpgradeRequest)
	a := args[1].(gate.Agent)

	skillId := m.SkillId
	player := a.UserData().(*entry.Player)

	hero, skill := data.Module.FindHeroSkill(player, skillId)
	if skill == nil {
		response := new(msg.SkillUpgradeResponse)
		response.Code = msg.ResponseCode_FAIL
		err := new(msg.Error)
		err.Code = define.SkillUpgradeExistErr
		err.Msg = define.ERRMAP[err.Code]
		response.Err = err
		a.WriteMsg(response)
		return
	}

	if hero.SkillPoint == 0 {
		response := new(msg.SkillUpgradeResponse)
		response.Code = msg.ResponseCode_FAIL
		err := new(msg.Error)
		err.Code = define.SkillUpgradeSPErr
		err.Msg = define.ERRMAP[err.Code]
		response.Err = err
		a.WriteMsg(response)
		return
	}

	data.Module.SkillUpgrade(hero, skill)

	response := new(msg.SkillUpgradeResponse)
	response.Code = msg.ResponseCode_SUCCESS
	a.WriteMsg(response)
}

func ConverSkillToMsgSkill(v *entry.Skill) *msg.Skill {
	skill := new(msg.Skill)
	skill.Id = v.Id
	skill.Name = v.Name
	skill.Level = v.Level
	skill.Type = v.Type
	skill.Desc = v.Desc
	skill.IsOpen = v.IsOpen
	skill.SkillId = v.SkillId
	skill.HeroId = v.HeroId
	return skill
}
