package database

// Img List
type Host struct {
	ID          int64  `json:"id" gorm:"primary_key, column:id"`
	Host        string `json:"host" gorm:"column:host;index:idx_host"`
	HostName    string `json:"hostname" gorm:"column:hostname"`
	CreatedTime int64  `json:"createdtime" gorm:"column:createdtime"`
}

// TableName change table name
func (Host) TableName() string {
	return "host"
}

// add admin
func (host *Host) Insert() (id int64, err error) {
	c := Eloquent.Create(&host)
	id = host.ID
	if c.Error != nil {
		err = c.Error
		return
	}
	return
}

// check if has host
func (host *Host) CheckHost(hashost string) (id int64, err error) {
	if err = Eloquent.First(&host, "host = ?", hashost).Error; err != nil {
		return
	}
	id = host.ID
	return
}

// check if has host
func (host *Host) FindHost(hashost string) (hosts Host, err error) {
	if err = Eloquent.First(&hosts, "host = ?", hashost).Error; err != nil {
		return
	}
	return
}
