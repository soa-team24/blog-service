package service

import "blog-service/repository"

type VoteService struct {
	VoteRepo *repository.VoteRepository
}
