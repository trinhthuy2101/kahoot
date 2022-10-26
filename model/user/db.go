package user

import (
	"crypto/md5"
	"encoding/hex"

	"gorm.io/gorm"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

type Database interface {
	Login(request *User) bool
	Register(request *User) bool
}
type db struct {
	db *gorm.DB
}

func (db *db) Login(request *User) bool {
	var result User
	encryptedPass := GetMD5Hash(request.Password)
	err := db.db.Where("username=? and password=?", request.Username, encryptedPass).First(&result).Error
	if err != nil {
		return false
	}
	return true
}
func (db *db) Register(request *User) bool {
	var result User
	encryptedPass := GetMD5Hash(request.Password)
	err := db.db.Create(&User{ID: request.ID, Username: request.Username, Password: encryptedPass}).Scan(&result).Error
	if err != nil {
		return false
	}
	return true
}
func NewDatabase(database *gorm.DB) Database {
	return &db{
		db: database,
	}
}
