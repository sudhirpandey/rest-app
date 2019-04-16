package model

import (
	//"encoding/json"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	
)

type Product struct {
	gorm.Model
	Name         string `gorm:"unique" json:"name"`
	Description  string `gorm:type:varchar(100) json:"description"`
	Price        float64 `gorm:"not null" json:"price,string"`  //canUse json.Number as type too.
}


func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{})
	//db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
