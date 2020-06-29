package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"reflect"
	"time"

	jsoniter "github.com/json-iterator/go"
)

//这个函数会破坏file文件，所以不能瞎用
func FileSha1(file io.Reader) string {
	_sha1 := sha1.New()
	_, _ = io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func StrMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func PathToCommon(str string) string {
	return filepath.FromSlash(str)
}
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

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

/**字符串->时间对象*/
func Str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime

}

/**字符串->时间戳*/
func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := Str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

/**时间对象->字符串*/
func Time2Str() string {
	const shortForm = "2006-01-01 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

/*时间对象->时间戳*/
func Time2Stamp() int64 {
	t := time.Now()
	millisecond := t.UnixNano() / 1e6
	return millisecond
}

/*时间戳->字符串*/
func Stamp2Str(stamp int64) string {
	timeLayout := "2006-01-02 15:04:05"
	str := time.Unix(stamp/1000, 0).Format(timeLayout)
	return str
}

/*时间戳->时间对象*/
func Stamp2Time(stamp int64) time.Time {
	stampStr := Stamp2Str(stamp)
	timer := Str2Time(stampStr)
	return timer
}

//获取file对象
func GetFileByHeader(header *multipart.FileHeader) (multipart.File, error) {
	file, err := header.Open()
	if err != nil {
		return nil, err
	}

	return file, nil
}
