package models

type Books struct {
	Id           uint       `json:"Id" gorm:"primaryKey"`
	Title        string     `json:"Title" gorm:"column:TITLE"`
	Author_Id    uint       `json:"AuthorId" gorm:"column:AUTHOR_ID"`
	Author       Authors    `json:"Author" gorm:"foreignKey:ID"`
	Publisher_Id uint       `json:"PublisherId" gorm:"column:PUBLISHER_ID"`
	Publisher    Publishers `json:"Publisher" gorm:"foreignKey:ID"`
	PriceSale    float32    `json:"PriceSale" sql:"type:decimal(10,2);" gorm:"column:PRICE_SALE"`
	PriceRent    float32    `json:"PriceRent" sql:"type:decimal(10,2);" gorm:"column:PRICE_RENT"`
	Currency     string     `json:"Currency" gorm:"column:CURRENCY"`
	Quantity     uint       `json:"Quantity" gorm:"column:QUANTITY"`
	InStock      bool       `json:"InStock" gorm:"column:IN_STOCK"`
	IsBestSeller bool       `json:"IsBestSeller" gorm:"column:IS_BEST_SELLER"`
	SaleOrRent   string     `json:"SaleOrRent" gorm:"column:SALE_OR_RENT"`
}