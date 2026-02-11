package member

import (
	"strings"
	"time"
)

// ID expresses the Member aggregate identifier.
type ID int64

// Member represents the member aggregate root.
type Member struct {
	createdAt time.Time
	firstName string
	lastName  string
	id        ID
}

// New instantiates a brand-new member aggregate ensuring invariants hold.
func New(firstName, lastName string) (*Member, error) {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	return &Member{
		firstName: firstName,
		lastName:  lastName,
		createdAt: time.Now(),
	}, nil
}

// Reconstitute rebuilds a persisted member aggregate.
func Reconstitute(id ID, firstName, lastName string, createdAt time.Time) *Member {
	return &Member{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		createdAt: createdAt,
	}
}

// ID returns the aggregate identifier.
func (m *Member) ID() ID {
	return m.id
}

// FirstName returns the member's first name.
func (m *Member) FirstName() string {
	return m.firstName
}

// LastName returns the member's last name.
func (m *Member) LastName() string {
	return m.lastName
}

// FullName returns the display-friendly full name.
func (m *Member) FullName() string {
	if m.firstName == "" && m.lastName == "" {
		return ""
	}
	if m.firstName == "" {
		return m.lastName
	}
	if m.lastName == "" {
		return m.firstName
	}
	return m.firstName + " " + m.lastName
}

// CreatedAt returns the creation time.
func (m *Member) CreatedAt() time.Time {
	return m.createdAt
}
