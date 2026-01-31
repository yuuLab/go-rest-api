package repository

import (
	"context"
	"fmt"

	"github.com/yuuLab/go-rest-api/internal/domain/member"
	"github.com/yuuLab/go-rest-api/internal/infrastructure/db/query/model"
)

// MemberRepository is repository for member.
type MemberRepository struct{}

// Save saves the member.
func (r MemberRepository) Save(ctx context.Context, entity member.Member) (*member.Member, error) {
	param := model.CreateMemberParams{
		FirstName: entity.FirstName(),
		LastName:  entity.FirstName(),
		CreatedAt: entity.CreatedAt(),
	}

	q := model.New(nil)
	result, err := q.CreateMember(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("failed to create member: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get lastInsertId: %w", err)
	}

	return member.Reconstitute(member.ID(id), entity.FirstName(), entity.LastName(), entity.CreatedAt()), nil
}
