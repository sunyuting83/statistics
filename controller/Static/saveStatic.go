package controller

import (
	"net/http"
	"statistics/database"
	"statistics/utils"

	"github.com/gin-gonic/gin"
)

type FormSaveData struct {
	StartDate int64  `form:"startdate" json:"startdate" xml:"startdate"  binding:"required"`
	EndDate   int64  `form:"enddate" json:"enddate" xml:"enddate"  binding:"required"`
	Host      string `form:"host" json:"host" xml:"host"  binding:"required"`
}

type DateData struct {
	Host     string      `json:"host"`
	HostName string      `json:"hostname"`
	Data     []*SaveData `json:"data"`
}

type SaveData struct {
	DateTime string            `json:"datetime"`
	Number   int64             `json:"number"`
	Data     []database.Static `json:"data"`
}

func SaveIt(c *gin.Context) {
	var form FormSaveData
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	d, err := MakeDateData(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "",
		"data":    d,
	})
}

func MakeDateData(form FormSaveData) (d *DateData, err error) {
	var gethost *database.Host
	Host, err := gethost.FindHost(form.Host)
	// fmt.Println(Host.ID)
	if err != nil {
		return nil, err
	}
	var static *database.Static
	StartDate, EndDate := utils.GetDateTimeUnix(form.StartDate, form.EndDate)
	datelist := utils.GetBetweenDates(StartDate, EndDate)
	// fmt.Println(datelist)
	data, err := static.GetStatic(StartDate, EndDate, Host.ID)
	if err != nil {
		return nil, err
	}
	// fmt.Println(data)
	var saveData []*SaveData
	for _, date := range datelist {
		var (
			d *SaveData = &SaveData{}
		)
		for _, item := range data {
			if date == utils.GetDateStr(item.CreatedTime) {
				d.DateTime = date
				d.Number = d.Number + item.Count
				d.Data = append(d.Data, item)
			}
		}
		if len(d.DateTime) > 0 {
			saveData = append(saveData, d)
		}
	}
	d = &DateData{
		Host:     Host.Host,
		HostName: Host.HostName,
		Data:     saveData,
	}
	return d, nil
}
