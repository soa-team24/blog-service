package service

import "blog-service/repository"

type VoteService struct {
	BlogRepo *repository.VoteRepository
}
