package define

const (
	LoginRegisteExistErr  int32 = 1
	LoginLoginNotExistErr int32 = 2

	HeroRandomLevelErr   int32 = 100
	HeroRandomGoldErr    int32 = 101
	HeroRandomDiamondErr int32 = 102

	HeroSelectExistErr int32 = 110
	HeroSelectPosErr   int32 = 111

	BattleGuanKaOpenErr  int32 = 200
	BattlePlayerPowerErr int32 = 201
	BattleNoneHeroErr    int32 = 202
)

var ERRMAP map[int32]string = map[int32]string{
	LoginRegisteExistErr:  "用户已存在",
	LoginLoginNotExistErr: "用户不存在",

	HeroRandomLevelErr:   "未知随机类型",
	HeroRandomGoldErr:    "金币不足",
	HeroRandomDiamondErr: "钻石不足",

	HeroSelectExistErr: "改英雄不存在",
	HeroSelectPosErr:   "英雄位置非法",

	BattleGuanKaOpenErr:  "关卡未开启",
	BattlePlayerPowerErr: "体力不足",
	BattleNoneHeroErr:    "没有上阵英雄",
}

const (
	BattleResult_Fail    int32 = 1
	BattleResult_Success int32 = 2
)

func init() {

}
