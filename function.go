package common

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"io"
	"math"
	mathRand "math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Md5Str md5 加密字符串
func Md5Str(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// Capitalize 字符首字母大写转换
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// CheckPassword 检测密码 6-18位且数字与英文组合
func CheckPassword(password string) bool {
	if password == "" {
		return false
	}
	result, _ := regexp.MatchString("^[a-zA-Z]\\w{5,17}$", password)
	return result
}
func MakePassword(password, salt string) string {
	return Md5Str(Md5Str(password) + salt)
}

// NowTime 当前时间
func NowTime() string {
	chinaTimezone, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(chinaTimezone).Format("2006-01-02 15:04:05")
}
func NowDate() string {
	chinaTimezone, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(chinaTimezone).Format("2006-01-02")
}
func NowTimestamp() int32 {
	chinaTimezone, _ := time.LoadLocation("Asia/Shanghai")
	return int32(time.Now().In(chinaTimezone).Unix())
}

// FileExists 判断文件夹或文件是否存在
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func AfterDays(day int) string {
	chinaTimezone, _ := time.LoadLocation("Asia/Shanghai")
	nTime := time.Now().In(chinaTimezone)
	yesTime := nTime.AddDate(0, 0, day)
	return yesTime.Format("2006-01-02")
}
func FileDel(path string) bool {
	exist, err := FileExists(path)
	if exist && err == nil {
		err = os.Remove(path)
		if err == nil {
			return true
		}
	}
	return false
}

// Round 四舍五入 保留两位小数
func Round(x float64) float64 {
	return cast.ToFloat64(strconv.FormatFloat(x, 'f', 2, 64))
}

// CheckMobile 检测手机号
func CheckMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	result, _ := regexp.MatchString(`^((\+86)|(86))?1([356789][0-9]|4[579]|6[67]|7[0135678]|9[189])[0-9]{8}$`, mobile)
	return result
}

// CheckPhone 检测固定电话
func CheckPhone(phone string) bool {
	if phone == "" {
		return false
	}
	result, _ := regexp.MatchString(`^(0\d{2,3}\-)?([2-9]\d{6,7})+(\-\d{1,6})?$`, phone)
	return result
}

// Empty 判断是否为空
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	mathRand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(b)
}
func TimeFormat(str string) string {
	tt, _ := time.Parse("2006-01-02T15:04:05Z07:00", str)
	return tt.Format("2006-01-02 15:04:05")
}
func DateFormat(str string) string {
	tt, _ := time.Parse("2006-01-02 15:04:05", str)
	return tt.Format("2006-01-02")
}

// ArrDuplicateInt int 数组去重
func ArrDuplicateInt(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// ArrDuplicateString string 数组去重
func ArrDuplicateString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
func CreateOrdersn(key string) string {
	u, _ := uuid.NewRandom()
	return key + time.Now().Format("20060102") + strconv.Itoa(int(u.ID()))
}
func FormatTimeUnix(str string, format string) int32 {
	formatTime, err := time.Parse(format, str)
	if err != nil {
		return 0
	}
	return int32(formatTime.Unix())
}
func FormatDifferenceDay(timeA, timeB, format string) int {
	a := FormatTimeUnix(timeA, format)
	b := FormatTimeUnix(timeB, format)
	if a < b {
		return 0
	}
	return cast.ToInt(math.Ceil(cast.ToFloat64((a - b) / int32(86400))))
}

// GetNumber 从字符串中获取数字
func GetNumber(str string) int {
	dictionary := "123456789"
	var num string
	for _, val := range str {
		if strings.Contains(dictionary, string(val)) {
			num += string(val)
		}
	}
	return cast.ToInt(num)
}
