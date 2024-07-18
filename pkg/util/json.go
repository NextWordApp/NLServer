package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"service/models"
)

func AnalyzeJsonFile(path string) {
	// 读取 JSON 文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open file %s: %v\n", filepath.Base(path), err)
		return
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close file %s: %v\n", filepath.Base(path), err)
		}
	}(file)

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("failed to read file %s: %v\n", filepath.Base(path), err)
		return
	}

	var words []models.WordMsg

	// 解析 JSON 数据
	err = json.Unmarshal(byteValue, &words)
	if err != nil {
		fmt.Printf("failed to unmarshal json from file %s: %v\n", filepath.Base(path), err)
		return
	}

	// 遍历并插入数据
	for _, word := range words {
		// 插入时不需要设置 ID，数据库会自动生成
		if err := models.InsertWord(&word); err != nil {
			fmt.Printf("failed to insert word %s: %v\n", word.Word, err)
		}
	}
}
