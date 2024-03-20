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
func (repo *BlogRepository) GetPaged(page, pageSize int) (model.PagedResultBlog, error) {
	var blogs []model.Blog
	var totalCount int64

	err := repo.DatabaseConnection.Model(&model.Blog{}).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&blogs).Error
	if err != nil {
		return model.PagedResultBlog{}, err
	}

	return model.PagedResultBlog{
		Results:    blogs,
		TotalCount: int(totalCount),
	}, nil
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
	dbResult := repo.DatabaseConnection.Delete(&model.Blog{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *BlogRepository) GetByAuthorId(userId int) ([]model.Blog, error) {
	var blogs []model.Blog
	result := repo.DatabaseConnection.Find(&blogs, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

func (repo *BlogRepository) GetByStatus(status model.Status) ([]model.Blog, error) {
	var blogs []model.Blog
	result := repo.DatabaseConnection.Find(&blogs, "status = ?", status)
	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
}
