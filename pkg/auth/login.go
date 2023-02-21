package auth

type Login struct {
	Username string `json:"username" db:"username" gorm:"not null"`
	Password string `json:"password" db:"password" gorm:"not null"`
}
