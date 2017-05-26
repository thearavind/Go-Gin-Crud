package models

// Users - Model for the uses table
type Users struct {
	UserId   int    `orm:"auto"`
	Email    string `orm:"size(128)"`
	Password string `orm:"size(64)"`
	UserName string `orm:"size(32)"`
	Bio      string `orm:"size(256);null"`
	Image    string `orm:"size(256);null"`
}
