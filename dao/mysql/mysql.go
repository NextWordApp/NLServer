package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"service/models"
	"service/pkg/setting"
)

func Init(config *setting.AppConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLConfig.User,
		config.MySQLConfig.Password,
		config.MySQLConfig.Host,
		config.MySQLConfig.Port,
		config.MySQLConfig.DBName,
	)
	MysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	models.MysqlDB = MysqlDB
	if err != nil {
		return err
	}

	// Auto migrate models
	err = MysqlDB.AutoMigrate(&models.User{}, &models.WordMsg{}, &models.UserWord{})
	if err != nil {
		return err
	}

	return nil
}
