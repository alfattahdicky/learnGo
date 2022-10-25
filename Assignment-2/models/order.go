package models

type Orders struct {
	Id uint `gorm:"primaryKey;autoIncrement:true;not null" json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt string `json:"orderedAt"`
	Items []Items `json:"items" gorm:"foreignKey:Id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}