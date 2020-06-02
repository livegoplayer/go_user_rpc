package util

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func JsonEncode(data interface{}) []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonByte, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("json加密出错:" + err.Error())
	}
	return jsonByte
}

//解析valjson对象
func JsonDecode(val []byte, data interface{}) interface{} {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//把结果放入val
	err := json.Unmarshal(val, &data)
	if err != nil {
		fmt.Printf("json解析出错:" + err.Error())
	}
	return data
}

//直接获取json中的值
func GetJsonVal(data []byte, path ...interface{}) string {
	return jsoniter.Get(data, path...).ToString()
}
