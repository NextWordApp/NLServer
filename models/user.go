package models

import (
	"fmt"
	"gorm.io/gorm"
)

var Uid uint

// Users 表模型
type User struct {
	UserID       uint   `gorm:"primaryKey;autoIncrement"`
	UserName     string `gorm:"size:255;not null;unique"`
	PasswordHash string `gorm:"size:255;not null"`
	PubKey       string `gorm:"size:255"`
	Email        string `gorm:"size:255;not null;unique"`
}

// 插入行
func InsertUser(user *User) {
	if err := MysqlDB.Create(user).Error; err != nil {
		fmt.Println("插入用户失败:", err)
	} else {
		fmt.Println("插入用户成功")
	}
}

// 删除行
func DeleteUser(userID uint) {
	if err := MysqlDB.Delete(&User{}, userID).Error; err != nil {
		fmt.Println("删除用户失败:", err)
	} else {
		fmt.Println("删除用户成功")
	}
}

// 修改行
func UpdateUser(userID uint, newUsername, newPasswordHash, newEmail string) {
	// 查找要更新的用户
	var user User
	if err := MysqlDB.First(&user, userID).Error; err != nil {
		fmt.Println("用户查找失败:", err)
		return
	}

	// 更新用户信息
	user.UserName = newUsername
	user.PasswordHash = newPasswordHash
	user.Email = newEmail

	// 保存更新后的用户信息
	if err := MysqlDB.Save(&user).Error; err != nil {
		fmt.Println("更新用户失败:", err)
	} else {
		fmt.Println("更新用户成功")
	}
}

// 验证 字段相关ID 是否存在
func PubKeyExists(walletPubKey string) (bool, error) {
	var user User
	result := MysqlDB.Where("pub_key = ?", walletPubKey).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

// FindUserIDByPubKey 根据 PubKey 查找 UserID
func FindUserIDByPubKey(pubKey string) (uint, error) {
	var user User
	if err := MysqlDB.Where("pub_key = ?", pubKey).First(&user).Error; err != nil {
		return 0, err
	}
	return user.UserID, nil
}
