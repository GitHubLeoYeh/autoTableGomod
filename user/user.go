package user

type UserInfo struct {
	ID       *uint64 `gorm:"column:id;type:int;primary_key;"`
	Name     *string `gorm:"column:name;type:varchar(255);"`
	Email    *string `gorm:"column:email;type:varchar(255);default:''"`
	Password *string `gorm:"column:password;type:varchar(255);default:''"`
}

func (UserInfo) TableName() string {
	return "user"
}
