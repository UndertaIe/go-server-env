package model

import (
	"github.com/UndertaIe/passwd/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	*database.BaseModel
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	ShareMode   int    `json:"share_mode"`
	Sex         int    `json:"gender"`
	Desc        string `json:"desc"`
	Role        int    `json:"role"`
}

func (u User) TableName() string {
	return "passwd_user"
}

func (u User) Get(db gorm.DB) (User, error) {
	err := db.Where("user_id = ? AND is_deleted = ?", u.UserId, false).First(&u).Error
	return u, err
}

func (u User) Create(db gorm.DB) (User, error) {
	err := db.Create(&u).Error
	return u, err
}

func (u User) Update(db gorm.DB, values interface{}) error {
	err := db.Model(&u).Where("user_id = ? AND is_deleted = ?", u.UserId, false).Updates(values).Error // 待测试
	return err
}

func (u User) Delete(db gorm.DB) error {
	err := db.Where("user_id = ? AND is_deleted = ?", u.UserId, false).Delete(&u).Error
	return err
}

type UserAccountRow struct {
	PlatformType,
	PlatformName,
	Password,
	PlatformDomain,
	PlatformLoginUrl,
	PlatformDesc string
}

func (u User) ListByUserID(db gorm.DB) ([]*UserAccountRow, error) {
	var resp []*UserAccountRow
	rows, err := db.Where(
		"user_id = ?", u.UserId).Joins(
		"JOIN passwd_platform p ON passwd_user_account.platform_id = p.platform_id AND passwd_user_account.user_id = ?", u.UserId).Select(
		"p.PlatformType, p.PlatformName, passwd_user_account.Password, p.PlatformDomain, p.PlatformLoginUrl, p.PlatformDesc").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		uar := &UserAccountRow{}
		err = rows.Scan(&uar.PlatformType, &uar.PlatformName, &uar.Password, &uar.PlatformDomain, &uar.PlatformLoginUrl, &uar.PlatformDesc)
		if err != nil {
			return nil, err
		}

		resp = append(resp, uar)
	}
	return resp, err
}
