package entry

import "fmt"

const (
	ItemTypeEquip    int32 = 1
	ItemTypeConsume  int32 = 2
	ItemTypeHeroChip int32 = 3
)

type Item struct {
	Id       int32
	Type     int32
	Name     string
	Price    int32
	Desc     string
	Equip    *Equip
	HeroChip *HeroChip

	ItemId string

	IsDirty bool
}

type Equip struct {
	HeroId string
	Effect string
	Mixs   []*Mix
}

type Mix struct {
	ItemId int32
	Num    int32
	Item   *Item
}

type Consume struct {
}

type HeroChip struct {
	Cnt        int32
	ComposeCnt int32
}

func NewItem() *Item {
	item := new(Item)
	item.IsDirty = true
	return item
}
func (m *Item) SetItemId(itemId string) {
	m.ItemId = itemId
	m.IsDirty = true
}

func (item *Item) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, Price : %d, Effect : %s, Desc : %s, Mixs : %v}",
		item.Id, item.Name, item.Price, "", "", "")
}

func (m *Mix) String() string {
	return fmt.Sprintf("\n		{ItemId : %d, Num : %d}",
		m.ItemId, m.Num)
}
