package service

import (
	"blog-service/model"
	"blog-service/repository"
	"fmt"
)

type BlogService struct {
	BlogRepo *repository.BlogRepository
}

func (service *BlogService) Get(id string) (*model.Blog, error) {
	blog, err := service.BlogRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("blog with id %s not found", id))
	}
	return &blog, nil
}

func (service *BlogService) GetAll() ([]model.Blog, error) {
	blogs, err := service.BlogRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve blogs: %v", err)
	}
	return blogs, nil
}

func (service *BlogService) GetAllPaged(page int, pageSize int) (model.PagedResultBlog, error) {
	blogs, err := service.BlogRepo.GetPaged(page, pageSize)
	if err != nil {
		return model.PagedResultBlog{}, fmt.Errorf("failed to retrieve blogs: %v", err)
	}
	return blogs, nil
}

func (service *BlogService) Save(blog *model.Blog) error {
	err := service.BlogRepo.Save(blog)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogService) Update(blog *model.Blog) error {
	existingBlog, err := service.BlogRepo.Get(blog.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find blog with ID %s: %v", blog.ID, err)
	}

	existingBlog.Title = blog.Title
	existingBlog.Description = blog.Description
	existingBlog.Status = blog.Status
	existingBlog.Image = blog.Image
	existingBlog.Category = blog.Category

	err = service.BlogRepo.Update(&existingBlog)
	if err != nil {
		return fmt.Errorf("failed to update blog: %v", err)
	}
	return nil
}

func (service *BlogService) Delete(id string) error {
	_, err := service.BlogRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find blog with ID %s: %v", id, err)
	}

	err = service.BlogRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete blog: %v", err)
	}
	return nil
}

func (service *BlogService) GetByAuthorId(userId int) ([]model.Blog, error) {
	blogs, err := service.BlogRepo.GetByAuthorId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve blogs for user with ID %d: %v", userId, err)
	}
	return blogs, nil
}
