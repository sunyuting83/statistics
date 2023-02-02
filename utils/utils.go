package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port       string `yaml:"port"`
	SECRET_KEY string `yaml:"SECRET_KEY"`
	CODE       string `yaml:"CODE"`
	AdminPWD   string `yaml:"AdminPWD"`
}

// GetCurrentPath Get Current Path
func GetCurrentPath() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(path)
	return dir, nil
}

// CheckConfig check config
func CheckConfig(OS, CurrentPath string) (conf *Config, err error) {
	LinkPathStr := "/"
	if OS == "windows" {
		LinkPathStr = "\\"
	}
	ConfigFile := strings.Join([]string{CurrentPath, "config.yaml"}, LinkPathStr)

	var confYaml *Config
	yamlFile, err := os.ReadFile(ConfigFile)
	if err != nil {
		return confYaml, errors.New("读取配置文件出错\n10秒后程序自动关闭")
	}
	err = yaml.Unmarshal(yamlFile, &confYaml)
	if err != nil {
		return confYaml, errors.New("读取配置文件出错\n10秒后程序自动关闭")
	}
	if len(confYaml.Port) <= 0 {
		confYaml.Port = "13002"
		config, _ := yaml.Marshal(&confYaml)
		os.WriteFile(ConfigFile, config, 0644)
	}
	if len(confYaml.SECRET_KEY) <= 0 {
		secret_key := RandSeq(32)
		confYaml.SECRET_KEY = secret_key
		config, _ := yaml.Marshal(&confYaml)
		os.WriteFile(ConfigFile, config, 0644)
	}
	if len(confYaml.CODE) <= 0 {
		code := "ASD123"
		confYaml.CODE = code
		config, _ := yaml.Marshal(&confYaml)
		os.WriteFile(ConfigFile, config, 0644)
	}
	if len(confYaml.AdminPWD) <= 0 {
		confYaml.AdminPWD = "admin888"
		config, _ := yaml.Marshal(&confYaml)
		os.WriteFile(ConfigFile, config, 0644)
	}
	return confYaml, nil
}

// CORSMiddleware cors middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IsExist(path string) bool {
	// 判断文件是否存在
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func GetDateTime() (int64, int64) {
	d := time.Now()
	date := d.Format("2006-01-02")
	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	// yesterday := d.AddDate(0, 0, -1)
	// yday := yesterday.Format("2006-01-02")
	// yDate := yday + "_00:00:00"
	// yTime, _ := time.ParseInLocation("2006-01-02_15:04:05", yDate, loc)

	//返回当天0点和23点59分的时间戳
	return startTime.Unix(), end.Unix()
}

func GetDateTimeUnix(start, end int64) (int64, int64) {
	startTimeStr, endTimeStr := time.Unix(start, 0), time.Unix(end, 0)
	startdate, enddate := startTimeStr.Format("2006-01-02"), endTimeStr.Format("2006-01-02")

	loc, _ := time.LoadLocation("Local")

	startDate, endDate := startdate+"_00:00:00", enddate+"_23:59:59"

	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)
	endTime, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	return startTime.Unix(), endTime.Unix()
}

func GetDateStr(start int64) string {
	startTimeStr := time.Unix(start, 0)
	startdate := startTimeStr.Format("2006-01-02")

	return startdate
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate int64) []string {
	startTimeStr, endTimeStr := time.Unix(sdate, 0), time.Unix(edate, 0)
	sdatea, edated := startTimeStr.Format("2006-01-02"), endTimeStr.Format("2006-01-02")
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdatea) {
		timeFormatTpl = timeFormatTpl[0:len(sdatea)]
	}
	date, err := time.Parse(timeFormatTpl, sdatea)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edated)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func MD5(a string) string {
	data := []byte(a)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
