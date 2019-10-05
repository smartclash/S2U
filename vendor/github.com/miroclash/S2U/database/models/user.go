package models

type User struct {
	Structure

	Name           string
	Email          string `gorm:"unique_index"`
	Password       string
	Bannerpicture  string `gorm:"column:bannerpicture"`
	Displaypicture string `gorm:"column:displaypicture"`
}
