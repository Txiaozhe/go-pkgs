package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spaolacci/murmur3"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

// define a table
type Like struct {
	ID       int    `gorm:"primary_key"`
	Ip       string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua       string `gorm:"type:varchar(256);not null;"`
	Title    string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash     uint64 `gorm:"unique_index:hash_idx;"`
	CreateAt time.Time
}

func main() {
	// create table
	if !db.HasTable(&Like{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}

	var (
		ip    = "127.0.0.1"
		ua    = "dddd"
		title = "hhh"
	)

	// insert
	like := &Like{
		Ip:    ip,
		Ua:    ua,
		Title: title,
		Hash:  murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
	}

	if err := db.Create(like).Error; err != nil {
		panic(err)
	}

	// delete
	var hash uint64 = 123455
	if err := db.Where(&Like{Hash: hash}).Delete(Like{}).Error; err != nil {
		panic(err)
	}

	// query
	var count int
	if err := db.Model(&Like{}).Where(&Like{Ip: ip, Ua: ua, Title: title}).Count(&count).Error; err != nil {
		panic(err)
	}

	// update
	db.Model(&Like{}).Update("Title", "newtitle")
	db.Model(&Like{}).Updates(Like{Title: "newtitle"})

	// set logger
	db.LogMode(true)
	// db.SetLogger(gorm.Logger{revel.TRACE})
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

func CreateAnimals(db *gorm.DB) error {
	tx := db.Begin()
	if err := tx.Create(&Like{Title: "hhh"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Like{Title: "kkk"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
