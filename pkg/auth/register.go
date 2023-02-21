package auth

type Register struct {
	Name     string `json:"name" db:"name" gorm:"not null"`
	Username string `json:"username" db:"username" gorm:"not null"`
	Password string `json:"password" db:"password" gorm:"not null"`
}
