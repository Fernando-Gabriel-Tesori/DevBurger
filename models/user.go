package models

// User representa a estrutura do usuário no banco de dados.
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Username string `gorm:"unique" json:"username"` // ← adicione isto se fizer sentido
	Orders   []Order
}
