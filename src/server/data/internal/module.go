package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"server/base"
	"server/data/entry"

	"github.com/go-simplejson"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
	players     map[string]*entry.Player
	playerHeros map[string][]*entry.Hero
	heros       []*entry.Hero
	skills      []*entry.Skill
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.players = make(map[string]*entry.Player)
	m.playerHeros = make(map[string][]*entry.Hero)
	m.heros = InitHeros()
	m.skills = InitSkills()
}

func (m *Module) OnDestroy() {

}

func (m *Module) FindPlayer(account string, pwd string) *entry.Player {
	return m.players[account+"-"+pwd]
}

func (m *Module) SavePlayer(player *entry.Player) error {
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	m.players[player.Account+"-"+player.Password] = player
	return nil
}

func (m *Module) SavePlayerHero(player *entry.Player, hero *entry.Hero) error {
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	if hero == nil {
		return errors.New("hero is nil")
	}

	heros := m.playerHeros[player.UserId]
	if hero == nil {
		heros = make([]*entry.Hero, 0)
	}
	heros = append(heros, hero)
	m.playerHeros[player.UserId] = heros

	return nil
}

func (m *Module) AllHeros() []*entry.Hero {
	return m.heros
}

func (m *Module) AllOwnHeros(player *entry.Player) ([]*entry.Hero, error) {
	if player == nil || len(player.UserId) == 0 {
		return nil, errors.New("player is nil or userId length is 0")
	}
	return m.playerHeros[player.UserId], nil
}

func InitHeros() []*entry.Hero {
	heros := make([]*entry.Hero, 0)

	data, err := ioutil.ReadFile("data/hero.json")
	if err != nil {
		log.Fatal("%v", err)
	}

	js, err := simplejson.NewJson([]byte(data))
	if err != nil {
		log.Fatal("%v", err)
	}

	m, err := js.Array()
	if err != nil {
		log.Fatal("%v", err)
	}

	for _, v := range m {
		// log.Debug("%v = %v", k, v)

		var hero entry.Hero
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &hero)
		if err != nil {
			log.Debug("err was %v", err)
		}

		heros = append(heros, &hero)
	}

	log.Debug("%v", heros)

	return heros
}

func InitSkills() []*entry.Skill {
	skills := make([]*entry.Skill, 0)

	data, err := ioutil.ReadFile("data/skill.json")
	if err != nil {
		log.Fatal("%v", err)
	}

	js, err := simplejson.NewJson([]byte(data))
	if err != nil {
		log.Fatal("%v", err)
	}

	m, err := js.Array()
	if err != nil {
		log.Fatal("%v", err)
	}

	for _, v := range m {
		// log.Debug("%v = %v", k, v)

		var skill entry.Skill
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &skill)
		if err != nil {
			log.Debug("err was %v", err)
		}

		skills = append(skills, &skill)
	}

	log.Debug("%v", skills)

	return skills
}
