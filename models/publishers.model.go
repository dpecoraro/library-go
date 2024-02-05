package models

type Publishers struct {
	ID                         uint
	Name                       string
	Cnpj                       string
	Ammount_of_Published_Books uint
	Has_Any_Book_In_Stock      bool
}