package domain

type Stock struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null"`
}
