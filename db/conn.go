package db

import (
	"github.com/fatih/color"
	donut_rpc "github.com/s1ntaxe770r/donut-rpc/proto"
	"github.com/s1ntaxe770r/donut-rpc/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dblg = utils.NewDBLogger()
)

// Connect instantiates a new database client
func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		dblg.Panicln("UNABLE TO CONNECT TO DATABASE")
	}
	db.AutoMigrate(&donut_rpc.Donut{})
	dblg.Println(color.GreenString("DB CONNECTION ESTABLISHED"))
	return db
}

func MakeDonut(db *gorm.DB, donut *donut_rpc.Donut) (*donut_rpc.Donut, error) {
	err := db.Create(&donut).Error
	if err != nil {
		dblg.Println(err.Error())
		return donut, err
	}
	return donut, nil
}

func GetDonuts(db *gorm.DB) (*donut_rpc.Donuts, error) {
	var donuts donut_rpc.Donuts
	err := db.Find(&donuts.Donuts).Error
	if err != nil {
		return nil, err
	}
	return &donuts, nil
}

func GetDonut(db *gorm.DB, donut *donut_rpc.DonutRequest) (*donut_rpc.Donut, error) {
	var res donut_rpc.Donut
	err := db.First(&res, donut.Name).Error
	if err != nil {
		dblg.Fatal(err.Error())
		return &res, err
	}
	return &res, nil
}
