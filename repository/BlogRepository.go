package repository

import (
	"blog-service/model"
	"errors"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func CheckDBConnection(db *gorm.DB) error {
	if db == nil {
		return errors.New("database connection is nil")
	}
	return nil
}

func (repo *BlogRepository) Get(id string) (model.Blog, error) {
	blog := model.Blog{}
	dbResult := repo.DatabaseConnection.First(&blog, "id = ?", id)

	if dbResult.Error != nil {
		return blog, dbResult.Error
	}

	return blog, nil
}

func (repo *BlogRepository) GetAll() ([]model.Blog, error) {
	var blogs []model.Blog
	dbResult := repo.DatabaseConnection.Find(&blogs)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return blogs, nil
}

func (repo *BlogRepository) Save(blog *model.Blog) error {
	dbResult := repo.DatabaseConnection.Create(blog)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *BlogRepository) Update(blog *model.Blog) error {
	dbResult := repo.DatabaseConnection.Save(blog)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *BlogRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Blog{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
