package model

//Base ...
type Base struct {
	ID string `gorm:"column:Id;primary_key;"`
}

//user.ID = uuid.New().String()
