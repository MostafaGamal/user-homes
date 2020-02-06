package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID			int			`gorm:"primary_key;AUTO_INCREMENT"`

	Name		string		`gorm:"not_null"`
	Password	string		`gorm:"not_null"`
	Email		string		`gorm:"not_null;unique_index"`
	Address		string		`gorm:"not_null"`
	Age			int			`gorm:"not_null"`
	Homes		[]Home
}

type Home struct {
	ID			string 		`gorm:"primary_key"`

	Serial		string		`gorm:"not_null"`
	Sensors		string		`gorm:"type:json;not_null"`
	UserID		string
}

// NewDB initialize gorm and automigrate
func NewDB() (db *gorm.DB, err error) {
	//For connecting to MySQL DB
	//db, err = gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("sqlite3", "file:test.db?_busy_timeout=300000")
	if err != nil {
		return
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Home{})

	return
}
