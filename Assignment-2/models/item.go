package models

type Items struct {
	Id uint `gorm:"primaryKey;autoIncrement:true;not null" json:"itemId"`
	OrdersID uint
	ItemCode string `json:"itemCode"`
	Description string `gorm:"varchar(191)" json:"description"`
	Quantity uint `json:"quantity"`
}