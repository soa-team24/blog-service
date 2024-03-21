package service

import (
	"blog-service/model"
	"blog-service/repository"
	"fmt"
	"log"
	"time"
)

type CommentService struct {
	CommentRepo *repository.CommentRepository
}

func (service *CommentService) Get(id string) (*model.Comment, error) {
	comment, err := service.CommentRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("comment with id %s not found", id))
	}
	return &comment, nil
}

func (service *CommentService) GetAllByBlogId(blogID string) ([]model.Comment, error) {
	comments, err := service.CommentRepo.GetAllByBlogId(blogID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve comments for blog with ID %s: %v", blogID, err)
	}
	return comments, nil
}

func (service *CommentService) GetTotalByBlogID(blogID string) (int, error) {

	comments, err := service.GetAllByBlogId(blogID)
	if err != nil {
		log.Fatal(err)
	}

	numComments := len(comments)

	return numComments, nil
}

func (service *CommentService) Save(comment *model.Comment) error {
	err := service.CommentRepo.Save(comment)
	if err != nil {
		return err
	}
	return nil
}

func (service *CommentService) Update(comment *model.Comment) error {
	existingComment, err := service.CommentRepo.Get(comment.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find blog with ID %s: %v", comment.ID, err)
	}

	existingComment.Text = comment.Text
	existingComment.LastModification = time.Now()

	err = service.CommentRepo.Update(&existingComment)
	if err != nil {
		return fmt.Errorf("failed to update comment: %v", err)
	}
	return nil
}

func (service *CommentService) Delete(id string) error {
	_, err := service.CommentRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find comment with ID %s: %v", id, err)
	}

	err = service.CommentRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %v", err)
	}
	return nil
}
