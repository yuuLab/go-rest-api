package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/yuuLab/go-rest-api/internal/domain/member"
)

// RegisterMemberInput represents the input data for member registration.
type RegisterMemberInput struct {
	FirstName   string
	LastName    string
	RawPassword string
}

// RegisterMemberOutput represents the output data after member registration.
type RegisterMemberOutput struct {
	ID        int
	FirstName string
	LastName  string
	CreatedAt time.Time
}

// RegisterMemberUseCaseI defines the interface for member registration use case.
type RegisterMemberUseCaseI interface {
	Execute(ctx context.Context, input RegisterMemberInput) (*RegisterMemberOutput, error)
}

// RegisterMemberUseCase handles the business logic for registering a new member.
type RegisterMemberUseCase struct {
	memberRepo member.MemberRepository
}

// NewRegisterMemberUseCase creates a new instance of RegisterMemberUseCase.
func NewRegisterMemberUseCase(memberRepo member.MemberRepository) *RegisterMemberUseCase {
	return &RegisterMemberUseCase{
		memberRepo: memberRepo,
	}
}

// Execute performs the member registration process.
func (uc *RegisterMemberUseCase) Execute(ctx context.Context, input RegisterMemberInput) (*RegisterMemberOutput, error) {
	newMember, err := member.New(input.FirstName, input.LastName)
	if err != nil {
		return nil, fmt.Errorf("failed to create member: %w", err)
	}

	savedMember, err := uc.memberRepo.Save(ctx, *newMember)
	if err != nil {
		return nil, fmt.Errorf("failed to save member: %w", err)
	}

	output := &RegisterMemberOutput{
		ID:        int(savedMember.ID()),
		FirstName: savedMember.FirstName(),
		LastName:  savedMember.LastName(),
		CreatedAt: savedMember.CreatedAt(),
	}

	return output, nil
}
