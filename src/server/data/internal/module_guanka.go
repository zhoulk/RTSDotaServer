package internal

import (
	"server/data/entry"
	"server/tool"
)

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
