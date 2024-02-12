package repository

import (
	"library/models"
	"library/utility"

	"gorm.io/gorm"
)


type BookRepositoryImpl struct {
	db *gorm.DB
}

func (b *BookRepositoryImpl) New() *BookRepositoryImpl{
	ctx := utility.DbConector{}
	db := ctx.Connect()
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	sql.Ping()
	return &BookRepositoryImpl{db: db}
	
}

func (repo *BookRepositoryImpl) Add(book *models.Books) {
	repo.db.Create(book)
}

func (repo *BookRepositoryImpl) List() *gorm.DB {
	return repo.db.Find(&models.Books{})
}

func (repo *BookRepositoryImpl) Update(model *models.Books) (*gorm.DB, error) {
	book := models.Books{}
	result := repo.db.Commit().First(model.Id).Take(&book)
	if result.RowsAffected  == 0 {
		panic("Livro não localizado")
	}
	return repo.db.UpdateColumns(&model), nil
}

func (repo *BookRepositoryImpl) Delete(id uint) (*gorm.DB, error) {
	book := models.Books{Id: id}
	result := repo.db.Commit().First(&book.Id)
	if result.RowsAffected  == 0 {
		panic("Livro não localizado")
	}
	return repo.db.Delete(book), nil
}

func (repo *BookRepositoryImpl) Get(id uint) (models.Books, error) {
	var book models.Books

	err := repo.db.Model(models.Books{Id: id}).First(&book)

	if err.Error != nil {
		panic("Livro não localizado")
	}
	
	return book, nil
}
