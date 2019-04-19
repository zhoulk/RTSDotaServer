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
			}
			guanKa.Times = guanKa.TotalTimes
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

func (m *Module) FindGuanKas(player *entry.Player, chapterId int32) []*entry.GuanKa {
	res := make([]*entry.GuanKa, 0)
	guanKas := m.AllGuanKas(player)
	if guanKas != nil {
		for _, guanKa := range guanKas {
			if guanKa.ChapterId == chapterId {
				res = append(res, guanKa)
			}
		}
	}
	return res
}

func (m *Module) FindChapter(player *entry.Player, chapterID int32) *entry.Chapter {
	chapters := m.AllChapters(player)
	if chapters != nil {
		for _, chapter := range chapters {
			if chapter.Id == chapterID {
				return chapter
			}
		}
	}
	return nil
}

func (m *Module) FindNextGuanKa(player *entry.Player, guanKaID int32) *entry.GuanKa {
	guanKas := m.AllGuanKas(player)
	if guanKas != nil {
		for _, guanKa := range guanKas {
			if guanKa.Id == guanKaID+1 {
				return guanKa
			}
		}
	}
	return nil
}

// 更新关卡
func (m *Module) UpdateGuanKa(player *entry.Player, gk *entry.GuanKa, result int32) {
	switch result {
	case entry.BattleResultStar0:
		break
	case entry.BattleResultStar1:
		m.updateGuanKa(player, gk, 1)
		break
	case entry.BattleResultStar2:
		m.updateGuanKa(player, gk, 2)
		break
	case entry.BattleResultStar3:
		m.updateGuanKa(player, gk, 3)
		break
	}

	gk.Times = gk.Times - 1
}

func (m *Module) updateGuanKa(player *entry.Player, gk *entry.GuanKa, star int32) {
	gk.Star = star
	gk.Status = entry.ChapterStatusCleared
	chapter := m.FindChapter(player, gk.ChapterId)
	chapter.Status = entry.ChapterStatusCleared
	m.calChapterStar(player, chapter)
	m.openNextGuanKa(player, gk)
}

func (m *Module) openNextGuanKa(player *entry.Player, gk *entry.GuanKa) {
	nextGk := m.FindNextGuanKa(player, gk.Id)
	if nextGk != nil {
		nextGk.IsOpen = true
		nextGk.Status = entry.ChapterStatusNormal
		chapter := m.FindChapter(player, nextGk.ChapterId)
		chapter.Status = entry.ChapterStatusNormal
	}
}

func (m *Module) calChapterStar(player *entry.Player, chapter *entry.Chapter) {
	gks := m.FindGuanKas(player, chapter.Id)

	var stars int32 = 0
	for _, gk := range gks {
		stars += gk.Star
	}
	chapter.Star = stars
}
