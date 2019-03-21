package entry

import "fmt"

type Chapter struct {
	Id   int32
	Name string
}

type GuanKa struct {
	Id        int32
	Name      string
	ChapterId int32
	Earn      *Earn
	Expend    *Expend
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

func (c *Chapter) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s}",
		c.Id, c.Name)
}

func (g *GuanKa) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, ChapterId : %d, Earn : %v, Expend : %v}",
		g.Id, g.Name, g.ChapterId, g.Earn, g.Expend)
}

func (e *Earn) String() string {
	return fmt.Sprintf("\n{ItemIds : %v, HeroExp : %d, PlayerExp : %d, Gold : %d}",
		e.ItemIds, e.HeroExp, e.PlayerExp, e.Gold)
}

func (e *Expend) String() string {
	return fmt.Sprintf("\n{Power : %d}",
		e.Power)
}
