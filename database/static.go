package database

// Img List
type Static struct {
	ID          int64  `json:"id" gorm:"primary_key, column:id"`
	HostID      int64  `json:"hostid" gorm:"column:hostid;index:idx_hostid"`
	Card        string `json:"card" gorm:"varchar(32);column:card;index:idx_card"`
	ClientIP    string `json:"clientip" gorm:"varchar(14);column:clientip"`
	Count       int64  `json:"count" gorm:"default:0;column:count"`
	CreatedTime int64  `json:"createdtime" gorm:"column:createdtime;index:idx_createdtime"`
}

// TableName change table name
func (Static) TableName() string {
	return "static"
}

// add admin
func (static *Static) Insert() (id int64, err error) {
	c := Eloquent.Create(&static)
	id = static.ID
	if c.Error != nil {
		err = c.Error
		return
	}
	return
}

func (static *Static) CheckHasStatic(startTime, endTime, hostid int64, card string) (statics Static, err error) {
	if err = Eloquent.First(&statics, "createdtime >= ? AND createdtime <= ? AND hostid = ? AND card = ?", startTime, endTime, hostid, card).Error; err != nil {
		return
	}
	return
}

func (static *Static) GetStatic(startTime, endTime, hostid int64) (statics []Static, err error) {
	if err = Eloquent.Find(&statics, "createdtime >= ? AND createdtime <= ? AND hostid = ?", startTime, endTime, hostid).Error; err != nil {
		return
	}
	return
}

func (static *Static) UpCount() (statics Static, err error) {
	static.Count = static.Count + 1
	if err = Eloquent.Model(&statics).Updates(&static).Error; err != nil {
		return
	}
	return
}

func (static *Static) GetTodayStatic(start, end int64) (statics []Static, err error) {
	if err = Eloquent.
		Select("id, count").
		Where("createdtime >= ? AND createdtime <= ?", start, end).
		Find(&statics).Error; err != nil {
		return
	}
	return
}

// makePage make page
func makePage(p int64) int64 {
	p = p - 1
	if p <= 0 {
		p = 0
	}
	page := p * 100
	return page
}
