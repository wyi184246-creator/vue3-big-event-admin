package model

import "time"

type User struct {
	ID        uint      `json:"id" db:"id" gorm:"column:id;primaryKey;autoIncrement;comment:用户唯一ID"`
	Username  string    `json:"username" db:"username" gorm:"column:username;type:varchar(50);not null;uniqueIndex:idx_username;comment:用户名"`
	Nickname  string    `json:"nickname" db:"nickname" gorm:"column:nickname;type:varchar(50);default:'';comment:昵称"`
	Email     string    `json:"email" db:"email" gorm:"column:email;type:varchar(100);default:'';comment:电子邮箱"`
	UserPic   string    `json:"user_pic" db:"user_pic" gorm:"column:user_pic;type:varchar(255);default:'';comment:用户头像URL"`
	CreatedAt time.Time `json:"created_at" db:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime;comment:最后更新时间"`
}

func (User) TableName() string {
	return "users"
}
