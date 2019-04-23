package internal

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/name5566/leaf/log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_Driver = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/pixel_farm?charset=utf8"
)

func (m *Module) PersistentData() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "dota_" + defaultTableName
	}

	m.ConnectDB()
	m.CreateTables()
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

			user := User{Uid: player.UserId, Account: player.Account, Password: player.Password}
			if m.db.NewRecord(user) {
				m.db.Create(&user)
			} else {
				m.db.Model(&user).Updates(user)
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
