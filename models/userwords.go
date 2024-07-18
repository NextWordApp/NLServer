package models

import (
	"fmt"
)

// UserWords 表模型
type UserWord struct {
	UserID uint `gorm:"not null;primaryKey;autoIncrement"`
	WordID uint `gorm:"not null"` //
}

func InsertUserWord(userWord *UserWord) {
	if err := MysqlDB.Create(userWord).Error; err != nil {
		fmt.Println("插入用户单词关系失败:", err)
	} else {
		fmt.Println("插入用户单词关系成功")
	}
}

func DeleteUserWord(userWordID uint) {
	if err := MysqlDB.Delete(&UserWord{}, userWordID).Error; err != nil {
		fmt.Println("删除用户单词关系失败:", err)
	} else {
		fmt.Println("删除用户单词关系成功")
	}
}

func FindMinUnlearnedWordID(userID uint) (uint, error) {
	var minWordID uint

	// 子查询：获取用户已经学习的单词ID
	subQuery := MysqlDB.Model(&UserWord{}).Select("word_id").Where("user_id = ?", userID)

	// 找出用户未学习的单词的最小ID
	result := MysqlDB.Model(&WordMsg{}).Where("id NOT IN (?)", subQuery).Select("MIN(id)").Scan(&minWordID)
	if result.Error != nil {
		return 0, result.Error
	}

	return minWordID, nil
}

func SelectWordFromUser(userId uint) (string, error) {
	wordId, err := FindMinUnlearnedWordID(userId)
	if err != nil {
		return "", err
	}
	word, err := GetNameByID(wordId)
	if err != nil {
		return "", err
	}
	return word, nil
}
