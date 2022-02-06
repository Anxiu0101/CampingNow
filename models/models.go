package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"

	"CampingNow/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Read the config file and Open the database
func init() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)

	// Read the config from app.ini file
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// Read value from config attribute name
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// pass config to dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	// open the database and buffer the config
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // set the prefix name of table
			SingularTable: true,        // use singular table by default
		},
		Logger: logger.Default.LogMode(logger.Info), // set log mode
	})

	mysqlDB, err := db.DB()
	if err != nil {
		log.Panicln("db.DB() err: ", err)
	}

	// some init set of database
	mysqlDB.SetMaxIdleConns(10)  // set max idle connections
	mysqlDB.SetMaxOpenConns(100) // set open connections, default is 0 (unlimited)
}

// CloseDB Close database
func CloseDB() {
	mysqlDB, _ := db.DB()
	defer mysqlDB.Close()
}
