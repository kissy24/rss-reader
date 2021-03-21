package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type XML struct {
	gorm.Model
	name    string
	content string
}

// Open DB
func Open() *gorm.DB {
	db, err := gorm.Open("sqlite3", "xml.sqlite3")
	if err != nil {
		panic("Don't open sqlite3")
	}
	return db
}

// Init DB
func Init() {
	db := Open()
	db.AutoMigrate(&XML{})
	defer db.Close()
}

// Insert DB
func Insert(name string, content string) {
	db := Open()
	db.Create(&XML{name: name, content: content})
	defer db.Close()
}

// Select DB
func Select(id int) XML {
	db := Open()
	var xml XML
	db.First(&xml, id)
	db.Close()
	return xml
}

// Update DB
func Update(id int, name string, content string) {
	db := Open()
	var xml XML
	xml.name = name
	xml.content = content
	db.Save(&xml)
	db.Close()
}

// Delete DB
func Delete(id int) {
	db := Open()
	var xml XML
	db.First(&xml, id)
	db.Delete(&xml)
	db.Close()
}
