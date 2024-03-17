package service

import "blog-service/repository"

type CommentService struct {
	CommentRepo *repository.CommentRepository
}
