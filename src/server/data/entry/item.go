package entry

import "fmt"

type Item struct {
	Id     int32
	Name   string
	Price  int32
	Effect string
	Desc   string
	Mixs   []*Mix
}

type Mix struct {
	ItemId int32
	Num    int32
	Item   *Item
}

func (item *Item) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, Price : %d, Effect : %s, Desc : %s, Mixs : %v}",
		item.Id, item.Name, item.Price, item.Effect, "", item.Mixs)
}

func (m *Mix) String() string {
	return fmt.Sprintf("\n		{ItemId : %d, Num : %d}",
		m.ItemId, m.Num)
}
