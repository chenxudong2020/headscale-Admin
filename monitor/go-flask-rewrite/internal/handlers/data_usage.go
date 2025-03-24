package handlers

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DataUsage(c *gin.Context) {
	// 读取 data.json 文件
	file, err := os.Open("data.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open data.json"})
		return
	}
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	var jsonData map[string]map[string]string
	json.Unmarshal(content, &jsonData)

	// 处理 recv 和 sent 数据
	newRecv := processData(jsonData["recv"])
	newSent := processData(jsonData["sent"])

	newData := gin.H{
		"recv": newRecv,
		"sent": newSent,
	}

	c.JSON(http.StatusOK, newData)
}

func processData(data map[string]string) map[string]string {
	keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y"}
	newData := make(map[string]string)

	for i := 0; i < len(keys)-1; i++ {
		val1 := data[keys[i]]
		val2 := data[keys[i+1]]
		diff := math.Ceil(float64(toInt(val2) - toInt(val1)))
		newData[keys[i]] = string(diff)
	}

	return newData
}
