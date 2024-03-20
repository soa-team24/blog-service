package repository

import (
	"blog-service/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CommentRepository) Get(id string) (model.Comment, error) {
	comment := model.Comment{}
	dbResult := repo.DatabaseConnection.First(&comment, "id = ?", id)

	if dbResult.Error != nil {
		return comment, dbResult.Error
	}

	return comment, nil
}

func (repo *CommentRepository) GetAllByBlogId(blogID string) ([]model.Comment, error) {
	var comments []model.Comment
	result := repo.DatabaseConnection.Find(&comments, "blog_id = ?", blogID)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (repo *CommentRepository) Save(comment *model.Comment) error {
	dbResult := repo.DatabaseConnection.Create(comment)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *CommentRepository) Update(comment *model.Comment) error {
	dbResult := repo.DatabaseConnection.Save(comment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *CommentRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Comment{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
