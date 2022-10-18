package models

import "time"

type Orders struct {
	Id uint `gorm:"primaryKey;autoIncrement:true;not null" json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt time.Time `json:"orderedAt"`
	Items []Items `json:"items"`
}