package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

//ParseJSONFile 解析JSON文件到对象
func ParseJSONFile(filename string, v interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		log.Errorf("JSON 文件 %s 读取失败: %v", filename, err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	err = dec.Decode(v)

	if err != nil {
		log.Errorf("JSON 文件 %s 解析失败: %v", filename, err)
	}
}

// PrettyStruct 输出漂亮的Json格式
func PrettyStruct(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("Struct marshal error: %v \n\n", err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	return out.String()
}
