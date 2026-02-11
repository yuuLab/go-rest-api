package repository

import (
	"context"
	"database/sql"
	"errors"
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

// FindByID finds a member by ID.
func (r MemberRepository) FindByID(ctx context.Context, id member.ID) (*member.Member, error) {
	q := model.New(nil)
	result, err := q.FindMember(ctx, uint64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("member not found: id=%d", id)
		}
		return nil, fmt.Errorf("failed to find member: %w", err)
	}

	return member.Reconstitute(
		member.ID(result.ID),
		result.FirstName,
		result.LastName,
		result.CreatedAt,
	), nil
}
