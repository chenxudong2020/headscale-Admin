package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DataRecord(c *gin.Context) {
	// 读取当前数据
	info := make(map[string]interface{})
	GetInfoData(&info)

	recvSpeed := info["recv_speed"].(float64)
	sentSpeed := info["sent_speed"].(float64)

	// 读取本地 data.json 文件
	file, err := os.Open("data.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open data.json"})
		return
	}
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	var jsonData map[string]map[string]string
	json.Unmarshal(content, &jsonData)

	// 更新数据
	keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y"}
	for i := 0; i < len(keys)-1; i++ {
		jsonData["sent"][keys[i]] = jsonData["sent"][keys[i+1]]
		jsonData["recv"][keys[i]] = jsonData["recv"][keys[i+1]]
	}

	jsonData["sent"]["y"] = string(sentSpeed)
	jsonData["recv"]["y"] = string(recvSpeed)

	// 写回文件
	file, _ = os.Create("data.json")
	defer file.Close()
	json.NewEncoder(file).Encode(jsonData)

	c.JSON(http.StatusOK, jsonData)
}
