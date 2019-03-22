package entry

import "fmt"

type Chapter struct {
	Id     int32
	Name   string
	IsOpen bool
}

type GuanKa struct {
	Id        int32
	Name      string
	ChapterId int32
	Earn      *Earn
	Expend    *Expend
	IsOpen    bool
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
