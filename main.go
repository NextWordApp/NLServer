package main

import (
	"fmt"
	"service/dao/mysql"
	"service/pkg/setting"
	"service/pkg/util"
	"service/routers"
)

func init() {

}

func main() {
	// 1. 加载配置
	if err := setting.Init("config/config.yaml"); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	// 3. 加载数据库
	err := mysql.Init(setting.Config)
	if err != nil {
		return
	}

	// 4. 初始化数据库
	util.AnalyzeJsonFile(setting.Config.RES.WordPath)

	// 5. 注册路由
	routers.Init()
}
