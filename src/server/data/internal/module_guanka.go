package internal

import (
	"server/data/entry"
	"server/tool"
)

// AllChapters ...
func (m *Module) AllChapters(player *entry.Player) []*entry.Chapter {
	if player == nil || len(player.UserId) == 0 {
		return nil
	}
	if m.playerChapters[player.UserId] == nil {
		chapters := make([]*entry.Chapter, 0)
		for i, ch := range m.chapters {
			chapter := new(entry.Chapter)
			tool.DeepCopy(chapter, ch)
			chapter.Status = entry.ChapterStatusLock
			if i == 0 {
				chapter.IsOpen = true
				chapter.Status = entry.ChapterStatusNormal
			}
			chapters = append(chapters, chapter)
		}
		m.playerChapters[player.UserId] = chapters
	}

	return m.playerChapters[player.UserId]
}

// AllGuanKas ...
func (m *Module) AllGuanKas(player *entry.Player) []*entry.GuanKa {
	if player == nil || len(player.UserId) == 0 {
		return nil
	}
	if m.playerGuanKas[player.UserId] == nil {
		guanKas := make([]*entry.GuanKa, 0)

		for i, gk := range m.guanKas {
			guanKa := new(entry.GuanKa)
			tool.DeepCopy(guanKa, gk)
			guanKa.Status = entry.ChapterStatusLock
			if i == 0 {
				guanKa.IsOpen = true
				guanKa.Status = entry.ChapterStatusNormal
				guanKa.Times = guanKa.TotalTimes
			}
			guanKas = append(guanKas, guanKa)
		}
		m.playerGuanKas[player.UserId] = guanKas
	}

	return m.playerGuanKas[player.UserId]
}

// FindChapterDefine ...
func (m *Module) FindChapterDefine(chapterID int32) *entry.Chapter {
	for _, chapter := range m.chapters {
		if chapter.Id == chapterID {
			return chapter
		}
	}
	return nil
}

// FindGuanKaDefine ...
func (m *Module) FindGuanKaDefine(guanKaID int32) *entry.GuanKa {
	for _, guanKa := range m.guanKas {
		if guanKa.Id == guanKaID {
			return guanKa
		}
	}
	return nil
}

// FindGuanKa ...
func (m *Module) FindGuanKa(player *entry.Player, guanKaID int32) *entry.GuanKa {
	guanKas := m.AllGuanKas(player)
	if guanKas != nil {
		for _, guanKa := range guanKas {
			if guanKa.Id == guanKaID {
				return guanKa
			}
		}
	}
	return nil
}
