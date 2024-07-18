package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type File struct {
	Type           string `json:"type"`
	TransferMethod string `json:"transfer_method"`
	URL            string `json:"url"`
}

type RequestBody struct {
	Inputs         map[string]interface{} `json:"inputs"`
	Query          string                 `json:"query"`
	ResponseMode   string                 `json:"response_mode"`
	ConversationID string                 `json:"conversation_id"`
	User           string                 `json:"user"`
	Files          []File                 `json:"files"`
}

func Request(word string) (msg string) {

	url := "https://api.dify.ai/v1/chat-messages"

	// 创建请求体
	requestBody := RequestBody{
		Inputs:         map[string]interface{}{},
		Query:          word,
		ResponseMode:   "blocking",
		ConversationID: "",
		User:           "abc-456",
		Files: []File{
			{
				Type:           "",
				TransferMethod: "",
				URL:            "",
			},
		},
	}

	// 将请求体序列化为 JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	// 设置头部信息
	req.Header.Set("Authorization", "Bearer app-e7uTHWwBs5BnLY2PNZ6TEifF")
	req.Header.Set("Content-Type", "application/json")

	// 发起 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	//// 提取 answer 字段并转化为 JSON
	//answer, ok := result["answer"]
	//if !ok {
	//	return
	//}

	responseJSON, err := json.Marshal(result)
	if err != nil {
		return
	}

	return string(responseJSON)
}
