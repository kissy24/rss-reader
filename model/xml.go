package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type XML struct {
	gorm.Model
	Name    string
	Content string
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
	db.Create(&XML{Name: name, Content: content})
	defer db.Close()
}

// Select All DB
func SelectAll() []XML {
	db := Open()
	var xmls []XML
	db.Order("created_at desc").Find(&xmls)
	db.Close()
	return xmls
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
	xml.ID = uint(id)
	xml.Name = name
	xml.Content = content
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
