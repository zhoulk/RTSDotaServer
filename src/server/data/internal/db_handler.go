package internal

import (
	"server/data/entry"
	"server/tool"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/name5566/leaf/log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_Driver = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/pixel_farm?charset=utf8&&parseTime=true"
)

func (m *Module) PersistentData() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "dota_" + defaultTableName
	}

	m.ConnectDB()
	m.CreateTables()
	m.IntializeTables()

	m.LoadFromDB()
}

func (m *Module) ConnectDB() {
	db, err := gorm.Open("mysql", DB_Driver)
	if err != nil {
		log.Error("%v", err)
		panic("failed to connect database")
	}
	// defer db.Close()
	m.db = db
}

func (m *Module) CreateTables() {
	if !m.db.HasTable(&Zone{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Zone{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&User{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserBaseInfo{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserBaseInfo{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&HeroDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&SkillDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&SkillDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ChapterDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ChapterDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&GuanKaDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&GuanKaDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&GuanKaEarnDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&GuanKaEarnDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&GuanKaExpendDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&GuanKaExpendDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&GuanKaHeroDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&GuanKaHeroDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&EquipDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&EquipDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&EquipMixDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&EquipMixDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ConsumeDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ConsumeDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&HeroChipDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroChipDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ExpHeroDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ExpHeroDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ExpPlayerDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ExpPlayerDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserTarvern{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserTarvern{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserHero{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserHero{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserHeroSkill{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserHeroSkill{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserChapter{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserChapter{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserGuanKa{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserGuanKa{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&Group{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Group{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&GroupMember{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&GroupMember{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserEquip{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserEquip{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserConsume{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserConsume{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserHeroChip{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserHeroChip{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&GameConfig{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&GameConfig{}).Error; err != nil {
			panic(err)
		}
	}
}

func (m *Module) IntializeTables() {
	m.IntializeZoneTable()
	m.IntializeHeroTable()
	m.IntializeSkillTable()
	m.IntializeChapterTable()
	m.IntializeGuanKaTable()
	m.IntializeEquipTable()
	m.IntializeExpPlayerTable()
	m.IntializeExpHeroTable()
	m.InitalizeGameConfig()
}

func (m *Module) InitalizeGameConfig() {
	config := GameConfig{
		ConfigKey:   entry.GameConfigKey_GroupPrice,
		ConfigValue: "500",
	}
	m.db.Where("config_key = ?", config.ConfigKey).FirstOrCreate(&config)
}

func (m *Module) IntializeZoneTable() {
	zones := entry.ZoneList
	for _, zone := range zones {
		z := Zone{
			ZoneId:     zone.Id,
			TCPAddr:    zone.TCPAddr,
			MaxConnNum: zone.MaxConnNum,
			Name:       zone.Name,
			IsNew:      zone.IsNew,
		}
		m.db.Where("zone_id = ?", z.ZoneId).FirstOrCreate(&z)
	}
}

func (m *Module) IntializeHeroTable() {
	heros := InitHeros()
	for _, hero := range heros {
		skillStr := ""
		if hero.SkillIds != nil {
			skillStr = tool.JoinInt32Arr(hero.SkillIds, ",")
		}
		heroDefine := HeroDefine{
			HeroId:           hero.Id,
			Name:             hero.Name,
			Type:             hero.Type,
			Strength:         hero.Strength,
			StrengthStep:     hero.StrengthStep,
			Agility:          hero.Agility,
			AgilityStep:      hero.AgilityStep,
			Intelligence:     hero.Intelligence,
			IntelligenceStep: hero.IntelligenceStep,
			Armor:            hero.Armor,
			AttackMin:        hero.AttackMin,
			AttackMax:        hero.AttackMax,
			Blood:            hero.Blood,
			MP:               hero.MP,
			SkillIds:         skillStr,
		}
		m.db.Where("hero_id = ?", heroDefine.HeroId).FirstOrCreate(&heroDefine)
	}
}

func (m *Module) IntializeSkillTable() {
	skills := InitSkills()
	for _, skill := range skills {
		skillDefine := SkillDefine{
			SkillId: skill.Id,
			Name:    skill.Name,
			Type:    skill.Type,
			Desc:    skill.Desc,
		}
		m.db.Where("skill_id = ?", skillDefine.SkillId).FirstOrCreate(&skillDefine)
	}
}

func (m *Module) IntializeChapterTable() {
	chapters := InitChapters()
	for _, chapter := range chapters {
		chapterDefine := ChapterDefine{
			ChapterId: chapter.Id,
			Name:      chapter.Name,
			GuanKaNum: chapter.GuanKaNum,
		}
		m.db.Where("chapter_id = ?", chapterDefine.ChapterId).FirstOrCreate(&chapterDefine)
	}
}

func (m *Module) IntializeGuanKaTable() {
	gks := InitGuanKas()
	for _, gk := range gks {
		gkDefine := GuanKaDefine{
			GuanKaId:   gk.Id,
			Name:       gk.Name,
			ChapterId:  gk.ChapterId,
			TotalTimes: gk.TotalTimes,
		}
		m.db.Where("guan_ka_id = ?", gkDefine.GuanKaId).FirstOrCreate(&gkDefine)

		if gk.Earn != nil {
			itemIdStr := ""
			if gk.Earn.ItemIds != nil {
				itemIdStr = tool.JoinInt32Arr(gk.Earn.ItemIds, ",")
			}
			gkEarnDefine := GuanKaEarnDefine{
				GuanKaId:  gk.Id,
				HeroExp:   gk.Earn.HeroExp,
				PlayerExp: gk.Earn.PlayerExp,
				Gold:      gk.Earn.Gold,
				ItemIds:   itemIdStr,
			}
			m.db.Where("guan_ka_id = ?", gkEarnDefine.GuanKaId).FirstOrCreate(&gkEarnDefine)
		}

		if gk.Earn != nil {
			gkExpendDefine := GuanKaExpendDefine{
				GuanKaId: gk.Id,
				Power:    gk.Expend.Power,
			}
			m.db.Where("guan_ka_id = ?", gkExpendDefine.GuanKaId).FirstOrCreate(&gkExpendDefine)
		}
	}
}

func (m *Module) IntializeEquipTable() {
	equips := InitEquips()
	for _, equip := range equips {
		equipDefine := EquipDefine{
			ItemId: equip.Id,
			Name:   equip.Name,
			Price:  equip.Price,
			Effect: equip.Equip.Effect,
			Desc:   equip.Desc,
		}
		m.db.Where("item_id = ?", equipDefine.ItemId).FirstOrCreate(&equipDefine)

		if equip.Equip.Mixs != nil {
			for _, mix := range equip.Equip.Mixs {
				mixDefine := EquipMixDefine{
					ItemId:  equip.Id,
					ChildId: mix.ItemId,
					Num:     mix.Num,
				}
				m.db.Where("item_id = ? and child_id = ?", mixDefine.ItemId, mixDefine.ChildId).FirstOrCreate(&mixDefine)
			}
		}
	}
}

func (m *Module) IntializeExpHeroTable() {
	exp := InitExpList()
	for index, e := range exp.Hero {
		expHeroDefine := ExpHeroDefine{
			Level: int32(index) + 1,
			Exp:   e,
		}
		m.db.Where("level = ?", expHeroDefine.Level).FirstOrCreate(&expHeroDefine)
	}
}

func (m *Module) IntializeExpPlayerTable() {
	exp := InitExpList()
	for index, e := range exp.Player {
		expPlayerDefine := ExpPlayerDefine{
			Level: int32(index) + 1,
			Exp:   e,
		}
		m.db.Where("level = ?", expPlayerDefine.Level).FirstOrCreate(&expPlayerDefine)
	}
}

func (m *Module) LoadFromDB() {
	log.Debug("loading data from db start ...")
	m.LoadZone()
	m.LoadPlayer()
	m.LoadHero()
	m.LoadSkill()
	m.LoadChapter()
	m.LoadGuanKa()
	m.LoadEquip()
	m.LoadExp()
	m.LoadUserHero()
	m.LoadUserChapter()
	m.LoadUserGuanKa()
	m.LoadUserEquip()
	m.LoadGameConfig()
	log.Debug("loading data from db end ...")
}

func (m *Module) LoadGameConfig() {
	var configs []GameConfig
	m.db.Find(&configs)
	for _, define := range configs {
		if entry.GameConfigKey_GroupPrice == define.ConfigKey {
			num, err := strconv.Atoi(define.ConfigValue)
			if err == nil {
				m.gameConfig.GroupPrice = int32(num)
			}
		}
	}
	log.Debug("load game configs  db %v ", len(configs))
}

func (m *Module) LoadZone() {
	var zones []*Zone
	m.db.Find(&zones)
	for _, define := range zones {
		z := new(entry.Zone)
		z.Id = define.ZoneId
		z.Name = define.Name
		z.TCPAddr = define.TCPAddr
		z.MaxConnNum = define.MaxConnNum
		z.IsNew = define.IsNew
		m.zones = append(m.zones, z)
	}
	log.Debug("load zones  db %v  mem %v", len(zones), len(m.zones))
}

func (m *Module) LoadPlayer() {
	var users []*User
	m.db.Find(&users)
	tempPlayers := make(map[string]*entry.Player)
	for _, user := range users {
		player := new(entry.Player)
		player.UserId = user.Uid
		player.Account = user.Account
		player.Password = user.Password
		player.LoginTime = user.LoginTime
		player.LogoutTime = user.LogoutTime
		m.SavePlayer(player)

		tempPlayers[user.Uid] = player
	}

	var userBaseInfos []*UserBaseInfo
	m.db.Find(&userBaseInfos)
	for _, baseInfo := range userBaseInfos {
		if tempPlayers[baseInfo.Uid] == nil {
			continue
		}
		info := new(entry.BaseInfo)
		info.Level = baseInfo.Level
		info.Exp = baseInfo.Exp
		info.LevelUpExp = baseInfo.LevelUpExp
		info.Diamond = baseInfo.Diamond
		info.Gold = baseInfo.Gold
		info.Power = baseInfo.Power
		info.MaxPower = baseInfo.MaxPower
		info.Military = baseInfo.Military
		tempPlayers[baseInfo.Uid].Name = baseInfo.Name
		tempPlayers[baseInfo.Uid].BaseInfo = info
	}

	var userTarverns []UserTarvern
	m.db.Find(&userTarverns)
	for _, userTarvern := range userTarverns {
		if tempPlayers[userTarvern.Uid] == nil {
			continue
		}
		extendInfo := tempPlayers[userTarvern.Uid].ExtendInfo
		if extendInfo == nil {
			extendInfo = new(entry.ExtendInfo)
		}
		extendInfo.FreeGoodLottery = userTarvern.FreeGoodLottery
		extendInfo.FreeBetterLottery = userTarvern.FreeBetterLottery
		extendInfo.MaxFreeGoodLottery = userTarvern.MaxFreeGoodLottery
		extendInfo.MaxFreeBetterLottery = userTarvern.MaxFreeBetterLottery
		extendInfo.NeedGoodLotteryCnt = userTarvern.NeedGoodLotteryCnt
		extendInfo.NeedBetterLotteryCnt = userTarvern.NeedBetterLotteryCnt
		extendInfo.LastFreeGoodLotteryStamp = userTarvern.LastFreeGoodLotteryStamp
		extendInfo.LastFreeBetterLotteryStamp = userTarvern.LastFreeBetterLotteryStamp
		extendInfo.GoodLotteryCnt = userTarvern.GoodLotteryCnt
		extendInfo.BetterLotteryCnt = userTarvern.BetterLotteryCnt
		tempPlayers[userTarvern.Uid].ExtendInfo = extendInfo
	}

	tempPlayers = nil

	log.Debug("load players  db %v  mem %v", len(users), len(m.players))
}

func (m *Module) LoadHero() {
	var heroDefines []*HeroDefine
	m.db.Find(&heroDefines)
	for _, define := range heroDefines {
		hero := new(entry.Hero)
		hero.Id = define.HeroId
		hero.Name = define.Name
		hero.Type = define.Type
		hero.Strength = define.Strength
		hero.StrengthStep = define.StrengthStep
		hero.Agility = define.Agility
		hero.AgilityStep = define.AgilityStep
		hero.Intelligence = define.Intelligence
		hero.IntelligenceStep = define.IntelligenceStep
		hero.Armor = define.Armor
		hero.AttackMin = define.AttackMin
		hero.AttackMax = define.AttackMax
		hero.Blood = define.Blood
		hero.MP = define.MP
		hero.MaxBlood = define.Blood
		hero.MaxMP = define.MP
		if len(define.SkillIds) > 0 {
			hero.SkillIds = tool.SpliteInt32Arr(define.SkillIds, ",")
		}
		m.heros = append(m.heros, hero)
	}

	log.Debug("load heros  db %v  mem %v", len(heroDefines), len(m.heros))
}

func (m *Module) LoadSkill() {
	var skillDefines []*SkillDefine
	m.db.Find(&skillDefines)
	for _, define := range skillDefines {
		skill := new(entry.Skill)
		skill.Id = define.SkillId
		skill.Name = define.Name
		skill.Type = define.Type
		skill.Desc = define.Desc
		m.skills = append(m.skills, skill)
	}

	log.Debug("load skills  db %v  mem %v", len(skillDefines), len(m.skills))
}

func (m *Module) LoadChapter() {
	var chapterDefines []*ChapterDefine
	m.db.Find(&chapterDefines)
	for _, define := range chapterDefines {
		chapter := new(entry.Chapter)
		chapter.Id = define.ChapterId
		chapter.Name = define.Name
		chapter.GuanKaNum = define.GuanKaNum
		m.chapters = append(m.chapters, chapter)
	}

	log.Debug("load chapters  db %v  mem %v", len(chapterDefines), len(m.chapters))
}

func (m *Module) LoadGuanKa() {
	var gkDefines []*GuanKaDefine
	m.db.Find(&gkDefines)
	for _, define := range gkDefines {
		gk := new(entry.GuanKa)
		gk.Id = define.GuanKaId
		gk.Name = define.Name
		gk.TotalTimes = define.TotalTimes
		gk.ChapterId = define.ChapterId

		var gkEarnDefine GuanKaEarnDefine
		m.db.Where("guan_ka_id = ?", define.GuanKaId).First(&gkEarnDefine)
		gk.Earn = new(entry.Earn)
		gk.Earn.ItemIds = tool.SpliteInt32Arr(gkEarnDefine.ItemIds, ",")
		gk.Earn.HeroExp = gkEarnDefine.HeroExp
		gk.Earn.PlayerExp = gkEarnDefine.PlayerExp
		gk.Earn.Gold = gkEarnDefine.Gold

		var gkExpendDefine GuanKaExpendDefine
		m.db.Where("guan_ka_id = ?", define.GuanKaId).First(&gkExpendDefine)
		gk.Expend = new(entry.Expend)
		gk.Expend.Power = gkExpendDefine.Power

		m.guanKas = append(m.guanKas, gk)
	}

	log.Debug("load guankas  db %v  mem %v", len(gkDefines), len(m.guanKas))
}

func (m *Module) LoadEquip() {
	var equipDefines []*EquipDefine
	m.db.Find(&equipDefines)
	for _, define := range equipDefines {
		item := new(entry.Item)
		item.Id = define.ItemId
		item.Type = entry.ItemTypeEquip
		item.Name = define.Name
		item.Price = define.Price
		item.Desc = define.Desc

		item.Equip = new(entry.Equip)
		item.Equip.Effect = define.Effect
		item.Equip.Mixs = make([]*entry.Mix, 0)

		var mixDefines []*EquipMixDefine
		m.db.Where("item_id = ?", define.ItemId).Find(&mixDefines)
		for _, m := range mixDefines {
			mix := new(entry.Mix)
			mix.ItemId = m.ChildId
			mix.Num = m.Num
			item.Equip.Mixs = append(item.Equip.Mixs, mix)
		}

		m.items = append(m.items, item)
	}

	cnt := 0
	for _, item := range m.items {
		if entry.ItemTypeEquip == item.Type {
			cnt++
		}
	}

	log.Debug("load equips  db %v  mem %v", len(equipDefines), cnt)
}

func (m *Module) LoadExp() {
	var expHeroDefines []ExpHeroDefine
	m.db.Find(&expHeroDefines)
	for _, define := range expHeroDefines {
		m.heroExpList = append(m.heroExpList, define.Exp)
	}
	var expPlayerDefines []ExpPlayerDefine
	m.db.Find(&expPlayerDefines)
	for _, define := range expPlayerDefines {
		m.playerExpList = append(m.playerExpList, define.Exp)
	}

	log.Debug("load exps hero  db %v  mem %v", len(expHeroDefines), len(m.heroExpList))
	log.Debug("load exps player  db %v  mem %v", len(expPlayerDefines), len(m.playerExpList))
}

func (m *Module) LoadUserHero() {
	var userHeros []UserHero
	m.db.Find(&userHeros)
	for _, define := range userHeros {
		heros := m.playerHeros[define.Uid]
		if heros == nil {
			heros = make([]*entry.Hero, 0)
		}
		hero := new(entry.Hero)
		hero.PlayerId = define.Uid
		hero.HeroId = define.HeroId
		hero.Level = define.Level
		hero.Exp = define.Exp
		hero.LevelUpExp = define.LevelUpExp
		hero.Strength = define.Strength
		hero.StrengthStep = define.StrengthStep
		hero.Agility = define.Agility
		hero.AgilityStep = define.AgilityStep
		hero.Intelligence = define.Intelligence
		hero.IntelligenceStep = define.IntelligenceStep
		hero.AttackMin = define.AttackMin
		hero.AttackMax = define.AttackMax
		hero.Armor = define.Armor
		hero.MaxBlood = define.Blood
		hero.MaxMP = define.MP
		hero.IsSelect = define.IsSelect
		hero.Pos = define.Pos

		var heroDefine HeroDefine
		m.db.Where("hero_id = ?", define.HeroDefineId).First(&heroDefine)
		hero.Name = heroDefine.Name
		hero.Type = heroDefine.Type

		hero.Skills = make([]*entry.Skill, 0)
		var heroSkills []UserHeroSkill
		m.db.Where("hero_id = ?", define.HeroId).First(&heroSkills)
		for _, heroSkill := range heroSkills {
			skill := new(entry.Skill)
			skill.HeroId = heroSkill.HeroId
			skill.SkillId = heroSkill.SkillId
			skill.Id = heroSkill.SkillDefineId
			skill.Level = heroSkill.Level
			skill.IsOpen = heroSkill.IsOpen

			var skillDefine SkillDefine
			m.db.Where("skill_id = ?", heroSkill.SkillDefineId).First(&skillDefine)
			skill.Name = skillDefine.Name
			skill.Type = skillDefine.Type
			skill.Desc = skillDefine.Desc

			hero.Skills = append(hero.Skills, skill)
		}

		heros = append(heros, hero)
		m.playerHeros[define.Uid] = heros
	}
	cnt := 0
	for _, heros := range m.playerHeros {
		cnt += len(heros)
	}
	log.Debug("load user hero  db %v  mem %v", len(userHeros), cnt)
}

func (m *Module) LoadUserChapter() {
	var userChapters []UserChapter
	m.db.Find(&userChapters)
	for _, define := range userChapters {
		chapter := new(entry.Chapter)
		chapter.Id = define.ChapterDefineId
		chapter.ChapterId = define.ChapterId
		chapter.Star = define.Star
		chapter.Status = define.Status
		chapter.IsOpen = define.IsOpen

		var chapterDefine ChapterDefine
		m.db.Where("chapter_id = ?", define.ChapterDefineId).First(&chapterDefine)
		chapter.Name = chapterDefine.Name
		chapter.GuanKaNum = chapterDefine.GuanKaNum

		chapters := m.playerChapters[define.Uid]
		if chapters == nil {
			chapters = make([]*entry.Chapter, 0)
		}

		chapters = append(chapters, chapter)

		m.playerChapters[define.Uid] = chapters
	}
	cnt := 0
	for _, chapters := range m.playerChapters {
		cnt += len(chapters)
	}

	log.Debug("load user chapter  db %v  mem %v", len(userChapters), cnt)
}

func (m *Module) LoadUserGuanKa() {
	var userGuanKas []UserGuanKa
	m.db.Find(&userGuanKas)
	for _, define := range userGuanKas {
		gk := new(entry.GuanKa)
		gk.Id = define.GuanKaDefineId
		gk.GuanKaId = define.GuanKaId
		gk.Star = define.Star
		gk.Status = define.Status
		gk.IsOpen = define.IsOpen
		gk.Times = define.Times

		var gkDefine GuanKaDefine
		m.db.Where("guan_ka_id = ?", define.GuanKaDefineId).First(&gkDefine)
		gk.Name = gkDefine.Name
		gk.ChapterId = gkDefine.ChapterId
		gk.TotalTimes = gkDefine.TotalTimes

		var gkEarnDefine GuanKaEarnDefine
		m.db.Where("guan_ka_id = ?", define.GuanKaDefineId).First(&gkEarnDefine)
		gk.Earn = new(entry.Earn)
		gk.Earn.ItemIds = tool.SpliteInt32Arr(gkEarnDefine.ItemIds, ",")
		gk.Earn.HeroExp = gkEarnDefine.HeroExp
		gk.Earn.PlayerExp = gkEarnDefine.PlayerExp
		gk.Earn.Gold = gkEarnDefine.Gold

		var gkExpendDefine GuanKaExpendDefine
		m.db.Where("guan_ka_id = ?", define.GuanKaDefineId).First(&gkExpendDefine)
		gk.Expend = new(entry.Expend)
		gk.Expend.Power = gkExpendDefine.Power

		gks := m.playerGuanKas[define.Uid]
		if gks == nil {
			gks = make([]*entry.GuanKa, 0)
		}

		gks = append(gks, gk)

		m.playerGuanKas[define.Uid] = gks
	}
	cnt := 0
	for _, gks := range m.playerGuanKas {
		cnt += len(gks)
	}

	log.Debug("load user guanka  db %v  mem %v", len(userGuanKas), cnt)
}

func (m *Module) LoadUserEquip() {
	var userEquips []UserEquip
	m.db.Find(&userEquips)
	for _, define := range userEquips {

		items := m.playerItems[define.Uid]
		if items == nil {
			items = make([]*entry.Item, 0)
		}

		item := new(entry.Item)
		item.Type = entry.ItemTypeEquip
		item.Id = define.ItemDefineId
		item.ItemId = define.EquipId
		item.Equip = new(entry.Equip)

		var equipDefine EquipDefine
		m.db.Where("item_id = ?", define.ItemDefineId).First(&equipDefine)
		item.Name = equipDefine.Name
		item.Price = equipDefine.Price
		item.Desc = equipDefine.Desc
		item.Equip.Effect = equipDefine.Effect
		item.Equip.HeroId = define.HeroId

		item.Equip.Mixs = make([]*entry.Mix, 0)

		var mixDefines []*EquipMixDefine
		m.db.Where("item_id = ?", equipDefine.ItemId).Find(&mixDefines)
		for _, m := range mixDefines {
			mix := new(entry.Mix)
			mix.ItemId = m.ChildId
			mix.Num = m.Num
			item.Equip.Mixs = append(item.Equip.Mixs, mix)
		}

		items = append(items, item)

		m.playerItems[define.Uid] = items
	}
	cnt := 0
	for _, items := range m.playerItems {
		for _, item := range items {
			if entry.ItemTypeEquip == item.Type {
				cnt++
			}
		}
	}

	log.Debug("load user equip  db %v  mem %v", len(userEquips), cnt)
}

func (m *Module) DoPersistent() {
	log.Debug("DoPersistent =====>  start")
	m.PersistentUser()
	m.PersistentUserHero()
	m.PersistentChapter()
	m.PersistentGuanKa()
	m.PersistentGroup()
	m.PersistentGroupMember()
	m.PersistentItems()
	log.Debug("DoPersistent =====>  end")
}

func (m *Module) PersistentItems() {
	cnt := 0

	for uid, items := range m.playerItems {
		for _, item := range items {
			if item.IsDirty {
				if entry.ItemTypeEquip == item.Type {
					userEquip := UserEquip{
						Uid:          uid,
						EquipId:      item.ItemId,
						ItemDefineId: item.Id,
						HeroId:       item.Equip.HeroId,
					}

					var oldUserEquip UserEquip
					m.db.Where("equip_id = ?", userEquip.EquipId).First(&oldUserEquip)
					if userEquip.EquipId != oldUserEquip.EquipId {
						m.db.Create(&userEquip)
					} else {
						m.db.Model(&userEquip).Where("equip_id = ?", userEquip.EquipId).Updates(userEquip)
					}

					cnt++
					item.IsDirty = false
				}
			}
		}
	}

	log.Debug("persistent user equip %v ", cnt)
}

func (m *Module) PersistentGroupMember() {
	cnt := 0
	for _, group := range m.groups {
		if group.GroupMembers != nil {
			for _, groupMember := range group.GroupMembers {
				if groupMember.IsDirty {
					gpm := GroupMember{
						GroupId:     group.GroupId,
						UserId:      groupMember.UserId,
						ContriToday: groupMember.ContriToday,
						ContriTotal: groupMember.ContriTotal,
						Job:         groupMember.Job,
					}

					var oldGroupMember GroupMember
					m.db.Where("user_id = ?", gpm.UserId).First(&oldGroupMember)
					if gpm.UserId != oldGroupMember.UserId {
						m.db.Create(&gpm)
					} else {
						m.db.Model(&gpm).Where("user_id = ?", gpm.GroupId).Updates(gpm)
					}

					cnt++
					groupMember.IsDirty = false
				}
			}
		}
	}
	log.Debug("persistent groupMember %v ", cnt)
}

func (m *Module) PersistentGroup() {
	cnt := 0
	for _, group := range m.groups {
		if group.IsDirty {

			gp := Group{
				GroupId:          group.GroupId,
				GroupName:        group.GroupName,
				GroupLeader:      group.GroupName,
				GroupDeclaration: group.GroupDeclaration,
				MemberCnt:        group.MemberCnt,
				MemberTotal:      group.MemberTotal,
				GroupLevel:       group.GroupLevel,
				ContriCurrent:    group.ContriCurrent,
				ContriLevelUp:    group.ContriLevelUp,
			}

			var oldGroup Group
			m.db.Where("group_id = ?", gp.GroupId).First(&oldGroup)
			if gp.GroupId != oldGroup.GroupId {
				m.db.Create(&gp)
			} else {
				m.db.Model(&gp).Where("group_id = ?", gp.GroupId).Updates(gp)
			}

			cnt++
			group.IsDirty = false
		}
	}
	log.Debug("persistent group %v ", cnt)
}

func (m *Module) PersistentGuanKa() {
	cnt := 0
	for uid, guanKas := range m.playerGuanKas {
		for _, guanKa := range guanKas {
			if guanKa.IsDirty {

				userGuanKa := UserGuanKa{
					Uid:            uid,
					GuanKaId:       guanKa.GuanKaId,
					GuanKaDefineId: guanKa.Id,
					Star:           guanKa.Star,
					Status:         guanKa.Status,
					IsOpen:         guanKa.IsOpen,
					Times:          guanKa.Times,
				}

				var oldUserGuanKa UserGuanKa
				m.db.Where("guan_ka_id = ?", userGuanKa.GuanKaId).First(&oldUserGuanKa)
				if userGuanKa.GuanKaId != oldUserGuanKa.GuanKaId {
					m.db.Create(&userGuanKa)
				} else {
					m.db.Model(&userGuanKa).Where("guan_ka_id = ?", userGuanKa.GuanKaId).Updates(userGuanKa)
				}

				cnt++
				guanKa.IsDirty = false
			}
		}
	}
	log.Debug("persistent user guanka %v ", cnt)
}

func (m *Module) PersistentChapter() {
	cnt := 0
	for uid, chapters := range m.playerChapters {
		for _, chapter := range chapters {
			if chapter.IsDirty {

				userChapter := UserChapter{
					Uid:             uid,
					ChapterId:       chapter.ChapterId,
					ChapterDefineId: chapter.Id,
					Star:            chapter.Star,
					Status:          chapter.Status,
					IsOpen:          chapter.IsOpen,
				}

				var oldUserChapter UserChapter
				m.db.Where("chapter_id = ?", userChapter.ChapterId).First(&oldUserChapter)
				if userChapter.ChapterId != oldUserChapter.ChapterId {
					m.db.Create(&userChapter)
				} else {
					m.db.Model(&userChapter).Where("chapter_id = ?", userChapter.ChapterId).Updates(userChapter)
				}

				cnt++
				chapter.IsDirty = false
			}
		}
	}
	log.Debug("persistent user chapter %v ", cnt)
}

func (m *Module) PersistentUserHero() {
	cnt := 0
	for uid, heros := range m.playerHeros {
		for _, hero := range heros {
			if hero.IsDirty {

				userHero := UserHero{
					Uid:              uid,
					HeroId:           hero.HeroId,
					HeroDefineId:     hero.Id,
					Level:            hero.Level,
					Exp:              hero.Exp,
					LevelUpExp:       hero.LevelUpExp,
					Strength:         hero.Strength,
					StrengthStep:     hero.StrengthStep,
					Agility:          hero.Agility,
					AgilityStep:      hero.AgilityStep,
					Intelligence:     hero.Intelligence,
					IntelligenceStep: hero.IntelligenceStep,
					Armor:            hero.Armor,
					AttackMin:        hero.AttackMin,
					AttackMax:        hero.AttackMax,
					Blood:            hero.MaxBlood,
					MP:               hero.MaxMP,
					IsSelect:         hero.IsSelect,
					Pos:              hero.Pos,
				}

				var oldUserHero UserHero
				m.db.Where("hero_id = ?", userHero.HeroId).First(&oldUserHero)
				if userHero.HeroId != oldUserHero.HeroId {
					m.db.Create(&userHero)
				} else {
					m.db.Model(&userHero).Where("hero_id = ?", userHero.HeroId).Updates(userHero)
				}

				for _, skill := range hero.Skills {
					if skill.IsDirty {

						heroSkill := UserHeroSkill{
							HeroId:        skill.HeroId,
							SkillId:       skill.SkillId,
							SkillDefineId: skill.Id,
							Level:         skill.Level,
							IsOpen:        skill.IsOpen,
						}

						var oldHeroSkill UserHeroSkill
						m.db.Where("skill_id = ?", heroSkill.SkillId).First(&oldHeroSkill)
						if heroSkill.SkillId != oldHeroSkill.SkillId {
							m.db.Create(&heroSkill)
						} else {
							m.db.Model(&heroSkill).Where("skill_id = ?", heroSkill.SkillId).Updates(heroSkill)
						}

						skill.IsDirty = false
					}
				}

				cnt++
				hero.IsDirty = false
			}
		}
	}
	log.Debug("persistent user hero %v ", cnt)
}

func (m *Module) PersistentUser() {
	cnt := 0
	for _, player := range m.players {
		if player.IsDirty {
			user := User{
				Uid:        player.UserId,
				Account:    player.Account,
				Password:   player.Password,
				LoginTime:  player.LoginTime,
				LogoutTime: player.LogoutTime,
			}

			var oldUser User
			m.db.Where("uid = ?", user.Uid).First(&oldUser)
			if user.Uid != oldUser.Uid {
				m.db.Create(&user)
			} else {
				m.db.Model(&user).Where("uid = ?", user.Uid).Updates(user)
			}

			player.IsDirty = false

			cnt++
		}

		if player.BaseInfo != nil && player.BaseInfo.IsDirty {
			userBaseInfo := UserBaseInfo{
				Uid:        player.UserId,
				Name:       player.Name,
				Level:      player.BaseInfo.Level,
				LevelUpExp: player.BaseInfo.LevelUpExp,
				Exp:        player.BaseInfo.Exp,
				Gold:       player.BaseInfo.Gold,
				Diamond:    player.BaseInfo.Diamond,
				Power:      player.BaseInfo.Power,
				MaxPower:   player.BaseInfo.MaxPower,
				Military:   player.BaseInfo.Military,
			}

			log.Debug("userBaseInfo === %v", userBaseInfo)

			var oldUserInfo UserBaseInfo
			m.db.Where("uid = ?", userBaseInfo.Uid).First(&oldUserInfo)
			if userBaseInfo.Uid != oldUserInfo.Uid {
				m.db.Create(&userBaseInfo)
			} else {
				m.db.Model(&userBaseInfo).Where("uid = ?", userBaseInfo.Uid).Updates(userBaseInfo)
			}

			player.BaseInfo.IsDirty = false

			cnt++
		}

		if player.ExtendInfo != nil && player.ExtendInfo.IsDirty {
			userTarvern := UserTarvern{
				Uid:                        player.UserId,
				FreeGoodLottery:            player.ExtendInfo.FreeGoodLottery,
				FreeBetterLottery:          player.ExtendInfo.FreeBetterLottery,
				MaxFreeGoodLottery:         player.ExtendInfo.MaxFreeGoodLottery,
				MaxFreeBetterLottery:       player.ExtendInfo.MaxFreeBetterLottery,
				LastFreeGoodLotteryStamp:   player.ExtendInfo.LastFreeGoodLotteryStamp,
				LastFreeBetterLotteryStamp: player.ExtendInfo.LastFreeBetterLotteryStamp,
				GoodLotteryCnt:             player.ExtendInfo.GoodLotteryCnt,
				BetterLotteryCnt:           player.ExtendInfo.BetterLotteryCnt,
				NeedGoodLotteryCnt:         player.ExtendInfo.NeedGoodLotteryCnt,
				NeedBetterLotteryCnt:       player.ExtendInfo.NeedBetterLotteryCnt,
			}

			var oldUserTarvern UserTarvern
			m.db.Where("uid = ?", userTarvern.Uid).First(&oldUserTarvern)
			if userTarvern.Uid != oldUserTarvern.Uid {
				m.db.Create(&userTarvern)
			} else {
				m.db.Model(&userTarvern).Where("uid = ?", userTarvern.Uid).Updates(userTarvern)
			}

			player.ExtendInfo.IsDirty = false

			cnt++
		}
	}
	log.Debug("persistent user %v ", cnt)
}

// func openDB() (success bool, db *sql.DB) {
// 	var isOpen bool
// 	db, err := sql.Open("mysql", DB_Driver)
// 	if err != nil {
// 		isOpen = false
// 	} else {
// 		isOpen = true
// 	}
// 	CheckErr(err)
// 	return isOpen, db
// }

// func insertToDB(db *sql.DB) {
// 	uid := GetNowtimeMD5()
// 	nowTimeStr := GetTime()
// 	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?,password=?,uid=?")
// 	CheckErr(err)
// 	res, err := stmt.Exec("wangbiao", "研发中心", nowTimeStr, "123456", uid)
// 	CheckErr(err)
// 	id, err := res.LastInsertId()
// 	CheckErr(err)
// 	if err != nil {
// 		fmt.Println("插入数据失败")
// 	} else {
// 		fmt.Println("插入数据成功：", id)
// 	}
// }

// func QueryFromDB(db *sql.DB) {
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	CheckErr(err)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 	}
// 	for rows.Next() {
// 		var uid string
// 		var username string
// 		var departmentname string
// 		var created string
// 		var password string
// 		var autid string
// 		CheckErr(err)
// 		err = rows.Scan(&uid, &username, &departmentname, &created, &password, &autid)
// 		fmt.Println(autid)
// 		fmt.Println(username)
// 		fmt.Println(departmentname)
// 		fmt.Println(created)
// 		fmt.Println(password)
// 		fmt.Println(uid)
// 	}
// }

// func UpdateDB(db *sql.DB, uid string) {
// 	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
// 	CheckErr(err)
// 	res, err := stmt.Exec("zhangqi", uid)
// 	affect, err := res.RowsAffected()
// 	fmt.Println("更新数据：", affect)
// 	CheckErr(err)
// }

// func DeleteFromDB(db *sql.DB, autid int) {
// 	stmt, err := db.Prepare("delete from userinfo where autid=?")
// 	CheckErr(err)
// 	res, err := stmt.Exec(autid)
// 	affect, err := res.RowsAffected()
// 	fmt.Println("删除数据：", affect)
// }
