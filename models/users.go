package models

// Model Struct
type Users struct {
	User_id  int    `orm:"auto"`
	Email    string `orm:"size(254)"`
	Token    string `orm:"size(254)"`
	Username string `orm:"size(254)"`
	Bio      string `orm:"size(254)"`
	Image    string `orm:"size(254)"`
}
