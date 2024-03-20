package service

import (
	"blog-service/model"
	"blog-service/repository"
	"fmt"
)

type VoteService struct {
	VoteRepo *repository.VoteRepository
}

/*func (service *VoteService) GetTotalByBlogId(blogID string) (int, error) {
	votes, err := service.VoteRepo.GetAllByBlogId(blogID)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve votes for blog ID %s: %v", blogID, err)
	}

	totalVotes := 0
	for _, vote := range votes {
		if vote.IsUpvote {
			totalVotes++
		} else {
			totalVotes--
		}
	}

	return totalVotes, nil
}*/


func (service *VoteService) GetTotalByBlogId(blogID string) (int, error) {
	votes, err := service.VoteRepo.GetAllByBlogId(blogID)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve votes for blog ID %s: %v", blogID, err)
	}

	totalVotes := 0
	for _, vote := range votes {
		if vote.IsUpvote {
			totalVotes++
		} else {
			totalVotes--
		}
	}

	return totalVotes, nil
}
func (service *VoteService) Save(vote *model.Vote) error {
	err := service.VoteRepo.Save(vote)
	if err != nil {
		return err
	}
	return nil
}

func (service *VoteService) Update(vote *model.Vote) error {
	existingVote, err := service.VoteRepo.Get(vote.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find vote with ID %s: %v", vote.ID, err)
	}

	existingVote.IsUpvote = vote.IsUpvote
	existingVote.CreationTime = vote.CreationTime

	err = service.VoteRepo.Update(&existingVote)
	if err != nil {
		return fmt.Errorf("failed to update vote: %v", err)
	}
	return nil
}

func (service *VoteService) GetAllByBlogId(blogID string) ([]model.Vote, error) {
	blogs, err := service.VoteRepo.GetAllByBlogId(blogID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve votes for blog with ID %s: %v", blogID, err)
	}
	return blogs, nil

}
