package member

import "context"

// MemberRepository saves and retrieves members.
type MemberRepository interface {
	Save(ctx context.Context, member Member) (Member, error)
	FindByID(ctx context.Context, id ID) (*Member, error)
}
