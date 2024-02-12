package models

type Authors struct {
	ID                         uint   `gorm:"column:ID"`
	Name                       string `gorm:"column:NAME"`
	Ammount_of_Published_Books uint   `gorm:"column:AMMOUNT_OF_PUBLISHED_BOOKS"`
	Has_Any_Book_In_Stock      bool   `gorm:"column:HAS_ANY_BOOK_IN_STOCK"`
}