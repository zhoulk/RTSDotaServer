package define

const (
	HeroRandomLevelErr   int32 = 100
	HeroRandomGoldErr    int32 = 101
	HeroRandomDiamondErr int32 = 102
)

var ERRMAP map[int32]string = map[int32]string{
	HeroRandomLevelErr:   "未知随机类型",
	HeroRandomGoldErr:    "金币不足",
	HeroRandomDiamondErr: "钻石不足",
}

func init() {

}
