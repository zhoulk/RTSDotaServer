package internal

import (
	"server/data/entry"
	"server/tool"
	"time"

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

	go m.ExecutePersistent()
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
	if !m.db.HasTable(&ItemDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ItemDefine{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ItemMixDefine{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ItemMixDefine{}).Error; err != nil {
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
}

func (m *Module) IntializeTables() {
	m.IntializeZoneTable()
	m.IntializeHeroTable()
	m.IntializeSkillTable()
	m.IntializeChapterTable()
	m.IntializeGuanKaTable()
	m.IntializeItemTable()
	m.IntializeExpPlayerTable()
	m.IntializeExpHeroTable()
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

func (m *Module) IntializeItemTable() {
	items := InitItems()
	for _, item := range items {
		itemDefine := ItemDefine{
			ItemId: item.Id,
			Name:   item.Name,
			Price:  item.Price,
			Effect: item.Effect,
			Desc:   item.Desc,
		}
		m.db.Where("item_id = ?", itemDefine.ItemId).FirstOrCreate(&itemDefine)

		if item.Mixs != nil {
			for _, mix := range item.Mixs {
				mixDefine := ItemMixDefine{
					ItemId:  item.Id,
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
	m.LoadItem()
	m.LoadExp()
	m.LoadUserHero()
	log.Debug("loading data from db end ...")
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
		m.SavePlayer(player)

		tempPlayers[user.Uid] = player
	}

	var userBaseInfos []*UserBaseInfo
	m.db.Find(&userBaseInfos)
	for _, baseInfo := range userBaseInfos {
		info := new(entry.BaseInfo)
		info.Level = baseInfo.Level
		info.Exp = baseInfo.Exp
		info.LevelUpExp = baseInfo.LevelUpExp
		info.Diamond = baseInfo.Diamond
		info.Gold = baseInfo.Gold
		info.Power = baseInfo.Power
		info.MaxPower = baseInfo.MaxPower
		tempPlayers[baseInfo.Uid].Name = baseInfo.Name
		tempPlayers[baseInfo.Uid].BaseInfo = info
	}

	var userTarverns []UserTarvern
	m.db.Find(&userTarverns)
	for _, userTarvern := range userTarverns {
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

func (m *Module) LoadItem() {
	var itemDefines []*ItemDefine
	m.db.Find(&itemDefines)
	for _, define := range itemDefines {
		item := new(entry.Item)
		item.Id = define.ItemId
		item.Name = define.Name
		item.Price = define.Price
		item.Effect = define.Effect
		item.Desc = define.Desc

		var mixDefines []*ItemMixDefine
		m.db.Where("item_id = ?", define.ItemId).Find(&mixDefines)
		item.Mixs = make([]*entry.Mix, 0)
		for _, m := range mixDefines {
			mix := new(entry.Mix)
			mix.ItemId = m.ChildId
			mix.Num = m.Num
			item.Mixs = append(item.Mixs, mix)
		}

		m.items = append(m.items, item)
	}

	log.Debug("load items  db %v  mem %v", len(itemDefines), len(m.items))
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
		heros = append(heros, hero)
		m.playerHeros[define.Uid] = heros
	}
	cnt := 0
	for _, heros := range m.playerHeros {
		cnt += len(heros)
	}
	log.Debug("load user hero  db %v  mem %v", len(userHeros), cnt)
}

func (m *Module) ExecutePersistent() {
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-timer.C:
			m.DoPersistent()
		}
	}
}

func (m *Module) DoPersistent() {
	log.Debug("DoPersistent =====>  start")
	m.PersistentUser()
	m.PersistentUserHero()
	log.Debug("DoPersistent =====>  end")
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
				}

				var oldUserHero UserHero
				m.db.Where("hero_id = ?", userHero.HeroId).First(&oldUserHero)
				if userHero.HeroId != oldUserHero.HeroId {
					m.db.Create(&userHero)
				} else {
					m.db.Model(&userHero).Updates(userHero)
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
				Uid:      player.UserId,
				Account:  player.Account,
				Password: player.Password,
			}

			var oldUser User
			m.db.Where("uid = ?", user.Uid).First(&oldUser)
			if user.Uid != oldUser.Uid {
				m.db.Create(&user)
			} else {
				m.db.Model(&user).Updates(user)
			}

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
			}

			var oldUserInfo UserBaseInfo
			m.db.Where("uid = ?", userBaseInfo.Uid).First(&oldUserInfo)
			if userBaseInfo.Uid != oldUserInfo.Uid {
				m.db.Create(&userBaseInfo)
			} else {
				m.db.Model(&userBaseInfo).Updates(userBaseInfo)
			}

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
				m.db.Model(&userTarvern).Updates(userTarvern)
			}

			cnt++

			player.IsDirty = false
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
