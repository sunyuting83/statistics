package database

// Img List
type Admin struct {
	ID          int64  `json:"id" gorm:"primary_key, column:id"`
	Username    string `json:"username" gorm:"varchar(128);index:idx_username_id;column:username"`
	Password    string `json:"password" gorm:"column:password"`
	Fuck        string `json:"fuck" gorm:"column:fuck"`
	UpdateTime  int64  `json:"updatetime" gorm:"column:updatetime"`
	CreatedTime int64  `json:"createdtime" gorm:"column:createdtime"`
}

// TableName change table name
func (Admin) TableName() string {
	return "admin"
}

// add admin
func (admin *Admin) Insert() error {
	Eloquent.Create(&admin)
	return nil
}

// login
func (admin *Admin) CheckAdminLogin(username, password string) (admins Admin, err error) {
	if err = Eloquent.First(&admins, "username = ? AND fuck = ? AND password = ?", username, "0", password).Error; err != nil {
		return
	}
	return
}

// Check UserName
func (admin *Admin) CheckUserName(username string) (admins Admin, err error) {
	if err = Eloquent.First(&admins, "username = ?", username).Error; err != nil {
		return
	}
	return
}

// Check ID
func (admin *Admin) CheckID(id int64) (admins Admin, err error) {
	if err = Eloquent.First(&admins, "id = ?", id).Error; err != nil {
		return
	}
	return
}

// Reset Password
func (admin *Admin) ResetPassword(username string) (admins Admin, err error) {
	// time.Sleep(time.Duration(100) * time.Millisecond)
	if err = Eloquent.First(&admins, "username = ?", username).Error; err != nil {
		return
	}
	if err = Eloquent.Model(&admins).Updates(&admin).Error; err != nil {
		return
	}
	return
}

// Update Status
func (admin *Admin) UpStatusOne(id int64) (admins Admin, err error) {
	if err = Eloquent.First(&admins, "id = ?", id).Error; err != nil {
		return
	}
	if err = Eloquent.Model(&admins).Updates(&admin).Error; err != nil {
		return
	}
	return
}

// Delete Admin
func (admin *Admin) DeleteOne(id int64) {
	// time.Sleep(time.Duration(100) * time.Millisecond)
	Eloquent.Where("id = ?", id).Delete(&admin)
}

// Get Count
func (admin *Admin) GetCount() (count int64, err error) {
	if err = Eloquent.Model(&admin).Count(&count).Error; err != nil {
		return
	}
	return
}

// Card List
func (admin *Admin) GetAdminList(page int64) (admins []Admin, err error) {
	p := makePage(page)
	if err = Eloquent.
		Select("id, username, fuck, createdtime").
		Order("id desc").
		Limit(100).Offset(p).
		Find(&admins).Error; err != nil {
		return
	}
	return
}
