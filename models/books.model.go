package models

type Books struct {
	Id           uint       `json:"Id" gorm:"primaryKey"`
	Title        string     `json:"Title"`
	Author_Id    uint       `json:"AuthorId"`
	Author       Authors    `json:"Author" gorm:"foreignKey:ID"`
	Publisher_Id uint       `json:"PublisherId"`
	Publisher    Publishers `json:"Publisher" gorm:"foreignKey:ID"`
	PriceSale    float32    `json:"PriceSale" sql:"type:decimal(10,2);"`
	PriceRent    float32    `json:"PriceRent" sql:"type:decimal(10,2);"`
	Currency     string     `json:"Currency"`
	Quantity     uint       `json:"Quantity"`
	InStock      bool       `json:"InStock"`
	IsBestSeller bool       `json:"IsBestSeller"`
	SaleOrRent   string     `json:"SaleOrRent"`
}