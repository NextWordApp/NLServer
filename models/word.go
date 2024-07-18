package models

import (
	"fmt"
)

// Words 表模型
type WordMsg struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Word     string `gorm:"size:255;not null"`
	Phonetic string `gorm:"size:255"`
	Pos      string `gorm:"size:50"`
	Meaning  string `gorm:"type:text"`
}

func InsertWord(word *WordMsg) error {
	if err := MysqlDB.Create(word).Error; err != nil {
		fmt.Println("insert word fail:", err)
	} else {
		fmt.Println("insert word success:", word.Word)
	}
	return nil
}

func DeleteWord(wordID uint) {
	if err := MysqlDB.Delete(&WordMsg{}, wordID).Error; err != nil {
		fmt.Println("delete word fail:", err)
	} else {
		fmt.Println("delete word success:")
	}
}

func GetNameByID(id uint) (string, error) {
	var word WordMsg
	if err := MysqlDB.First(&word, id).Error; err != nil {
		return "", err
	}
	return word.Word, nil
}

func FindWordById(wordId uint) (*WordMsg, error) {
	var wordMsg WordMsg
	result := MysqlDB.First(&wordMsg, wordId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wordMsg, nil
}
