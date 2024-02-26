package models

type CarModel struct {
	Id int
	Name string 
	ManufacturerID int // foreign key
	CategoryId int // foreign key
	Year int
	Specs Specifications // specification object 
	Phot string // address to photo
}