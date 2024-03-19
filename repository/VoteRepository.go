package repository

import (
	"blog-service/model"

	"gorm.io/gorm"
)

type VoteRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *VoteRepository) Get(id string) (model.Vote, error) {
	vote := model.Vote{}
	dbResult := repo.DatabaseConnection.First(&vote, "id = ?", id)

	if dbResult.Error != nil {
		return vote, dbResult.Error
	}

	return vote, nil
}

func (repo *VoteRepository) Save(vote *model.Vote) error {
	dbResult := repo.DatabaseConnection.Create(vote)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *VoteRepository) Update(vote *model.Vote) error {
	dbResult := repo.DatabaseConnection.Save(vote)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *VoteRepository) GetAllByBlogId(blogID string) ([]model.Vote, error) {
	var votes []model.Vote
	result := repo.DatabaseConnection.Find(&votes, "blog_id = ?", blogID)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}
