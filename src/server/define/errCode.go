package define

const (
	SysRequestArgsErr int32 = 1

	LoginRegisteExistErr     int32 = 501
	LoginLoginNotExistErr    int32 = 502
	LoginRegisteAccountExist int32 = 503

	HeroRandomLevelErr   int32 = 100
	HeroRandomGoldErr    int32 = 101
	HeroRandomDiamondErr int32 = 102

	HeroSelectExistErr int32 = 110
	HeroSelectPosErr   int32 = 111

	BattleGuanKaOpenErr  int32 = 200
	BattlePlayerPowerErr int32 = 201
	BattleNoneHeroErr    int32 = 202
	BattleResultExistErr int32 = 203
	BattleGuanKaTimesErr int32 = 204

	SkillUpgradeExistErr int32 = 300
	SkillUpgradeSPErr    int32 = 301

	GroupCreateDiamondErr  int32 = 400
	GroupCreateExistErr    int32 = 401
	GroupApplyExistErr     int32 = 402
	GroupApplyIsInErr      int32 = 403
	GroupOperExistErr      int32 = 404
	GroupOperIsInErr       int32 = 405
	GroupOperNoApplyErr    int32 = 406
	GroupOperMemberFullErr int32 = 407
	GroupOperNoLeaderErr   int32 = 408
	GroupOperIsNotInErr    int32 = 409
)

var ERRMAP map[int32]string = map[int32]string{
	SysRequestArgsErr: "请求参数异常",

	LoginRegisteExistErr:     "用户已存在",
	LoginLoginNotExistErr:    "用户不存在",
	LoginRegisteAccountExist: "账号已存在",

	HeroRandomLevelErr:   "未知随机类型",
	HeroRandomGoldErr:    "金币不足",
	HeroRandomDiamondErr: "钻石不足",

	HeroSelectExistErr: "改英雄不存在",
	HeroSelectPosErr:   "英雄位置非法",

	BattleGuanKaOpenErr:  "关卡未开启",
	BattlePlayerPowerErr: "体力不足",
	BattleNoneHeroErr:    "没有上阵英雄",
	BattleResultExistErr: "不存在的战斗",
	BattleGuanKaTimesErr: "没有战斗次数",

	SkillUpgradeExistErr: "技能不存在",
	SkillUpgradeSPErr:    "技能点不足",

	GroupCreateDiamondErr:  "钻石不足",
	GroupCreateExistErr:    "已经拥有一个军团",
	GroupApplyExistErr:     "军团不存在",
	GroupApplyIsInErr:      "你已经加入该军团",
	GroupOperExistErr:      "军团不存在",
	GroupOperIsInErr:       "你已经加入该军团",
	GroupOperNoApplyErr:    "申请不存在",
	GroupOperMemberFullErr: "军团已经满员",
	GroupOperNoLeaderErr:   "没有操作权限",
	GroupOperIsNotInErr:    "还未加入该军团",
}

const (
	BattleResult_Fail    int32 = 1
	BattleResult_Success int32 = 2
)

func init() {

}
