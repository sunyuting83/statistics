package controller

import (
	"net/http"
	"statistics/database"
	"statistics/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Putit(c *gin.Context) {
	var (
		code       string = c.Query("code")
		host       string = c.Query("host")
		SECRET_KEY string = c.Query("secretkey")
	)
	clientIP := GetIp(c)
	secret_code, _ := c.Get("code")
	scode := secret_code.(string)
	if code != scode {
		if len(code) != 32 {
			c.String(http.StatusOK, "1")
			return
		}
	}
	if len(host) <= 0 {
		c.String(http.StatusOK, "1")
		return
	}
	if len(SECRET_KEY) <= 0 {
		c.String(http.StatusOK, "1")
		return
	}
	if !ItsIP(host) {
		c.String(http.StatusOK, "1")
		return
	}
	secret_key, _ := c.Get("secret_key")
	sk := secret_key.(string)
	if SECRET_KEY != sk {
		c.String(http.StatusOK, "1")
		return
	}

	var gethost database.Host
	T := time.Now().Unix()
	ID, err := gethost.CheckHost(host)
	if err != nil && err.Error() == "record not found" {
		gethost.CreatedTime = T
		gethost.Host = host
		id, _ := gethost.Insert()
		ID = id
	}

	startTime, endTime := utils.GetDateTime()

	var static database.Static
	sc, err := static.CheckHasStatic(startTime, endTime, ID, code)
	if err != nil && err.Error() == "record not found" {
		static.HostID = ID
		static.Card = code
		static.ClientIP = clientIP
		static.Count = 1
		static.CreatedTime = T
		static.Insert()
		c.String(http.StatusOK, "0")
		return
	}
	sc.UpCount()
	c.String(http.StatusOK, "0")
}

func GetIp(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

func ItsIP(host string) bool {
	var (
		x bool = false
	)
	if strings.Contains(host, ".") {
		if strings.Contains(host, ":") {
			ip := strings.Split(host, ":")[0]
			if strings.Contains(ip, ".") {
				st := strings.Split(ip, ".")
				if len(st) == 4 {
					x = true
				}
			}
		}
	}
	return x
}
