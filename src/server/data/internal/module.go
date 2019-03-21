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
	items       []*entry.Item
	chapters    []*entry.Chapter
	guanKas     []*entry.GuanKa
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.players = make(map[string]*entry.Player)
	m.playerHeros = make(map[string][]*entry.Hero)
	m.heros = InitHeros()
	m.skills = InitSkills()
	m.items = InitItems()
	m.chapters = InitChapters()
	m.guanKas = InitGuanKas()
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

func (m *Module) AllSkills() []*entry.Skill {
	return m.skills
}

func (m *Module) AllItems() []*entry.Item {
	return m.items
}

func (m *Module) AllChapters() []*entry.Chapter {
	return m.chapters
}

func (m *Module) AllGuanKas() []*entry.GuanKa {
	return m.guanKas
}

func (m *Module) AllOwnHeros(player *entry.Player) ([]*entry.Hero, error) {
	if player == nil || len(player.UserId) == 0 {
		return nil, errors.New("player is nil or userId length is 0")
	}
	return m.playerHeros[player.UserId], nil
}

func InitHeros() []*entry.Hero {
	heros := make([]*entry.Hero, 0)

	m, _ := ReadFile("data/json/hero.json")

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

	m, _ := ReadFile("data/json/skill.json")

	for _, v := range m {
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

func InitItems() []*entry.Item {
	items := make([]*entry.Item, 0)

	m, _ := ReadFile("data/json/item.json")

	for _, v := range m {
		var item entry.Item
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &item)
		if err != nil {
			log.Debug("err was %v", err)
		}

		items = append(items, &item)
	}

	log.Debug("%v", items)

	return items
}

func InitChapters() []*entry.Chapter {
	chapters := make([]*entry.Chapter, 0)

	m, _ := ReadFile("data/json/chapter.json")

	for _, v := range m {
		var chapter entry.Chapter
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &chapter)
		if err != nil {
			log.Debug("err was %v", err)
		}

		chapters = append(chapters, &chapter)
	}

	log.Debug("%v", chapters)

	return chapters
}

func InitGuanKas() []*entry.GuanKa {
	guanKas := make([]*entry.GuanKa, 0)

	m, _ := ReadFile("data/json/guanKa.json")

	for _, v := range m {
		var guanKa entry.GuanKa
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &guanKa)
		if err != nil {
			log.Debug("err was %v", err)
		}

		guanKas = append(guanKas, &guanKa)
	}

	log.Debug("%v", guanKas)

	return guanKas
}

func ReadFile(path string) ([]interface{}, error) {
	data, err := ioutil.ReadFile(path)
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

	return m, err
}
