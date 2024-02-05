package repository

import (
	"library/models"
	"library/utility"

	"gorm.io/gorm"
)


type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func (b *AuthorRepositoryImpl) New() *AuthorRepositoryImpl{
	ctx := utility.DbConector{}
	db := ctx.Connect()
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	sql.Ping()

	return &AuthorRepositoryImpl{db: db}
	
}

func (repo *AuthorRepositoryImpl) Add(author *models.Authors) {
	repo.db.Create(author)
}

func (repo *AuthorRepositoryImpl) List() *gorm.DB {
	return repo.db.Find(&models.Authors{})
}

func (repo *AuthorRepositoryImpl) Update(model *models.Authors) (*gorm.DB, error) {
	author := models.Authors{}
	result := repo.db.Table("Authors").First(model.ID).Take(&author)
	if result.Error != nil {
		panic("Livro não localizado")
	}
	return repo.db.UpdateColumns(model), nil
}

func (repo *AuthorRepositoryImpl) Delete(id uint) (*gorm.DB, error) {
	author := models.Authors{}
	result := repo.db.Table("Authors").First(id).Take(&author)
	if result.Error != nil {
		panic("Livro não localizado")
	}
	return repo.db.Delete(&author), nil
}

func (repo *AuthorRepositoryImpl) Get(id uint) (*gorm.DB, error) {
	result := repo.db.Table("Authors").First(id)
	if result.Error != nil {
		panic("Livro não localizado")
	}
	return result, nil
}
