package member

import "context"

// MemberRepository saves the member.
type MemberRepository interface {
	Save(ctx context.Context, member Member) (Member, error)
}
