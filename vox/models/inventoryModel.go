package models

type User struct {
	ID   uint8  `gorm:"primary_key"`
	Name string `gorm:"type:varchar(20);not null"`
}

type Laptops struct {
	ID     uint8  `gorm:"primary_key"`
	InName string `gorm:"type:varchar(10);not null"`
	Sum    uint8
	Uid    uint8
	User   User `gorm:"foreignKey:Uid"`
}
