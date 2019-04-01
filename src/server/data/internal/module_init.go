package internal

import (
	"encoding/json"
	"io/ioutil"
	"server/data/entry"

	"github.com/go-simplejson"
	"github.com/name5566/leaf/log"
)

func InitHeros() []*entry.Hero {
	heros := make([]*entry.Hero, 0)

	m, _ := ReadFile("data/json/hero.json")

	for _, v := range m {
		// log.Debug("%v = %v", k, v)

		var hero *entry.Hero
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &hero)
		if err != nil {
			log.Debug("err was %v", err)
		}

		hero.Exp = 0
		hero.Level = 1
		hero.LevelUpExp = 90
		hero.SkillPoint = 1
		hero.MaxBlood = hero.Blood
		hero.MaxMP = hero.MP
		heros = append(heros, hero)
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
