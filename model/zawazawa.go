package model

// ZawazawaUser [...]
type ZawazawaUser struct {
	ID     int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Mobile string `gorm:"index:mobile;column:mobile;type:varchar(32)" json:"mobile"` // 手机号码
}

// ZawazawaUserToken [...]
type ZawazawaUserToken struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	UId          int    `gorm:"column:uid;type:int(11)" json:"uid"`
	Token        string `gorm:"column:token;type:varchar(255)" json:"token"`
	RefreshToken string `gorm:"column:refresh_token;type:varchar(255)" json:"refresh_token"`
	Type         string `gorm:"column:type;type:varchar(255)" json:"type"`
}
