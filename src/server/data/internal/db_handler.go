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
	m.LoadFromDB()
	m.IntializeTables()

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
}

func (m *Module) IntializeTables() {
	m.IntializeHeroTable()
	m.IntializeSkillTable()
	m.IntializeChapterTable()
	m.IntializeGuanKaTable()
	m.IntializeItemTable()
}

func (m *Module) IntializeHeroTable() {
	heros := InitHeros()
	for _, hero := range heros {

		if m.HasHero(hero.Id) {
			continue
		}

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
		m.db.Create(&heroDefine)
	}
}

func (m *Module) IntializeSkillTable() {
	skills := InitSkills()
	for _, skill := range skills {

		if m.HasSkill(skill.Id) {
			continue
		}

		skillDefine := SkillDefine{
			SkillId: skill.Id,
			Name:    skill.Name,
			Type:    skill.Type,
			Desc:    skill.Desc,
		}
		m.db.Create(&skillDefine)
	}
}

func (m *Module) IntializeChapterTable() {
	chapters := InitChapters()
	for _, chapter := range chapters {

		if m.HasChapter(chapter.Id) {
			continue
		}

		chapterDefine := ChapterDefine{
			ChapterId: chapter.Id,
			Name:      chapter.Name,
			GuanKaNum: chapter.GuanKaNum,
		}
		m.db.Create(&chapterDefine)
	}
}

func (m *Module) IntializeGuanKaTable() {
	gks := InitGuanKas()
	for _, gk := range gks {

		if m.HasGuanKa(gk.Id) {
			continue
		}

		gkDefine := GuanKaDefine{
			GuanKaId:   gk.Id,
			Name:       gk.Name,
			ChapterId:  gk.ChapterId,
			TotalTimes: gk.TotalTimes,
		}
		m.db.Create(&gkDefine)
	}
}

func (m *Module) IntializeItemTable() {
	items := InitItems()
	for _, item := range items {

		if m.HasItem(item.Id) {
			continue
		}

		itemDefine := ItemDefine{
			ItemId: item.Id,
			Name:   item.Name,
			Price:  item.Price,
			Effect: item.Effect,
			Desc:   item.Desc,
		}
		m.db.Create(&itemDefine)

		if item.Mixs != nil {
			for _, mix := range item.Mixs {
				mixDefine := ItemMixDefine{
					ItemId:  item.Id,
					ChildId: mix.ItemId,
					Num:     mix.Num,
				}
				m.db.Create(&mixDefine)
			}
		}
	}
}

func (m *Module) LoadFromDB() {
	m.LoadPlayer()
	m.LoadHero()
	m.LoadSkill()
	m.LoadChapter()
	m.LoadGuanKa()
	m.LoadItem()
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
		tempPlayers[baseInfo.Uid].BaseInfo = info
	}

	tempPlayers = nil
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
		m.guanKas = append(m.guanKas, gk)
	}
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
	for _, player := range m.players {
		if player.IsDirty {
			log.Debug("save %v", player)

			user := User{
				Uid:      player.UserId,
				Account:  player.Account,
				Password: player.Password}
			if m.db.NewRecord(user) {
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
				MaxPower:   player.BaseInfo.MaxPower}
			if m.db.NewRecord(userBaseInfo) {
				m.db.Create(&userBaseInfo)
			} else {
				m.db.Model(&userBaseInfo).Updates(userBaseInfo)
			}

			player.IsDirty = false
		}
	}
	log.Debug("DoPersistent =====>  end")
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
