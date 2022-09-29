package model

import (
	"github.com/UndertaIe/passwd/database"
	"github.com/UndertaIe/passwd/pkg/page"
	"gorm.io/gorm"
)

type User struct {
	UserId        int    `json:"user_id" gorm:"column:user_id"`
	UserName      string `json:"user_name" gorm:"column:user_name"`
	Password      string `json:"password" gorm:"column:password"`
	Salt          string `json:"salt" gorm:"column:salt"`
	PhoneNumber   string `json:"phone_number" gorm:"column:phone_number"`
	Email         string `json:"email" gorm:"column:email; default:null"`
	ShareMode     int    `json:"share_mode" gorm:"column:share_mode; default:0"`
	Role          int    `json:"role" gorm:"column:role"`
	ProfileImgUrl string `json:"profile_img_url" gorm:"column:profile_img_url; default:''"`
	Description   string `json:"description" gorm:"column:description; default:''"`
	Sex           int    `json:"sex" gorm:"column:sex; default:0"`
	*database.BaseModel
}

func (u User) TableName() string {
	return "passwd_user"
}

func (u User) Get(db *gorm.DB) (User, error) {
	err := db.Where("user_id = ? AND is_deleted = ?", u.UserId, false).Take(&u).Error
	return u, err
}

func (u User) GetUserByPhone(db *gorm.DB) (User, error) {
	err := db.Where("phone_number = ? AND is_deleted = ?", u.PhoneNumber, false).Take(&u).Error
	return u, err
}

func (u User) PhoneExists(db *gorm.DB) (bool, error) {
	err := db.Where("phone_number = ? AND is_deleted = ?", u.PhoneNumber, false).Take(&u).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u User) EmailExists(db *gorm.DB) (bool, error) {
	err := db.Where("email = ? AND is_deleted = ?", u.Email, false).Take(&u).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u User) GetUserByEmail(db *gorm.DB) (User, error) {
	err := db.Where("email = ? AND is_deleted = ?", u.Email, false).Take(&u).Error
	return u, err
}

func (u User) GetUserByName(db *gorm.DB) (User, error) {
	err := db.Where("user_name = ? AND is_deleted = ?", u.UserName, false).Take(&u).Error
	return u, err
}

func (u User) Create(db *gorm.DB) error {
	err := db.Omit("user_id").Create(&u).Error
	return err
}

func (u User) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&u).Where("user_id = ? AND is_deleted = ?", u.UserId, false).Updates(values).Error
	return err
}

func (u User) Delete(db *gorm.DB) error {
	err := db.Model(&u).Where("user_id = ? AND is_deleted = ?", u.UserId, false).Update("is_deleted", true).Error
	return err
}

type UserRow struct {
	UserId      int    `json:"user_id" gorm:"column:user_id"`
	UserName    string `json:"user_name" gorm:"column:user_name"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Email       string `json:"email" gorm:"column:email"`
	ShareMode   int    `json:"share_mode" gorm:"column:share_mode"`
	Sex         int    `json:"sex" gorm:"column:sex"`
	Description string `json:"description" gorm:"column:description"`
	Role        int    `json:"role" gorm:"column:role"`
}

func (u User) GetUserList(db *gorm.DB, pager *page.Pager) ([]UserRow, error) {
	db = db.Offset(pager.Offset()).Limit(pager.Limit())
	var users []UserRow
	tx := db.Model(&User{}).Where(map[string]interface{}{"is_deleted": false}).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
