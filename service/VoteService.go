package service

import (
	"blog-service/model"
	"blog-service/repository"
	"fmt"
)

type VoteService struct {
	VoteRepo       *repository.VoteRepository
	CommentService *CommentService
	BlogService    *BlogService
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

func CheckBlogStatus(service *VoteService, blogID string) {
	totalVotes, err := service.GetTotalByBlogId(blogID)
	if err != nil {
		fmt.Printf("Failed to get total votes for blog ID %s: %v\n", blogID, err)
		return
	}

	blog, err := service.BlogService.Get(blogID)
	if err != nil {
		fmt.Printf("Failed to get blog: %v\n", err)
		return
	}

	if totalVotes <= -10 {
		blog.Status = 2
	} else if totalVotes >= 1 {
		blog.Status = 3
	} else if totalVotes > 3 {
		blog.Status = 4
	}

	err = service.BlogService.Update(blog)
	if err != nil {
		fmt.Printf("Failed to update blog: %v\n", err)
	}
}

func (service *VoteService) Save(vote *model.Vote) error {
	err := service.VoteRepo.Save(vote)
	if err != nil {
		return err
	}
	CheckBlogStatus(service, vote.BlogId.String())

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
	CheckBlogStatus(service, vote.BlogId.String())

	return nil
}

func (service *VoteService) GetAllByBlogId(blogID string) ([]model.Vote, error) {
	blogs, err := service.VoteRepo.GetAllByBlogId(blogID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve votes for blog with ID %s: %v", blogID, err)
	}
	return blogs, nil

}
