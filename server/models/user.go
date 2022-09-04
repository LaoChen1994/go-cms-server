package models

type User struct {
	Model
	Mobile   string `gorm: "column: mobile" json: "mobile"`
	Username string `gorm: "column: username" json: "username"`
	Passwprd string `gorm: "column: password" json: "password"`
}
