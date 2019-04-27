package entry

import (
	"fmt"
)

const (
	ChapterStatusLock    int32 = 1
	ChapterStatusNormal  int32 = 2
	ChapterStatusCleared int32 = 3
)

const (
	BattleResultStar1 int32 = 1
	BattleResultStar2 int32 = 2
	BattleResultStar3 int32 = 3
	BattleResultStar0 int32 = 4
)

type Chapter struct {
	Id        int32
	Name      string
	Star      int32
	Status    int32
	GuanKaNum int32
	IsOpen    bool

	ChapterId string
	IsDirty   bool
}

type GuanKa struct {
	Id         int32
	Name       string
	ChapterId  int32
	Heros      []*Hero
	Earn       *Earn
	Expend     *Expend
	IsOpen     bool
	Star       int32
	Status     int32
	Times      int32
	TotalTimes int32

	GuanKaId string
	IsDirty  bool
}

type Earn struct {
	ItemIds   []int32
	Items     []*Item
	HeroExp   int32
	PlayerExp int32
	Gold      int32
}

type Expend struct {
	Power int32
}

func NewChapter() *Chapter {
	chapter := new(Chapter)
	chapter.IsDirty = true
	return chapter
}

func (c *Chapter) SetChapterId(chapterId string) {
	c.ChapterId = chapterId
	c.IsDirty = true
}

func (c *Chapter) SetOpen(open bool) {
	c.IsOpen = open
	c.IsDirty = true
}

func (c *Chapter) SetStatus(status int32) {
	c.Status = status
	c.IsDirty = true
}

func (c *Chapter) SetStar(star int32) {
	c.Star = star
	c.IsDirty = true
}

func NewGuanKa() *GuanKa {
	gk := new(GuanKa)
	gk.IsDirty = true
	return gk
}

func (g *GuanKa) SetGuanKaId(gkId string) {
	g.GuanKaId = gkId
	g.IsDirty = true
}

func (g *GuanKa) SetOpen(open bool) {
	g.IsOpen = open
	g.IsDirty = true
}

func (g *GuanKa) SetStatus(status int32) {
	g.Status = status
	g.IsDirty = true
}

func (g *GuanKa) SetStar(star int32) {
	g.Star = star
	g.IsDirty = true
}

func (g *GuanKa) SetTimes(times int32) {
	g.Times = times
	g.IsDirty = true
}

func (c *Chapter) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, IsOpen : %v}",
		c.Id, c.Name, c.IsOpen)
}

func (g *GuanKa) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, ChapterId : %d, IsOpen : %v, Earn : %v, Expend : %v}",
		g.Id, g.Name, g.ChapterId, g.IsOpen, g.Earn, g.Expend)
}

func (e *Earn) String() string {
	return fmt.Sprintf("\n{ItemIds : %v, HeroExp : %d, PlayerExp : %d, Gold : %d}",
		e.ItemIds, e.HeroExp, e.PlayerExp, e.Gold)
}

func (e *Expend) String() string {
	return fmt.Sprintf("\n{Power : %d}",
		e.Power)
}
