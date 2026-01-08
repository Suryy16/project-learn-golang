package models

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Todos     []Todo `gorm:"foreignKey:UserID" json:"todos,omitempty"`
}
