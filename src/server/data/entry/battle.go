package entry

const (
	BattleTypeGuanKa int32 = 1
)

type BattleInfo struct {
	BattleId string
	Type     int32
	Guanka   *GuanKa
}
