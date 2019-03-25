package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"server/base"
	"server/data/entry"
	"server/tool"

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
	players        map[string]*entry.Player
	playerHeros    map[string][]*entry.Hero
	heros          []*entry.Hero
	skills         []*entry.Skill
	items          []*entry.Item
	chapters       []*entry.Chapter
	playerChapters map[string][]*entry.Chapter
	guanKas        []*entry.GuanKa
	playerGuanKas  map[string][]*entry.GuanKa
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.players = make(map[string]*entry.Player)
	m.playerHeros = make(map[string][]*entry.Hero)
	m.playerChapters = make(map[string][]*entry.Chapter)
	m.playerGuanKas = make(map[string][]*entry.GuanKa)
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
	if hero.Skills == nil {
		hero.Skills = m.FindHeroSkills(hero)
		if len(hero.Skills) > 0 {
			hero.Skills[0].IsOpen = true
		}
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

func (m *Module) AllChapters(player *entry.Player) []*entry.Chapter {
	if player == nil || len(player.UserId) == 0 {
		return nil
	}
	if m.playerChapters[player.UserId] == nil {
		chapters := make([]*entry.Chapter, 0)
		ch := m.FindChapterDefine(1)
		chapter := new(entry.Chapter)
		tool.DeepCopy(chapter, ch)
		chapter.IsOpen = true
		chapters = append(chapters, chapter)
		m.playerChapters[player.UserId] = chapters
	}

	return m.playerChapters[player.UserId]
}

func (m *Module) AllGuanKas(player *entry.Player) []*entry.GuanKa {
	if player == nil || len(player.UserId) == 0 {
		return nil
	}
	if m.playerGuanKas[player.UserId] == nil {
		guanKas := make([]*entry.GuanKa, 0)
		gk := m.FindGuanKaDefine(1)
		guanKa := new(entry.GuanKa)
		tool.DeepCopy(guanKa, gk)
		guanKa.IsOpen = true
		guanKas = append(guanKas, guanKa)
		m.playerGuanKas[player.UserId] = guanKas
	}

	return m.playerGuanKas[player.UserId]
}

func (m *Module) AllOwnHeros(player *entry.Player) ([]*entry.Hero, error) {
	if player == nil || len(player.UserId) == 0 {
		return nil, errors.New("player is nil or userId length is 0")
	}
	return m.playerHeros[player.UserId], nil
}

func (m *Module) FindAHero(player *entry.Player, heroId string) *entry.Hero {
	heros, _ := m.AllOwnHeros(player)
	if heros != nil {
		for _, hero := range heros {
			if hero.HeroId == heroId {
				return hero
			}
		}
	}
	return nil
}

func (m *Module) FindAHeroAt(player *entry.Player, pos int32) *entry.Hero {
	heros, _ := m.AllOwnHeros(player)
	if heros != nil {
		for _, hero := range heros {
			if hero.Pos == pos {
				return hero
			}
		}
	}
	return nil
}

func (m *Module) FindHeroSkills(hero *entry.Hero) []*entry.Skill {
	skills := make([]*entry.Skill, 0)
	if hero.SkillIds != nil {
		for _, skillId := range hero.SkillIds {
			skill := m.FindASkill(skillId)
			if skill != nil {
				skills = append(skills, skill)
			} else {
				log.Error("[FindHeroSkills ] skill is not exist , skillId = %v", skillId)
			}
		}
	}
	return skills
}

func (m *Module) FindASkill(skillId int32) *entry.Skill {
	skills := m.AllSkills()
	for _, skill := range skills {
		if skillId == skill.Id {
			return skill
		}
	}
	return nil
}

func (m *Module) RemoveHero(player *entry.Player, heroId string) *entry.Hero {
	heros, _ := m.AllOwnHeros(player)
	var oldHero *entry.Hero
	if heros != nil {
		target := heros[:0]
		for _, hero := range heros {
			if hero.HeroId != heroId {
				target = append(target, hero)
			} else {
				oldHero = hero
			}
		}
	}
	return oldHero
}

func (m *Module) UnSelectHero(player *entry.Player, hero *entry.Hero) {
	hero.IsSelect = false
	hero.Pos = 0
}

func (m *Module) SelectHero(player *entry.Player, hero *entry.Hero, pos int32) {
	hero.IsSelect = true
	hero.Pos = pos
}

func (m *Module) SelectHeroIds(player *entry.Player) []string {
	heros, _ := m.AllOwnHeros(player)
	heroIds := make([]string, 0)
	if heros != nil {
		for _, hero := range heros {
			if hero.IsSelect {
				heroIds = append(heroIds, hero.HeroId)
			}
		}
	}
	return heroIds
}

func (m *Module) SelectHeros(player *entry.Player) []*entry.Hero {
	heros, _ := m.AllOwnHeros(player)
	selelctHeros := make([]*entry.Hero, 0)
	if heros != nil {
		for _, hero := range heros {
			if hero.IsSelect {
				selelctHeros = append(selelctHeros, hero)
			}
		}
	}
	return selelctHeros
}

func (m *Module) FindChapterDefine(chapterId int32) *entry.Chapter {
	for _, chapter := range m.chapters {
		if chapter.Id == chapterId {
			return chapter
		}
	}
	return nil
}

func (m *Module) FindGuanKaDefine(guanKaId int32) *entry.GuanKa {
	for _, guanKa := range m.guanKas {
		if guanKa.Id == guanKaId {
			return guanKa
		}
	}
	return nil
}

func (m *Module) FindGuanKa(player *entry.Player, guanKaId int32) *entry.GuanKa {
	guanKas := m.AllGuanKas(player)
	if guanKas != nil {
		for _, guanKa := range guanKas {
			if guanKa.Id == guanKaId {
				return guanKa
			}
		}
	}
	return nil
}

func (m *Module) EffectByEarn(player *entry.Player, earn *entry.Earn) {

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
