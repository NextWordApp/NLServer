package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/ai"
	"service/models"
)

// 公钥
type PublicKey struct {
	Address string `json:"address"`
}

func Login(c *gin.Context) {
	var request PublicKey
	// 将请求体绑定到结构体
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address := request.Address

	// 查询数据库是否存在
	exists, err := models.PubKeyExists(address)
	if err != nil {
		c.JSON(http.StatusOK, nil)
		return
	}

	fmt.Println("账户不存在，已自动注册")
	// 不存在: 插入
	if !exists {
		user := models.User{
			UserID:       models.Uid,
			UserName:     "",
			PasswordHash: "",
			PubKey:       address,
			Email:        "",
		}
		models.Uid++
		models.InsertUser(&user)
	}

	// 存在 nil
	fmt.Println("账户存在，退出")
}

func GetMessage(c *gin.Context) {
	var request PublicKey
	// 将请求体绑定到结构体
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 拿到公钥
	address := request.Address

	fmt.Printf("address: %s\n", address)
	// 通过公钥拿到用户 UID
	id, err := models.FindUserIDByPubKey(address)
	if err != nil {
		return
	}

	word, err := models.SelectWordFromUser(id)
	if err != nil {
		return
	}
	fmt.Printf("address: %s\n", word)

	// 通过UID 查找单词
	wordMsg, err := models.FindWordById(id)
	if err != nil {
		return
	}

	// 传入 word
	msg := ai.Request(wordMsg.Word)
	fmt.Printf("msg: %s\n", msg)
	// 拼接 json

	// 返回json 给前端
	c.JSON(http.StatusOK, msg)

}
