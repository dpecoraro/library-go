package repository

import (
	"library/models"
	"library/utility"

	"gorm.io/gorm"
)


type PublisherRepositoryImpl struct {
	db *gorm.DB
}

func (b *PublisherRepositoryImpl) New() *PublisherRepositoryImpl{
	ctx := utility.DbConector{}
	db := ctx.Connect()
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	sql.Ping()

	return &PublisherRepositoryImpl{db: db}
	
}

func (repo *PublisherRepositoryImpl) Add(Publisher *models.Publishers) {
	repo.db.Create(Publisher)
}

func (repo *PublisherRepositoryImpl) List() *gorm.DB {
	return repo.db.Find(&models.Publishers{})
}

func (repo *PublisherRepositoryImpl) Update(model *models.Publishers) (*gorm.DB, error) {
	publisher := models.Publishers{}
	result := repo.db.Table("Publishers").First(model.ID).Take(&publisher)
	if result.Error  != nil {
		panic("Livro não localizado")
	}
	return repo.db.UpdateColumns(&model), nil
}

func (repo *PublisherRepositoryImpl) Delete(id uint) (*gorm.DB, error) {
	publisher := models.Publishers{}
	result := repo.db.Table("Publishers").First(id).Take(&publisher)
	if result.Error  != nil {
		panic("Livro não localizado")
	}
	return repo.db.Delete(publisher), nil
}

func (repo *PublisherRepositoryImpl) Get(id uint) (*gorm.DB, error) {
	publisher := models.Publishers{}
	result := repo.db.Table("Publishers").First(id).Take(&publisher)
	if result.Error  != nil {
		panic("Livro não localizado")
	}
	return result, nil
}
