package service

import (
	"blog-service/repository"
	"fmt"
)

type VoteService struct {
	VoteRepo *repository.VoteRepository
}

func (service *VoteService) GetAllByBlogId(blogID string) (int, error) {
	votes, err := service.VoteRepo.GetAllByBlogId(blogID)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve votes for blog ID %s: %v", blogID, err)
	}

	var totalVotes int
	for _, vote := range votes {
		if vote.IsUpvote {
			totalVotes++
		} else {
			totalVotes--
		}
	}

	return totalVotes, nil
}
