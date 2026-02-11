package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/yuuLab/go-rest-api/internal/domain/member"
)

// GetMemberInput represents the input data for getting a member.
type GetMemberInput struct {
	ID int
}

// GetMemberOutput represents the output data of a member.
type GetMemberOutput struct {
	ID        int
	FirstName string
	LastName  string
	FullName  string
	CreatedAt time.Time
}

// GetMemberUseCaseI defines the interface for member retrieval use case.
type GetMemberUseCaseI interface {
	Execute(ctx context.Context, input GetMemberInput) (*GetMemberOutput, error)
}

// GetMemberUseCase handles the business logic for retrieving a member.
type GetMemberUseCase struct {
	memberRepo member.MemberRepository
}

// NewGetMemberUseCase creates a new instance of GetMemberUseCase.
func NewGetMemberUseCase(memberRepo member.MemberRepository) *GetMemberUseCase {
	return &GetMemberUseCase{
		memberRepo: memberRepo,
	}
}

// Execute retrieves a member by ID.
func (uc *GetMemberUseCase) Execute(ctx context.Context, input GetMemberInput) (*GetMemberOutput, error) {
	// Find member by ID
	m, err := uc.memberRepo.FindByID(ctx, member.ID(input.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to find member: %w", err)
	}

	// Convert to output DTO
	output := &GetMemberOutput{
		ID:        int(m.ID()),
		FirstName: m.FirstName(),
		LastName:  m.LastName(),
		FullName:  m.FullName(),
		CreatedAt: m.CreatedAt(),
	}

	return output, nil
}
